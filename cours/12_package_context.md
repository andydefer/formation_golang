# Module 12 – Le package `context`

Bienvenue dans ce douzième module ! À la fin de ce module, vous saurez utiliser le package `context` pour gérer les annulations, les timeouts, et propager des valeurs à travers la chaîne d'appels. Vous pourrez ainsi créer des programmes réactifs et contrôlables, capables de s'arrêter proprement quand c'est nécessaire.

---

## 0. Pourquoi le package `context` ?

### 0.1. Problème : comment arrêter une goroutine ?

Imaginez un serveur qui traite des requêtes longues. Un client annule sa requête (fermeture du navigateur, timeout). Comment arrêter proprement le travail en cours ?

**Sans `context`** :

```go
func travailLong() {
    for i := 0; i < 100; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println("Étape", i)
    }
}

func main() {
    go travailLong()

    time.Sleep(1 * time.Second)
    fmt.Println("L'utilisateur annule...")
    // ❌ Impossible d'arrêter la goroutine !
    time.Sleep(5 * time.Second)
}
```

**Avec `context`** :

```go
func travailLong(ctx context.Context) {
    for i := 0; i < 100; i++ {
        select {
        case <-ctx.Done():
            fmt.Println("Annulation reçue, arrêt propre")
            return
        default:
            time.Sleep(100 * time.Millisecond)
            fmt.Println("Étape", i)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    go travailLong(ctx)

    time.Sleep(1 * time.Second)
    fmt.Println("L'utilisateur annule...")
    cancel()  // ✅ La goroutine s'arrête proprement

    time.Sleep(500 * time.Millisecond)
}
```

### 0.2. Le package `context` : à quoi ça sert ?

| Fonctionnalité | Description |
|----------------|-------------|
| **Annulation** | Propager un signal d'annulation à travers les goroutines |
| **Timeout** | Annuler automatiquement après une durée |
| **Deadline** | Annuler à une date/heure précise |
| **Propagation** | Passer des valeurs (trace ID, auth) le long de la chaîne |

💡 **Règle d'or :** Le `context.Context` est **toujours le premier paramètre** d'une fonction qui peut être bloquante ou annulable.

```go
func MaFonction(ctx context.Context, autreParam string) error
```

---

## 1. Créer des contextes

### 1.1. Contextes racines

```go
import "context"

// Contexte vide, jamais annulé (point de départ)
ctx := context.Background()

// Contexte vide pour tests (ou quand on ne sait pas quoi mettre)
ctx := context.TODO()
```

| Fonction | Quand l'utiliser |
|----------|------------------|
| `context.Background()` | Point d'entrée (main, serveur HTTP, handlers) |
| `context.TODO()` | En développement, refactoring, ou code temporaire |

### 1.2. Créer des contextes dérivés

```go
// Annulation manuelle
ctx, cancel := context.WithCancel(parentCtx)

// Annulation automatique après 5 secondes
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)

// Annulation à une heure précise
ctx, cancel := context.WithDeadline(parentCtx, time.Now().Add(5*time.Second))

// Contexte avec une valeur (trace ID, auth, etc.)
ctx := context.WithValue(parentCtx, key, value)
```

⚠️ **Toujours appeler `cancel()`** pour libérer les ressources :

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()  // IMPORTANT : libère les ressources
```

---

## 2. Annuler des goroutines

### 2.1. Pattern basique d'annulation

```go
func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
    defer wg.Done()

    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d: arrêt demandé\n", id)
            return
        default:
            fmt.Printf("Worker %d: travail...\n", id)
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    var wg sync.WaitGroup

    // Lancement de 3 workers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(ctx, i, &wg)
    }

    time.Sleep(2 * time.Second)
    fmt.Println("Annulation générale !")
    cancel()  // Tous les workers s'arrêtent

    wg.Wait()
    fmt.Println("Tous les workers sont arrêtés")
}
```

### 2.2. Annulation avec `select` dans une boucle

```go
func telechargerFichier(ctx context.Context, url string) error {
    // Simuler un téléchargement long
    for progression := 0; progression <= 100; progression += 10 {
        select {
        case <-ctx.Done():
            return fmt.Errorf("téléchargement annulé: %w", ctx.Err())
        case <-time.After(100 * time.Millisecond):
            fmt.Printf("%s: %d%%\n", url, progression)
        }
    }
    return nil
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 350*time.Millisecond)
    defer cancel()

    err := telechargerFichier(ctx, "https://example.com/fichier.zip")
    if err != nil {
        fmt.Println("Erreur:", err)
        // Output: Erreur: téléchargement annulé: context deadline exceeded
    }
}
```

### 2.3. Propagation d'annulation en cascade

```go
func niveau3(ctx context.Context) {
    select {
    case <-ctx.Done():
        fmt.Println("Niveau 3: annulation reçue")
        return
    case <-time.After(2 * time.Second):
        fmt.Println("Niveau 3: travail terminé")
    }
}

