# Module 11 – Concurrence avancée

Bienvenue dans ce onzième module ! À la fin de ce module, vous saurez protéger les données partagées avec les mutex, éviter les conditions de course, utiliser les patterns avancés de concurrence, et concevoir des programmes concurrents robustes et sans fuite de mémoire.

---

## 0. Pourquoi la concurrence avancée ?

### 0.1. Problème : les données partagées non protégées

Dans le module 9, nous avons vu comment lancer des goroutines et les faire communiquer via des channels. Mais que se passe-t-il quand plusieurs goroutines **accèdent à la même variable** ?

```go
var compteur int

func incrementer(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
        compteur++  // ⚠️ Opération NON atomique !
    }
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go incrementer(&wg)
    }
    wg.Wait()
    fmt.Println("Compteur final:", compteur)
    // Résultat attendu: 100000
    // Résultat réel:   82341, 91234, 100000 (variable !)
}
```

**Pourquoi ?** L'opération `compteur++` n'est pas atomique :
1. Lire la valeur actuelle de `compteur`
2. Ajouter 1
3. Écrire la nouvelle valeur

Si deux goroutines lisent la même valeur (ex: 42), toutes deux ajoutent 1 et écrivent 43. On a perdu un incrément !

C'est une **condition de course** (race condition).

### 0.2. Solutions de synchronisation

| Mécanisme | Quand l'utiliser |
|-----------|------------------|
| **Channel** | Communication entre goroutines |
| **Mutex** | Protection de données partagées |
| **RWMutex** | Lectures fréquentes, écritures rares |
| **Once** | Initialisation unique |
| **Cond** | Attente conditionnelle |

---

## 1. Les mutex (`sync.Mutex`)

Un **mutex** (mutual exclusion) agit comme un verrou : une seule goroutine peut le détenir à la fois.

### 1.1. Utilisation basique

```go
import "sync"

type Compteur struct {
    mu      sync.Mutex  // Le mutex protège compteur
    valeur  int
}

func (c *Compteur) Incrementer() {
    c.mu.Lock()         // Prendre le verrou
    defer c.mu.Unlock() // Le libérer à la fin
    c.valeur++
}

func (c *Compteur) Valeur() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.valeur
}

func main() {
    compteur := &Compteur{}
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            compteur.Incrementer()
        }()
    }

    wg.Wait()
    fmt.Println(compteur.Valeur())  // 1000 (toujours !)
}
```

### 1.2. Règles d'utilisation des mutex

✅ **À faire :**
- Toujours `defer Unlock()` après `Lock()`
- Garder les sections critiques **courtes**
- Ne jamais copier un mutex (passer par pointeur)

❌ **À ne pas faire :**
- Oublier de déverrouiller → **deadlock**
- Verrouiller puis appeler une fonction qui reverrouille → **deadlock**
- Passer un mutex par valeur (il serait copié)

### 1.3. Piège classique : copie de mutex

```go
// ❌ MAUVAIS : la méthode reçoit une copie du compteur
func (c Compteur) Incrementer() {  // c est une copie !
    c.mu.Lock()  // Le mutex original n'est pas verrouillé
    defer c.mu.Unlock()
    c.valeur++
}

// ✅ BON : utiliser un pointeur
func (c *Compteur) Incrementer() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.valeur++
}
```

---

## 2. `sync.RWMutex` – Lecture/Écriture

`RWMutex` permet :
- **Plusieurs lecteurs** simultanés (`RLock`)
- **Un seul écrivain** exclusif (`Lock`)

```go
type Cache struct {
    mu    sync.RWMutex
    data  map[string]string
}

func (c *Cache) Get(key string) string {
    c.mu.RLock()         // Verrou de lecture
    defer c.mu.RUnlock()
    return c.data[key]
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()          // Verrou d'écriture (exclusif)
    defer c.mu.Unlock()
    c.data[key] = value
}

func (c *Cache) Size() int {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return len(c.data)
}
```

**Comparaison des performances :**

```go
// Test avec 1000 lectures et 10 écritures
// Mutex simple:  ~150ms
// RWMutex:      ~45ms  (gain de 3x pour les lectures)
```

💡 **Quand utiliser RWMutex ?**
- Lectures **beaucoup plus fréquentes** que les écritures
- Cache, configuration, compteurs, etc.

---

## 3. `sync.Once` – Initialisation unique

`sync.Once` garantit qu'une fonction est exécutée **une seule fois**, même si appelée par plusieurs goroutines.

### 3.1. Pattern du singleton

```go
type Config struct {
    Port int
    Host string
}

var (
    instance *Config
    once     sync.Once
)

func GetConfig() *Config {
    once.Do(func() {
        fmt.Println("Initialisation unique...")
        instance = &Config{
            Port: 8080,
            Host: "localhost",
        }
    })
    return instance
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            cfg := GetConfig()
            fmt.Printf("Config: %+v\n", cfg)
        }()
    }
    wg.Wait()
}
// "Initialisation unique..." n'apparaît qu'une seule fois
```

### 3.2. Initialisation de ressource coûteuse

```go
var (
    db     *sql.DB
    dbOnce sync.Once
)

func GetDB() *sql.DB {
    dbOnce.Do(func() {
        db, _ = sql.Open("mysql", "user:pass@/dbname")
        db.SetMaxOpenConns(25)
    })
    return db
}
```

---

## 4. `sync.Map` – Map concurrente

Pour les cas simples, Go fournit `sync.Map` (thread-safe).

```go
var m sync.Map

func main() {
    var wg sync.WaitGroup

    // Écritures concurrentes
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            m.Store(fmt.Sprintf("key%d", i), i*10)
        }(i)
    }

    // Lectures concurrentes
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            if val, ok := m.Load(fmt.Sprintf("key%d", i)); ok {
                fmt.Printf("key%d = %v\n", i, val)
            }
        }(i)
    }

    wg.Wait()

    // Parcours
    m.Range(func(key, value interface{}) bool {
        fmt.Printf("%v: %v\n", key, value)
        return true  // continue
    })
}
```

**⚠️ Quand ne PAS utiliser `sync.Map` :**
- Opérations fréquentes sur les mêmes clés
- Nombre de clés fixe et connu
- Dans ces cas, un `map` + `RWMutex` est plus performant

---

## 5. `sync.Cond` – Attente conditionnelle

`Cond` permet à des goroutines d'attendre qu'une **condition** devienne vraie.

### 5.1. Producteur-Consommateur avec Cond

```go
type Queue struct {
    items []int
    cond  *sync.Cond
}

func NewQueue() *Queue {
    return &Queue{
        items: make([]int, 0),
        cond:  sync.NewCond(&sync.Mutex{}),
    }
}

func (q *Queue) Push(item int) {
    q.cond.L.Lock()
    defer q.cond.L.Unlock()

    q.items = append(q.items, item)
    fmt.Printf("📥 Produit: %d (taille: %d)\n", item, len(q.items))

    // Réveiller UNE goroutine qui attend
    q.cond.Signal()
    // Ou toutes: q.cond.Broadcast()
}

func (q *Queue) Pop() int {
    q.cond.L.Lock()
    defer q.cond.L.Unlock()

    // Attendre qu'il y ait des éléments
    for len(q.items) == 0 {
        fmt.Println("😴 Consommateur attend...")
        q.cond.Wait()  // Libère le verrou et attend
    }

    item := q.items[0]
    q.items = q.items[1:]
    fmt.Printf("📤 Consomme: %d (taille: %d)\n", item, len(q.items))
    return item
}

func main() {
    queue := NewQueue()
    var wg sync.WaitGroup

    // 2 consommateurs
    for i := 1; i <= 2; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 5; j++ {
                item := queue.Pop()
                time.Sleep(100 * time.Millisecond)
            }
        }(i)
    }

    // 1 producteur
    time.Sleep(500 * time.Millisecond)
    for i := 1; i <= 10; i++ {
        queue.Push(i)
        time.Sleep(50 * time.Millisecond)
    }

    wg.Wait()
}
```

### 5.2. Différence entre Signal et Broadcast

| Méthode | Effet |
|---------|-------|
| `Signal()` | Réveille **une** goroutine qui attend |
| `Broadcast()` | Réveille **toutes** les goroutines qui attendent |

---

## 6. Patterns avancés de concurrence

### 6.1. Pattern Worker Pool (pool de workers)