func niveau2(ctx context.Context) {
    go niveau3(ctx)

    select {
    case <-ctx.Done():
        fmt.Println("Niveau 2: annulation reçue")
        return
    case <-time.After(1 * time.Second):
        fmt.Println("Niveau 2: travail terminé")
    }
}

func niveau1(ctx context.Context) {
    go niveau2(ctx)

    select {
    case <-ctx.Done():
        fmt.Println("Niveau 1: annulation reçue")
        return
    case <-time.After(3 * time.Second):
        fmt.Println("Niveau 1: travail terminé")
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
    defer cancel()

    niveau1(ctx)
    time.Sleep(500 * time.Millisecond)
}
```

**Sortie :**
```
Niveau 2: annulation reçue
Niveau 3: annulation reçue
Niveau 1: annulation reçue
```

---

## 3. Timeout et Deadline

### 3.1. `WithTimeout` – timeout relatif

```go
// Timeout de 3 secondes
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

// Utilisation avec une requête HTTP
req, _ := http.NewRequestWithContext(ctx, "GET", "https://api.example.com", nil)
client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        fmt.Println("Timeout dépassé")
    }
}
```

### 3.2. `WithDeadline` – deadline absolue

```go
// Deadline à 15h30 précises
deadline := time.Date(2024, 1, 15, 15, 30, 0, 0, time.Local)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

### 3.3. Combinaison de timeouts

```go
func operationComplexe(ctx context.Context) error {
    // Étape 1: timeout 1 seconde
    ctx1, cancel1 := context.WithTimeout(ctx, 1*time.Second)
    defer cancel1()

    if err := etape1(ctx1); err != nil {
        return fmt.Errorf("étape 1: %w", err)
    }

    // Étape 2: timeout 2 secondes
    ctx2, cancel2 := context.WithTimeout(ctx, 2*time.Second)
    defer cancel2()

    return etape2(ctx2)
}
```

---

## 4. Propager des valeurs avec `WithValue`

### 4.1. Utilisation basique

```go
type contextKey string  // Type personnalisé pour éviter les collisions

const (
    UserIDKey   contextKey = "userID"
    RequestIDKey contextKey = "requestID"
)

func middleware(ctx context.Context) context.Context {
    // Ajouter des valeurs
    ctx = context.WithValue(ctx, UserIDKey, 12345)
    ctx = context.WithValue(ctx, RequestIDKey, "req-abc-123")
    return ctx
}

func handler(ctx context.Context) {
    // Récupérer les valeurs
    userID := ctx.Value(UserIDKey).(int)
    requestID := ctx.Value(RequestIDKey).(string)

    fmt.Printf("Traitement de la requête %s pour l'utilisateur %d\n", requestID, userID)
}

func main() {
    ctx := context.Background()
    ctx = middleware(ctx)
    handler(ctx)
}
```

### 4.2. Règles importantes pour `WithValue`

✅ **À faire :**
- Utiliser un **type personnalisé** pour les clés (évite les collisions)
- Stocker des valeurs **immutables** et **sécurisées** (auth, trace ID)
- Utiliser pour les **données de requête** (métadonnées, pas données métier)

❌ **À ne pas faire :**
- Utiliser `string` comme type de clé (risque de collision)
- Stocker des paramètres optionnels de fonction
- Remplacer des paramètres explicites

```go
// ✅ BON : type personnalisé
type userContextKey struct{}
var userKey = userContextKey{}

// ❌ MAUVAIS : string comme clé
ctx = context.WithValue(ctx, "userID", 123)

// ❌ MAUVAIS : passer des données métier importantes
ctx = context.WithValue(ctx, "userPassword", "secret")
```