```go
type Job struct {
    ID     int
    Payload string
}

func worker(id int, jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d traite job %d: %s\n", id, job.ID, job.Payload)
        time.Sleep(100 * time.Millisecond)  // Simulation de travail
        results <- fmt.Sprintf("Job %d terminé par worker %d", job.ID, id)
    }
}

func main() {
    const nbWorkers = 3
    const nbJobs = 10

    jobs := make(chan Job, nbJobs)
    results := make(chan string, nbJobs)
    var wg sync.WaitGroup

    // Lancement des workers
    for i := 1; i <= nbWorkers; i++ {
        wg.Add(1)
        go worker(i, jobs, results, &wg)
    }

    // Envoi des jobs
    for i := 1; i <= nbJobs; i++ {
        jobs <- Job{ID: i, Payload: fmt.Sprintf("Donnée %d", i)}
    }
    close(jobs)

    // Attente des workers
    go func() {
        wg.Wait()
        close(results)
    }()

    // Collecte des résultats
    for result := range results {
        fmt.Println("✅", result)
    }
}
```

### 6.2. Pattern Fan-Out / Fan-In

**Fan-Out** : distribuer le travail à plusieurs goroutines
**Fan-In** : rassembler les résultats

```go
// Générateur de nombres
func generate(numbers ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range numbers {
            out <- n
        }
        close(out)
    }()
    return out
}

// Transformation (carré)
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Fan-In: combine plusieurs channels en un seul
func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup

    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for n := range c {
                out <- n
            }
        }(ch)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}

func main() {
    // Génération
    nums := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

    // Fan-Out: 3 workers qui calculent le carré
    worker1 := square(nums)
    worker2 := square(nums)
    worker3 := square(nums)

    // Fan-In: combinaison des résultats
    results := fanIn(worker1, worker2, worker3)

    // Collecte
    for result := range results {
        fmt.Println(result)
    }
}
```

### 6.3. Pattern Pipeline

```go
// Étape 1: génère des nombres
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Étape 2: multiplie par 2
func mult(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * 2
        }
        close(out)
    }()
    return out
}

// Étape 3: filtre les nombres > 10
func filter(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            if n > 10 {
                out <- n
            }
        }
        close(out)
    }()
    return out
}

func main() {
    // Pipeline: gen → mult → filter
    pipeline := filter(mult(gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)))

    for result := range pipeline {
        fmt.Println(result)
    }
    // Output: 12, 14, 16, 18, 20
}
```

### 6.4. Pattern Rate Limiting (limitation de débit)

```go
func rateLimiter(requests []string, limit int) {
    // Token bucket: canal bufferisé avec le nombre de jetons
    tokens := make(chan time.Time, limit)

    // Remplissage du bucket
    for i := 0; i < limit; i++ {
        tokens <- time.Now()
    }

    // Goroutine qui ajoute un jeton toutes les secondes
    go func() {
        ticker := time.NewTicker(1 * time.Second)
        for range ticker.C {
            select {
            case tokens <- time.Now():
                fmt.Println("➕ Jeton ajouté")
            default:
                fmt.Println("⚠️  Bucket plein")
            }
        }
    }()

    for _, req := range requests {
        <-tokens  // Prendre un jeton (bloque si vide)
        fmt.Printf("✅ Traitement de: %s\n", req)
    }
}
```

### 6.5. Pattern Circuit Breaker

```go
type CircuitBreaker struct {
    mu           sync.Mutex
    maxFailures  int
    failures     int
    state        string  // "closed", "open", "half-open"
    lastFailure  time.Time
    timeout      time.Duration
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        maxFailures: maxFailures,
        state:       "closed",
        timeout:     timeout,
    }
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mu.Lock()

    if cb.state == "open" {
        if time.Since(cb.lastFailure) > cb.timeout {
            cb.state = "half-open"
            fmt.Println("🔄 Circuit: half-open (test)")
        } else {
            cb.mu.Unlock()
            return fmt.Errorf("circuit breaker open")
        }
    }
    cb.mu.Unlock()

    // Exécution
    err := fn()

    cb.mu.Lock()
    defer cb.mu.Unlock()

    if err != nil {
        cb.failures++
        cb.lastFailure = time.Now()

        if cb.failures >= cb.maxFailures {
            cb.state = "open"
            fmt.Printf("🔌 Circuit ouvert après %d échecs\n", cb.failures)
        }
        return err
    }

    // Succès
    if cb.state == "half-open" {
        fmt.Println("✅ Circuit rétabli")
        cb.state = "closed"
    }
    cb.failures = 0
    return nil
}

// Utilisation
func appelExterne() error {
    if time.Now().Second()%3 == 0 {
        return fmt.Errorf("service indisponible")
    }
    return nil
}

func main() {
    cb := NewCircuitBreaker(3, 5*time.Second)

    for i := 1; i <= 20; i++ {
        err := cb.Call(appelExterne)
        if err != nil {
            fmt.Printf("[%d] ❌ Erreur: %v\n", i, err)
        } else {
            fmt.Printf("[%d] ✅ Succès\n", i)
        }
        time.Sleep(1 * time.Second)
    }
}
```