### 4.3. Pattern du trace ID

```go
type traceIDKey struct{}

func WithTraceID(ctx context.Context, traceID string) context.Context {
    return context.WithValue(ctx, traceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) string {
    if val := ctx.Value(traceIDKey{}); val != nil {
        return val.(string)
    }
    return "unknown"
}

func serviceA(ctx context.Context) {
    traceID := GetTraceID(ctx)
    fmt.Printf("[%s] Service A: début\n", traceID)
    serviceB(ctx)
}

func serviceB(ctx context.Context) {
    traceID := GetTraceID(ctx)
    fmt.Printf("[%s] Service B: traitement\n", traceID)
}

func main() {
    ctx := context.Background()
    ctx = WithTraceID(ctx, "trace-001")
    serviceA(ctx)
}
```

---

## 5. `context` avec les packages standards

### 5.1. `net/http` – serveur HTTP

```go
func longHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()  // Contexte de la requête

    // Le contexte est annulé si le client ferme la connexion
    select {
    case <-ctx.Done():
        http.Error(w, "Requête annulée", http.StatusRequestTimeout)
        return
    case <-time.After(5 * time.Second):
        w.Write([]byte("Traitement terminé"))
    }
}

func main() {
    http.HandleFunc("/long", longHandler)
    http.ListenAndServe(":8080", nil)
}
```

### 5.2. `net/http` – client HTTP

```go
func requeteAvecTimeout(url string, timeout time.Duration) (*http.Response, error) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    client := &http.Client{}
    return client.Do(req)
}
```

### 5.3. `database/sql` – requêtes avec contexte

```go
func getUserByID(ctx context.Context, db *sql.DB, id int) (*User, error) {
    // Timeout de 2 secondes pour cette requête
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()

    row := db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id = ?", id)

    var user User
    err := row.Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            return nil, fmt.Errorf("requête trop lente")
        }
        return nil, err
    }
    return &user, nil
}
```

---

## 6. Erreurs et vérifications

### 6.1. Vérifier pourquoi un contexte a été annulé

```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel()

time.Sleep(2 * time.Second)

select {
case <-ctx.Done():
    err := ctx.Err()
    if errors.Is(err, context.DeadlineExceeded) {
        fmt.Println("Timeout dépassé")
    } else if errors.Is(err, context.Canceled) {
        fmt.Println("Annulation manuelle")
    }
}
```

| Erreur | Signification |
|--------|---------------|
| `context.DeadlineExceeded` | Timeout ou deadline dépassé(e) |
| `context.Canceled` | Annulation manuelle (`cancel()`) |

### 6.2. Vérifier si une deadline existe

```go
func afficherDeadline(ctx context.Context) {
    deadline, ok := ctx.Deadline()
    if ok {
        fmt.Printf("Deadline: %v (dans %v)\n", deadline, time.Until(deadline))
    } else {
        fmt.Println("Pas de deadline")
    }
}
```

---

## 7. Patterns avancés

### 7.1. Groupe de goroutines annulables

```go
func groupeAvecAnnulation(taches []func(context.Context) error) error {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    errCh := make(chan error, len(taches))
    var wg sync.WaitGroup

    for _, tache := range taches {
        wg.Add(1)
        go func(tache func(context.Context) error) {
            defer wg.Done()
            if err := tache(ctx); err != nil {
                errCh <- err
                cancel()  // Une erreur → tout le groupe s'arrête
            }
        }(tache)
    }

    // Attendre la fin
    go func() {
        wg.Wait()
        close(errCh)
    }()

    // Retourner la première erreur
    for err := range errCh {
        return err
    }
    return nil
}
```

### 7.2. Timeout personnalisé avec channel

```go
func operation(ctx context.Context) error {
    resultCh := make(chan error, 1)

    go func() {
        // Travail long
        time.Sleep(2 * time.Second)
        resultCh <- nil
    }()

    select {
    case <-ctx.Done():
        return ctx.Err()
    case err := <-resultCh:
        return err
    }
}
```

### 7.3. Contexte avec valeur de requête (exemple complet)

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "sync/atomic"
)

type contextKey string

const (
    RequestIDKey contextKey = "requestID"
    UserKey      contextKey = "user"
)

var requestCounter uint64

func requestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Générer un ID unique
        id := atomic.AddUint64(&requestCounter, 1)
        ctx := context.WithValue(r.Context(), RequestIDKey, id)
        next(w, r.WithContext(ctx))
    }
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Simuler l'authentification
        user := r.Header.Get("X-User")
        if user == "" {
            user = "anonymous"
        }
        ctx := context.WithValue(r.Context(), UserKey, user)
        next(w, r.WithContext(ctx))
    }
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        requestID := ctx.Value(RequestIDKey)
        user := ctx.Value(UserKey)

        log.Printf("[RequestID: %v] User: %v - %s %s",
            requestID, user, r.Method, r.URL.Path)

        next(w, r)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    requestID := ctx.Value(RequestIDKey)
    user := ctx.Value(UserKey)

    fmt.Fprintf(w, "Hello %s! (RequestID: %v)\n", user, requestID)
}

func main() {
    handler := requestIDMiddleware(authMiddleware(logMiddleware(helloHandler)))
    http.HandleFunc("/hello", handler)

    log.Println("Serveur démarré sur :8080")
    http.ListenAndServe(":8080", nil)
}
```

---

## 8. Pièges et bonnes pratiques

### 8.1. Piège : oublier d'appeler `cancel()`

```go
// ❌ MAUVAIS : fuite de mémoire
ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// cancel n'est pas appelée → ressources non libérées

// ✅ BON : toujours defer cancel
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

### 8.2. Piège : passer un contexte nil

```go
// ❌ MAUVAIS : panic
func process(ctx context.Context) {
    // ...
}

process(nil)  // panic: context nil

// ✅ BON : utiliser Background() si pas de contexte
process(context.Background())
```

### 8.3. Piège : stocker des mutex dans le contexte

```go
// ❌ MAUVAIS : le contexte est immuable
type Data struct {
    mu sync.Mutex
    val int
}

ctx := context.WithValue(ctx, "data", &Data{})  // Pointeur OK

// ✅ BON : les valeurs doivent être immutables ou des pointeurs
```

### 8.4. Bonnes pratiques récapitulatives

| Règle | Explication |
|-------|-------------|
| `context` est le **premier paramètre** | `func(ctx context.Context, ...)` |
| Toujours `defer cancel()` | Évite les fuites de mémoire |
| Ne pas stocker dans une structure | Passer explicitement en paramètre |
| Types personnalisés pour les clés | Évite les collisions |
| Valeurs immutables ou pointeurs | Le contexte ne doit pas muter |
| Vérifier `ctx.Err()` après `<-ctx.Done()` | Pour connaître la raison |

---

## TP final – Module 12 – Annulation de téléchargements

### Énoncé

Créez un programme qui télécharge plusieurs fichiers en parallèle avec **possibilité d'annulation** et **timeout global**.

**Fonctionnalités :**

1. **Téléchargement parallèle** : plusieurs goroutines téléchargent simultanément
2. **Timeout global** : flag `-timeout` pour limiter le temps total
3. **Annulation manuelle** : appui sur `Entrée` pour annuler
4. **Propagation** : l'annulation stoppe tous les téléchargements en cours
5. **Affichage** : progression en temps réel, arrêt propre

**Structure du programme :**

```
Téléchargement de 5 fichiers (timeout: 10s)
Appuyez sur Entrée pour annuler...

[1/5] Fichier 1: 10% 20% 30% 40% 50% 60% 70% 80% 90% 100% ✅
[2/5] Fichier 2: 5% 10% 15% 20% 25% 30% 35% 40% 45% 50%...
[3/5] Fichier 3: 0%...

⏱️ Annulation demandée par l'utilisateur !
Téléchargements interrompus.

=== Résumé ===
Succès: 1/5
Annulés: 4/5
```

### Étapes à suivre

**Étape 1 – Structure de base**

```go
package main

import (
    "context"
    "flag"
    "fmt"
    "os"
    "sync"
    "time"
)

type Result struct {
    ID       int
    Success  bool
    Error    error
    Progress int
}

func main() {
    var timeout time.Duration
    flag.DurationVar(&timeout, "timeout", 10*time.Second, "Timeout global")
    flag.Parse()

    nbFichiers := 5
    fmt.Printf("Téléchargement de %d fichiers (timeout: %v)\n", nbFichiers, timeout)
    fmt.Println("Appuyez sur Entrée pour annuler...")
}
```