---

## 7. Détection des conditions de course

Go intègre un **détecteur de race conditions** :

```bash
# Exécuter avec le détecteur
go run -race main.go

# Compiler avec le détecteur
go build -race main.go

# Tester avec le détecteur
go test -race ./...
```

**Exemple de détection :**

```go
// race.go
var compteur int

func main() {
    go func() { compteur++ }()
    go func() { compteur++ }()
    time.Sleep(time.Second)
}
```

```bash
$ go run -race race.go
==================
WARNING: DATA RACE
Write at 0x00c0000a2008 by goroutine 7:
  main.main.func2()
      /tmp/race.go:8 +0x3e

Previous write at 0x00c0000a2008 by goroutine 6:
  main.main.func1()
      /tmp/race.go:7 +0x3e
==================
Found 1 data race(s)
```

---

## 8. Éviter les deadlocks (interblocages)

### 8.1. Qu'est-ce qu'un deadlock ?

Deux goroutines s'attendent mutuellement, aucune ne peut avancer.

```go
// ❌ Deadlock simple
func main() {
    ch := make(chan int)
    ch <- 42  // Bloque indéfiniment (personne ne reçoit)
}
```

```go
// ❌ Deadlock avec mutex
var mu1, mu2 sync.Mutex

func f1() {
    mu1.Lock()
    time.Sleep(time.Millisecond)
    mu2.Lock()  // Attend que f2 libère mu2
    mu2.Unlock()
    mu1.Unlock()
}

func f2() {
    mu2.Lock()
    time.Sleep(time.Millisecond)
    mu1.Lock()  // Attend que f1 libère mu1
    mu1.Unlock()
    mu2.Unlock()
}
```

### 8.2. Outils de détection

```bash
# Go peut détecter certains deadlocks à l'exécution
go run main.go
# fatal error: all goroutines are asleep - deadlock!
```

### 8.3. Bonnes pratiques anti-deadlock

✅ **Toujours :**
- `defer Unlock()` après chaque `Lock()`
- Utiliser un ordre d'acquisition des verrous
- Éviter les mutex dans les mutex

```go
// ✅ Ordre fixe d'acquisition
func transfert(compteA, compteB *Compte, montant int) {
    // Toujours verrouiller le compte avec l'ID le plus petit en premier
    if compteA.ID < compteB.ID {
        compteA.mu.Lock()
        compteB.mu.Lock()
    } else {
        compteB.mu.Lock()
        compteA.mu.Lock()
    }
    defer compteA.mu.Unlock()
    defer compteB.mu.Unlock()

    compteA.Solde -= montant
    compteB.Solde += montant
}
```

---

## TP final – Module 11 – Cache concurrent avec expiration

### Énoncé

Créez un **cache concurrent** avec les fonctionnalités suivantes :

1. **Stockage clé-valeur** thread-safe
2. **Expiration** automatique des entrées (TTL)
3. **Nettoyage** périodique des entrées expirées
4. **Statistiques** : nombre d'entrées, hits/miss
5. **Mécanisme de lecture** qui rafraîchit le TTL

### Structure du cache

```go
type Item struct {
    valeur     interface{}
    expiration time.Time
}

type Cache struct {
    mu       sync.RWMutex
    items    map[string]Item
    ttl      time.Duration
    stats    Stats
}

type Stats struct {
    Hits  int
    Miss  int
    Size  int
}
```

### Fonctions à implémenter

- `NewCache(ttl time.Duration, cleanupInterval time.Duration) *Cache`
- `Set(key string, value interface{})`
- `Get(key string) (interface{}, bool)` – rafraîchit le TTL si trouvé
- `Delete(key string)`
- `Stats() Stats`
- Nettoyage automatique en arrière-plan

### À vous de jouer !

Codez la solution par vous-même avant de regarder la correction.

---
.
.
.
.
.
.
.
.
.
.
---
### Solution complète

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Item représente une entrée du cache
type Item struct {
    valeur     interface{}
    expiration time.Time
}