**Étape 2 – Fonction de téléchargement annulable**

```go
func telecharger(ctx context.Context, id int, duree time.Duration, progress chan<- Result) {
    ticker := time.NewTicker(50 * time.Millisecond)
    defer ticker.Stop()

    for progressPct := 0; progressPct <= 100; progressPct += 5 {
        select {
        case <-ctx.Done():
            progress <- Result{ID: id, Success: false, Error: ctx.Err(), Progress: progressPct}
            return
        case <-ticker.C:
            progress <- Result{ID: id, Success: false, Progress: progressPct}
        }
        time.Sleep(duree / 20) // Simulation
    }

    progress <- Result{ID: id, Success: true, Progress: 100}
}
```

**Étape 3 – Gestion de l'annulation utilisateur**

```go
func attendAnnulation(ctx context.Context, cancel context.CancelFunc) {
    fmt.Scanln() // Attend Entrée
    fmt.Println("\n⏱️ Annulation demandée par l'utilisateur !")
    cancel()
}
```

**Étape 4 – Assemblage complet**

```go
func main() {
    var timeout time.Duration
    flag.DurationVar(&timeout, "timeout", 10*time.Second, "Timeout global")
    flag.Parse()

    nbFichiers := 5
    fmt.Printf("Téléchargement de %d fichiers (timeout: %v)\n", nbFichiers, timeout)
    fmt.Println("Appuyez sur Entrée pour annuler...\n")

    // Contexte avec timeout
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    // Channel de progression
    progress := make(chan Result, nbFichiers)
    var wg sync.WaitGroup

    // Lancement des téléchargements
    for i := 1; i <= nbFichiers; i++ {
        wg.Add(1)
        duree := time.Duration(500+100*i) * time.Millisecond
        go func(id int, duree time.Duration) {
            defer wg.Done()
            telecharger(ctx, id, duree, progress)
        }(i, duree)
    }

    // Attente de l'annulation utilisateur
    go attendAnnulation(ctx, cancel)

    // Affichage des résultats
    affichage := make(map[int]int)
    success := 0
    cancelled := 0

    go func() {
        wg.Wait()
        close(progress)
    }()

    for res := range progress {
        affichage[res.ID] = res.Progress
        if res.Progress == 100 && res.Success {
            success++
            fmt.Printf("\r✅ Fichier %d terminé\n", res.ID)
        } else if res.Error != nil {
            if cancelled == 0 {
                fmt.Println()
            }
            cancelled++
        }

        // Afficher la progression
        fmt.Print("\rProgression: ")
        for i := 1; i <= nbFichiers; i++ {
            if p, ok := affichage[i]; ok {
                fmt.Printf("[%d:%d%%] ", i, p)
            } else {
                fmt.Printf("[%d:---] ", i)
            }
        }
    }

    fmt.Printf("\n\n=== Résumé ===\n")
    fmt.Printf("✅ Succès: %d/%d\n", success, nbFichiers)
    fmt.Printf("⏹️ Annulés: %d/%d\n", cancelled, nbFichiers-success-cancelled)

    if ctx.Err() == context.DeadlineExceeded {
        fmt.Println("⚠️ Timeout global dépassé")
    }
}
```

### Solution complète