// Cache concurrent avec expiration
type Cache struct {
    mu              sync.RWMutex
    items           map[string]Item
    ttl             time.Duration
    cleanupInterval time.Duration
    stats           Stats
    stopCleanup     chan bool
}

// Statistiques du cache
type Stats struct {
    Hits  int
    Miss  int
    Size  int
    mu    sync.RWMutex
}

// Nouveau cache
func NewCache(ttl, cleanupInterval time.Duration) *Cache {
    cache := &Cache{
        items:           make(map[string]Item),
        ttl:             ttl,
        cleanupInterval: cleanupInterval,
        stopCleanup:     make(chan bool),
    }

    // Démarrer le nettoyeur automatique
    go cache.startCleanup()

    return cache
}

// Ajouter ou modifier une valeur
func (c *Cache) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.items[key] = Item{
        valeur:     value,
        expiration: time.Now().Add(c.ttl),
    }

    c.updateStatsSize()
}

// Récupérer une valeur (rafraîchit le TTL)
func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    item, exists := c.items[key]
    if !exists {
        c.incrementMiss()
        return nil, false
    }

    // Vérifier l'expiration
    if time.Now().After(item.expiration) {
        delete(c.items, key)
        c.updateStatsSize()
        c.incrementMiss()
        return nil, false
    }

    // Rafraîchir le TTL (lecture = prolongation)
    item.expiration = time.Now().Add(c.ttl)
    c.items[key] = item

    c.incrementHit()
    return item.valeur, true
}

// Supprimer une clé
func (c *Cache) Delete(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()

    delete(c.items, key)
    c.updateStatsSize()
}

// Obtenir les statistiques
func (c *Cache) Stats() (hits, miss, size int) {
    c.stats.mu.RLock()
    defer c.stats.mu.RUnlock()

    c.mu.RLock()
    defer c.mu.RUnlock()

    return c.stats.Hits, c.stats.Miss, len(c.items)
}

// Nettoyage périodique
func (c *Cache) startCleanup() {
    ticker := time.NewTicker(c.cleanupInterval)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            c.cleanup()
        case <-c.stopCleanup:
            return
        }
    }
}

// Supprimer les entrées expirées
func (c *Cache) cleanup() {
    c.mu.Lock()
    defer c.mu.Unlock()

    now := time.Now()
    for key, item := range c.items {
        if now.After(item.expiration) {
            delete(c.items, key)
        }
    }
    c.updateStatsSize()
}

// Arrêter le nettoyeur
func (c *Cache) Stop() {
    close(c.stopCleanup)
}

// Métriques internes
func (c *Cache) incrementHit() {
    c.stats.mu.Lock()
    defer c.stats.mu.Unlock()
    c.stats.Hits++
}

func (c *Cache) incrementMiss() {
    c.stats.mu.Lock()
    defer c.stats.mu.Unlock()
    c.stats.Miss++
}

func (c *Cache) updateStatsSize() {
    c.stats.mu.Lock()
    defer c.stats.mu.Unlock()
    c.stats.Size = len(c.items)
}

// Affichage convivial
func (c *Cache) String() string {
    c.mu.RLock()
    defer c.mu.RUnlock()

    hits, miss, size := c.Stats()
    return fmt.Sprintf("Cache{size: %d, hits: %d, miss: %d, hitrate: %.1f%%}",
        size, hits, miss, float64(hits)/float64(hits+miss)*100)
}

// ===== TESTS =====