```go
package main

import (
    "context"
    "flag"
    "fmt"
    "os"
    "sync"
    "time"
)

type DownloadResult struct {
    ID       int
    Success  bool
    Error    error
    Progress int
}

func download(ctx context.Context, id int, duration time.Duration, progress chan<- DownloadResult) {
    ticker := time.NewTicker(50 * time.Millisecond)
    defer ticker.Stop()

    steps := 20
    stepDuration := duration / time.Duration(steps)

    for i := 0; i <= steps; i++ {
        progressPct := (i * 100) / steps

        select {
        case <-ctx.Done():
            progress <- DownloadResult{
                ID:       id,
                Success:  false,
                Error:    ctx.Err(),
                Progress: progressPct,
            }
            return
        case <-ticker.C:
            progress <- DownloadResult{
                ID:       id,
                Success:  false,
                Progress: progressPct,
            }
        }

        time.Sleep(stepDuration)
    }

    progress <- DownloadResult{
        ID:      id,
        Success: true,
        Progress: 100,
    }
}

func waitForCancel(ctx context.Context, cancel context.CancelFunc) {
    fmt.Scanln()
    fmt.Println("\n⚠️ Annulation demandée par l'utilisateur !")
    cancel()
}

func main() {
    var timeout time.Duration
    var numFiles int

    flag.DurationVar(&timeout, "timeout", 10*time.Second, "Timeout global")
    flag.IntVar(&numFiles, "files", 5, "Nombre de fichiers à télécharger")
    flag.Parse()

    fmt.Printf("📥 Téléchargement de %d fichiers (timeout: %v)\n", numFiles, timeout)
    fmt.Println("⌨️  Appuyez sur Entrée pour annuler...\n")

    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    progress := make(chan DownloadResult, numFiles)
    var wg sync.WaitGroup

    // Lancement des downloads
    for i := 1; i <= numFiles; i++ {
        wg.Add(1)
        // Durée aléatoire entre 500ms et 3s
        duration := time.Duration(500+200*i) * time.Millisecond
        go func(id int, d time.Duration) {
            defer wg.Done()
            download(ctx, id, d, progress)
        }(i, duration)
    }

    // Attente annulation utilisateur
    go waitForCancel(ctx, cancel)

    // Affichage des résultats
    display := make(map[int]int)
    success := 0
    cancelled := 0

    go func() {
        wg.Wait()
        close(progress)
    }()

    for res := range progress {
        display[res.ID] = res.Progress

        if res.Success {
            success++
            fmt.Printf("\n✅ Fichier %d téléchargé avec succès\n", res.ID)
        } else if res.Error != nil && cancelled == 0 {
            fmt.Println()
        }

        if res.Error != nil && !res.Success && res.Error != context.DeadlineExceeded {
            cancelled++
        }

        // Barre de progression compacte
        fmt.Print("\r📊 Progression: ")
        for i := 1; i <= numFiles; i++ {
            if p, ok := display[i]; ok {
                fmt.Printf("[%d:%3d%%] ", i, p)
            } else {
                fmt.Printf("[%d:---] ", i)
            }
        }
    }

    fmt.Printf("\n\n=== 📊 Résumé ===\n")
    fmt.Printf("✅ Succès: %d/%d\n", success, numFiles)
    fmt.Printf("⏹️  Interrompus: %d/%d\n", cancelled, numFiles-success-cancelled)

    if ctx.Err() == context.DeadlineExceeded {
        fmt.Printf("⏰ Timeout global dépassé après %v\n", timeout)
    } else if ctx.Err() == context.Canceled {
        fmt.Println("👤 Annulation par l'utilisateur")
    }
}
```

### Tests à effectuer

```bash
# Exécution normale
go run main.go

# Avec timeout court
go run main.go -timeout 2s

# Avec plus de fichiers
go run main.go -files 10

# Annulation manuelle (appuyez sur Entrée)
go run main.go
# (pendant l'exécution, appuyez sur Entrée)
```

---

## Récapitulatif des acquis

| Concept | Syntaxe | Utilisation |
|---------|---------|-------------|
| Contexte racine | `context.Background()`, `context.TODO()` | Point d'entrée |
| Annulation manuelle | `WithCancel()` | Arrêt contrôlé |
| Timeout | `WithTimeout()` | Durée maximale |
| Deadline | `WithDeadline()` | Date/heure limite |
| Valeurs | `WithValue()` | Propagation de métadonnées |
| Vérifier annulation | `ctx.Done()` | Dans une boucle ou `select` |
| Raison annulation | `ctx.Err()` | `DeadlineExceeded` ou `Canceled` |

---

## Prochain module (aperçu)

Dans le Module 13, nous découvrirons :
- Le développement web avec `net/http`
- Les routes, handlers et middleware
- La manipulation de JSON
- La création d'API REST

---

**Félicitations !** Vous maîtrisez maintenant le package `context`, un outil essentiel pour écrire des programmes Go robustes, annulables et réactifs. Vous pouvez gérer des timeouts, propager des annulations, et passer des métadonnées dans votre chaîne d'appels. 🚀