func main() {
    fmt.Println("=== Test du Cache Concurrent ===\n")

    // Cache avec TTL de 2 secondes, nettoyage toutes les 1 seconde
    cache := NewCache(2*time.Second, 1*time.Second)
    defer cache.Stop()

    // Test 1: Set et Get
    fmt.Println("📝 Test 1: Set/Get basique")
    cache.Set("nom", "Alice")
    cache.Set("age", 30)

    if val, ok := cache.Get("nom"); ok {
        fmt.Printf("   ✅ Get('nom') = %v\n", val)
    }

    // Test 2: Expiration
    fmt.Println("\n⏰ Test 2: Expiration")
    cache.Set("temp", "expire vite")
    fmt.Printf("   Avant expiration: %v\n", cache.getOrPanic("temp"))
    time.Sleep(3 * time.Second)
    if _, ok := cache.Get("temp"); !ok {
        fmt.Println("   ✅ La clé 'temp' a bien expiré")
    }

    // Test 3: Concurrence
    fmt.Println("\n🚀 Test 3: Accès concurrents")
    var wg sync.WaitGroup

    // 100 lectures concurrentes
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            cache.Set(fmt.Sprintf("key%d", id), id*10)
            cache.Get(fmt.Sprintf("key%d", id))
        }(i)
    }

    wg.Wait()
    hits, miss, size := cache.Stats()
    fmt.Printf("   Statistiques: hits=%d, miss=%d, size=%d\n", hits, miss, size)

    // Test 4: Rafraîchissement du TTL
    fmt.Println("\n🔄 Test 4: Rafraîchissement du TTL")
    cache.Set("refresh", "valeur")
    fmt.Printf("   Get #1: %v\n", cache.getOrPanic("refresh"))

    time.Sleep(1500 * time.Millisecond)
    fmt.Printf("   Get #2 (après 1.5s): %v\n", cache.getOrPanic("refresh"))

    time.Sleep(1500 * time.Millisecond)
    fmt.Printf("   Get #3 (après encore 1.5s): %v\n", cache.getOrPanic("refresh"))

    // Test 5: Statistiques finales
    fmt.Printf("\n📊 Statistiques finales: %s\n", cache)
}

func (c *Cache) getOrPanic(key string) string {
    if val, ok := c.Get(key); ok {
        return fmt.Sprintf("%v", val)
    }
    return "<expiré>"
}
```

### Exemple d'utilisation avancée

```go
// Cache avec génération automatique
type AutoCache struct {
    *Cache
    loader func(string) (interface{}, error)
}

func NewAutoCache(ttl, cleanup time.Duration, loader func(string) (interface{}, error)) *AutoCache {
    return &AutoCache{
        Cache:  NewCache(ttl, cleanup),
        loader: loader,
    }
}

func (ac *AutoCache) GetOrLoad(key string) (interface{}, error) {
    if val, ok := ac.Get(key); ok {
        return val, nil
    }

    // Chargement depuis la source
    val, err := ac.loader(key)
    if err != nil {
        return nil, err
    }

    ac.Set(key, val)
    return val, nil
}

// Utilisation
func main() {
    loader := func(key string) (interface{}, error) {
        fmt.Printf("Chargement de '%s' depuis la base...\n", key)
        time.Sleep(100 * time.Millisecond)
        return fmt.Sprintf("Donnée de %s", key), nil
    }

    ac := NewAutoCache(5*time.Second, 1*time.Second, loader)

    // Premier accès : charge
    val, _ := ac.GetOrLoad("user:123")
    fmt.Println(val)

    // Deuxième accès : cache
    val, _ = ac.GetOrLoad("user:123")
    fmt.Println(val)
}
```

### Tests supplémentaires

```go
// Test de performance
func BenchmarkCacheSet(b *testing.B) {
    cache := NewCache(1*time.Minute, 10*time.Second)
    defer cache.Stop()

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        cache.Set(fmt.Sprintf("key%d", i), i)
    }
}

// Test de concurrence
func TestCacheConcurrent(t *testing.T) {
    cache := NewCache(1*time.Minute, 10*time.Second)
    defer cache.Stop()

    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            cache.Set(fmt.Sprintf("k%d", i), i)
            if val, ok := cache.Get(fmt.Sprintf("k%d", i)); !ok || val != i {
                t.Errorf("Valeur incorrecte pour k%d", i)
            }
        }(i)
    }
    wg.Wait()
}
```

---

## Récapitulatif des acquis

| Concept | Syntaxe | Utilisation |
|---------|---------|-------------|
| Mutex | `sync.Mutex`, `Lock()`, `Unlock()` | Protection de données partagées |
| RWMutex | `RLock()`, `RUnlock()` | Lectures fréquentes |
| Once | `sync.Once`, `Do()` | Initialisation unique |
| Map concurrent | `sync.Map` | Map thread-safe simple |
| Cond | `sync.Cond`, `Wait()`, `Signal()` | Attente conditionnelle |
| Race detector | `go run -race` | Détection des conditions de course |

---

## Prochain module (aperçu)

Dans le Module 12, nous découvrirons :
- Le package `context` pour l'annulation et les timeouts
- La propagation de valeurs à travers la chaîne d'appels
- Les patterns d'annulation de goroutines

---

**Félicitations !** Vous maîtrisez maintenant les techniques avancées de concurrence en Go. Vous savez protéger les données partagées, éviter les deadlocks, et concevoir des patterns robustes comme les worker pools et les circuit breakers. 🚀
