# Module 9 – Concurrence de base

Bienvenue dans ce neuvième module ! À la fin de ce module, vous saurez écrire des programmes capables d'exécuter plusieurs tâches **simultanément** grâce aux goroutines, communiquer entre ces tâches avec les channels, et synchroniser leur terminaison avec `WaitGroup`.

---

## 0. Pourquoi la concurrence ?

### 0.1. Problème : les tâches lentes bloquent tout

Imaginez un programme qui doit télécharger 10 fichiers depuis internet.

**Sans concurrence** (exécution séquentielle) :

```go
func main() {
    debut := time.Now()

    for i := 1; i <= 10; i++ {
        telechargerFichier(i)  // Chaque téléchargement prend 1 seconde
    }

    fmt.Println("Temps total:", time.Since(debut))
    // Temps total: 10 secondes !
}
```

Chaque téléchargement attend que le précédent soit terminé. C'est **lent** et inefficace.

**Avec concurrence** :

```go
func main() {
    debut := time.Now()

    for i := 1; i <= 10; i++ {
        go telechargerFichier(i)  // Lancement en parallèle !
    }

    time.Sleep(2 * time.Second)  // On attend que tout soit fini
    fmt.Println("Temps total:", time.Since(debut))
    // Temps total: ~1-2 secondes !
}
```

Les 10 téléchargements s'exécutent **simultanément**. Le temps total est celui du téléchargement le plus long.

### 0.2. Concurrence ≠ Parallélisme

Go a une vision particulière :

| Concept | Définition | Analogie |
|---------|------------|----------|
| **Concurrence** | Gérer plusieurs tâches à la fois (structure) | Une personne qui lit plusieurs livres, un chapitre à la fois |
| **Parallélisme** | Exécuter plusieurs tâches au même moment (exécution) | Plusieurs personnes lisant chacun un livre différent |

Go excelle dans la **concurrence** : il peut gérer des milliers de tâches simultanément, même sur un processeur à un seul cœur.

---

## 1. Les goroutines (le "go" léger)

### 1.1. Qu'est-ce qu'une goroutine ?

Une **goroutine** est une fonction qui s'exécute **simultanément** avec le reste du programme.

```go
func direBonjour() {
    fmt.Println("Bonjour !")
}

func main() {
    go direBonjour()  // Lance la goroutine
    fmt.Println("Programme principal")

    // On attend un peu pour voir le message
    time.Sleep(100 * time.Millisecond)
}
```

**Sortie possible :**
```
Programme principal
Bonjour !
```

⚠️ **Piège classique :** Si on enlève `time.Sleep()`, le programme principal se termine immédiatement, et la goroutine est **tuée** avant d'avoir pu s'exécuter.

### 1.2. Syntaxe de base

```go
// 1. Avec une fonction nommée
func maFonction() { ... }
go maFonction()

// 2. Avec une fonction anonyme
go func() {
    fmt.Println("Dans une goroutine")
}()

// 3. Avec paramètres
go func(texte string) {
    fmt.Println(texte)
}("Hello depuis goroutine")
```

### 1.3. Les goroutines sont légères

| Caractéristique | Thread classique | Goroutine Go |
|-----------------|------------------|--------------|
| Taille de pile initiale | ~1 Mo | ~2 Ko |
| Création | Lente (appel système) | Très rapide |
| Nombre max typique | ~milliers | ~millions |

On peut créer **des centaines de milliers** de goroutines sans problème :

```go
func main() {
    for i := 0; i < 100000; i++ {
        go func(id int) {
            fmt.Println("Goroutine", id)
        }(i)
    }
    time.Sleep(time.Second)
}
```

---

## 2. `sync.WaitGroup` – Attendre que les goroutines finissent

`time.Sleep()` n'est pas fiable : on ne sait pas combien de temps attendre. `sync.WaitGroup` permet d'attendre **exactement** que toutes les goroutines aient terminé.

### 2.1. Fonctionnement

```go
import "sync"

func main() {
    var wg sync.WaitGroup  // 1. Créer un WaitGroup

    // 2. Ajouter une goroutine au compteur
    wg.Add(1)
    go func() {
        defer wg.Done()    // 3. Décrémenter le compteur quand terminé
        fmt.Println("Travail 1")
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("Travail 2")
    }()

    wg.Wait()  // 4. Bloquer jusqu'à ce que compteur = 0
    fmt.Println("Tous terminés")
}
```

### 2.2. Les trois méthodes essentielles

| Méthode | Effet |
|---------|-------|
| `wg.Add(n)` | Ajoute n au compteur (doit être appelé avant la goroutine) |
| `wg.Done()` | Diminue le compteur de 1 (équivalent à `Add(-1)`) |
| `wg.Wait()` | Bloque jusqu'à ce que le compteur soit 0 |

### 2.3. Pattern standard avec boucle

```go
func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)  // Important : avant le go
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Goroutine %d travaille\n", id)
            time.Sleep(100 * time.Millisecond)
        }(i)  // Passer i par copie pour éviter le piège !
    }

    wg.Wait()
    fmt.Println("Toutes les goroutines ont fini")
}
```

⚠️ **Piège classique : la variable de boucle**

```go
// ❌ MAUVAIS : capture de la variable de boucle
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // Affiche 5,5,5,5,5 (i est partagé !)
    }()
}

// ✅ BON : passer la valeur en paramètre
for i := 0; i < 5; i++ {
    go func(id int) {
        fmt.Println(id)  // Affiche 0,1,2,3,4
    }(i)
}

// ✅ BON : créer une nouvelle variable
for i := 0; i < 5; i++ {
    i := i  // Nouvelle variable locale
    go func() {
        fmt.Println(i)
    }()
}
```

---

## 3. Les channels (canaux)

Les goroutines s'exécutent de manière indépendante. Pour qu'elles **communiquent** entre elles, on utilise les **channels**.

### 3.1. Philosophie : "Don't communicate by sharing memory, share memory by communicating"

Au lieu de partager des variables (avec des verrous complexes), Go encourage à **passer des messages** via des channels.

### 3.2. Créer et utiliser un channel

```go
// Création d'un channel d'entiers
ch := make(chan int)

// Envoyer une valeur dans le channel (←)
go func() {
    ch <- 42  // Envoi
}()

// Recevoir une valeur du channel (←)
valeur := <-ch  // Réception
fmt.Println(valeur)  // 42
```

**La flèche indique le sens du flux :**
- `ch <- valeur` : envoi (on met dans le canal)
- `valeur := <-ch` : réception (on sort du canal)

### 3.3. Channels non bufferisés (synchrones)

Par défaut, un channel est **non bufferisé** : l'envoi bloque jusqu'à ce qu'une goroutine soit prête à recevoir.

```go
func main() {
    ch := make(chan int)

    go func() {
        fmt.Println("Envoi en cours...")
        ch <- 42        // Bloque jusqu'à réception
        fmt.Println("Envoi terminé")
    }()

    time.Sleep(100 * time.Millisecond)
    fmt.Println("Réception en cours...")
    valeur := <-ch     // Débloque l'envoi
    fmt.Println("Reçu:", valeur)
}
```

**Sortie :**
```
Envoi en cours...
Réception en cours...
Reçu: 42
Envoi terminé
```

C'est un mécanisme de **synchronisation** parfait.

### 3.4. Channels bufferisés

On peut créer un channel avec un **buffer** (tampon) :

```go
// Channel bufferisé : peut contenir jusqu'à 3 messages
ch := make(chan int, 3)

// Envoi : ne bloque que si le buffer est plein
ch <- 1
ch <- 2
ch <- 3
// ch <- 4  // Bloque ici (buffer plein)

// Réception : ne bloque que si le buffer est vide
fmt.Println(<-ch)  // 1
fmt.Println(<-ch)  // 2
fmt.Println(<-ch)  // 3
```

| Type | Comportement |
|------|--------------|
| Non bufferisé (`make(chan T)`) | L'envoi bloque jusqu'à ce qu'une goroutine reçoive |
| Bufferisé (`make(chan T, n)`) | Envoi bloque seulement si buffer plein |

💡 **Règle :** Commencez toujours par des channels non bufferisés. Ils sont plus simples à comprendre et forcent une bonne synchronisation.

---

## 4. Exemples concrets

### 4.1. Producteur / Consommateur

```go
func producteur(ch chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i <= 5; i++ {
        fmt.Printf("Producteur: envoie %d\n", i)
        ch <- i
        time.Sleep(100 * time.Millisecond)
    }
    close(ch)  // Fermeture du channel (important !)
}

func consommateur(ch <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for valeur := range ch {  // range s'arrête quand channel fermé
        fmt.Printf("Consommateur: reçu %d\n", valeur)
        time.Sleep(200 * time.Millisecond)
    }
}

func main() {
    ch := make(chan int)
    var wg sync.WaitGroup

    wg.Add(2)
    go producteur(ch, &wg)
    go consommateur(ch, &wg)

    wg.Wait()
}
```

**Notation de direction :**
- `chan<- int` : channel **envoi uniquement**
- `<-chan int` : channel **réception uniquement**
- `chan int` : **envoi et réception**

### 4.2. Calcul parallèle

```go
// Calcul du carré d'un nombre dans une goroutine
func carre(n int, ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    resultat := n * n
    ch <- resultat
}

func main() {
    nombres := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    ch := make(chan int, len(nombres))  // Buffer pour ne pas bloquer
    var wg sync.WaitGroup

    // Lancement des goroutines
    for _, n := range nombres {
        wg.Add(1)
        go carre(n, ch, &wg)
    }

    // Attente dans une goroutine séparée
    go func() {
        wg.Wait()
        close(ch)  // Fermer le channel quand tous ont fini
    }()

    // Collecte des résultats
    somme := 0
    for resultat := range ch {
        somme += resultat
        fmt.Printf("Carré reçu: %d, somme partielle: %d\n", resultat, somme)
    }

    fmt.Printf("Somme des carrés: %d\n", somme)
}
```

---

## 5. Le `select` – attendre sur plusieurs channels

`select` permet d'attendre sur **plusieurs channels** à la fois. Il bloque jusqu'à ce qu'un des channels soit prêt.

```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Message de ch1"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Message de ch2"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Reçu de ch1:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Reçu de ch2:", msg2)
        }
    }
}
// Output:
// Reçu de ch1: Message de ch1  (après 1s)
// Reçu de ch2: Message de ch2  (après 2s)
```

### 5.1. `select` avec timeout

```go
func travailLong() chan string {
    ch := make(chan string)
    go func() {
        time.Sleep(3 * time.Second)
        ch <- "Résultat"
    }()
    return ch
}

func main() {
    ch := travailLong()

    select {
    case resultat := <-ch:
        fmt.Println("Succès:", resultat)
    case <-time.After(2 * time.Second):
        fmt.Println("Timeout ! L'opération a pris trop de temps")
    }
}
// Output: Timeout ! L'opération a pris trop de temps
```

### 5.2. `select` avec `default` (non bloquant)

```go
ch := make(chan int)

select {
case val := <-ch:
    fmt.Println("Reçu:", val)
default:
    fmt.Println("Rien à recevoir, on continue")
}
```

---

## 6. Fermeture des channels

### 6.1. Pourquoi fermer ?

Fermer un channel indique aux récepteurs qu'il n'y aura **plus de données** à venir.

```go
ch := make(chan int)

go func() {
    for i := 0; i < 3; i++ {
        ch <- i
    }
    close(ch)  // Indique la fin
}()

for val := range ch {  // Boucle jusqu'à fermeture
    fmt.Println(val)
}
// Output: 0, 1, 2
```

### 6.2. Vérifier si un channel est fermé

```go
valeur, ouvert := <-ch
if ouvert {
    fmt.Println("Channel ouvert, valeur reçue:", valeur)
} else {
    fmt.Println("Channel fermé")
}
```

### 6.3. Règles importantes

| Action sur channel fermé | Résultat |
|--------------------------|----------|
| `v, ok := <-ch` | `v` = valeur zéro, `ok` = false |
| `ch <- x` | **Panic** ! |
| `close(ch)` | **Panic** ! |

✅ **Règle d'or :** C'est l'**expéditeur** qui ferme le channel, jamais le récepteur.

---

## TP final – Module 9 – Téléchargement parallèle

### Énoncé

Créez un programme qui télécharge plusieurs fichiers "simulés" en parallèle, avec affichage de la progression.

**Fonctionnalités :**

1. **Simulation de téléchargement** : fonction `telecharger(id int, duree time.Duration)`
2. **Téléchargement parallèle** : utilisez des goroutines
3. **Collecte des résultats** : utilisez un channel pour récupérer les statuts
4. **Affichage en temps réel** : affichez chaque fin de téléchargement
5. **Barre de progression** : affichez le nombre de fichiers restants

**Structure du programme :**

```
Téléchargement de 5 fichiers...
Fichier 2 téléchargé en 800ms
Fichier 4 téléchargé en 900ms
Fichier 1 téléchargé en 1.0s
Fichier 3 téléchargé en 1.2s
Fichier 5 téléchargé en 1.5s

=== Résumé ===
Succès: 5/5
Temps total: 1.5s
```

### Étapes à suivre

**Étape 1 – Créer le fichier**

Créez un dossier `module9` et un fichier `download.go`.

**Étape 2 – Structure de base**

```go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func main() {
    // Votre code ici
}
```

**Étape 3 – Fonction de téléchargement simulé**

```go
func telecharger(id int, duree time.Duration, result chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()

    start := time.Now()
    time.Sleep(duree)  // Simulation du téléchargement
    elapsed := time.Since(start)

    result <- fmt.Sprintf("Fichier %d téléchargé en %v", id, elapsed)
}
```

**Étape 4 – Dans main()**

```go
func main() {
    rand.Seed(time.Now().UnixNano())

    nbFichiers := 5
    result := make(chan string, nbFichiers)  // Channel bufferisé
    var wg sync.WaitGroup

    fmt.Printf("Téléchargement de %d fichiers...\n\n", nbFichiers)

    startTotal := time.Now()

    // Lancement des téléchargements
    for i := 1; i <= nbFichiers; i++ {
        wg.Add(1)
        duree := time.Duration(rand.Intn(1500)+500) * time.Millisecond
        go telecharger(i, duree, result, &wg)
    }

    // Attente dans une goroutine séparée
    go func() {
        wg.Wait()
        close(result)
    }()

    // Affichage des résultats en temps réel
    for msg := range result {
        fmt.Println(msg)
    }

    fmt.Printf("\n=== Résumé ===\n")
    fmt.Printf("Succès: %d/%d\n", nbFichiers, nbFichiers)
    fmt.Printf("Temps total: %v\n", time.Since(startTotal))
}
```

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

### Proposition de correction

```go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// Simulation de téléchargement
func telecharger(id int, duree time.Duration, result chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()

    start := time.Now()
    time.Sleep(duree)
    elapsed := time.Since(start)

    result <- fmt.Sprintf("Fichier %d téléchargé en %v", id, elapsed.Round(time.Millisecond))
}

func main() {
    rand.Seed(time.Now().UnixNano())

    const nbFichiers = 5
    result := make(chan string, nbFichiers)
    var wg sync.WaitGroup

    fmt.Printf("📥 Téléchargement de %d fichiers...\n\n", nbFichiers)

    startTotal := time.Now()

    // Lancement des téléchargements
    for i := 1; i <= nbFichiers; i++ {
        wg.Add(1)
        // Durée aléatoire entre 500ms et 2000ms
        duree := time.Duration(rand.Intn(1500)+500) * time.Millisecond
        go telecharger(i, duree, result, &wg)
    }

    // Goroutine qui attend la fin et ferme le channel
    go func() {
        wg.Wait()
        close(result)
    }()

    // Compteur pour la progression
    termines := 0
    for msg := range result {
        termines++
        fmt.Printf("[%d/%d] %s\n", termines, nbFichiers, msg)
    }

    fmt.Printf("\n=== 📊 Résumé ===\n")
    fmt.Printf("✅ Succès: %d/%d\n", termines, nbFichiers)
    fmt.Printf("⏱️  Temps total: %v\n", time.Since(startTotal).Round(time.Millisecond))
}
```

**Exemple de sortie :**
```
📥 Téléchargement de 5 fichiers...

[1/5] Fichier 3 téléchargé en 623ms
[2/5] Fichier 1 téléchargé en 847ms
[3/5] Fichier 5 téléchargé en 1.045s
[4/5] Fichier 2 téléchargé en 1.234s
[5/5] Fichier 4 téléchargé en 1.456s

=== 📊 Résumé ===
✅ Succès: 5/5
⏱️  Temps total: 1.456s
```

### Version améliorée avec gestion des erreurs

```go
type Resultat struct {
    ID     int
    Duree  time.Duration
    Erreur error
}

func telechargerAvecErreur(id int, duree time.Duration) Resultat {
    start := time.Now()

    // 10% de chance d'échec
    if rand.Intn(100) < 10 {
        time.Sleep(duree / 2)
        return Resultat{ID: id, Duree: time.Since(start), Erreur: fmt.Errorf("échec réseau")}
    }

    time.Sleep(duree)
    return Resultat{ID: id, Duree: time.Since(start), Erreur: nil}
}

func main() {
    rand.Seed(time.Now().UnixNano())

    const nbFichiers = 10
    result := make(chan Resultat, nbFichiers)
    var wg sync.WaitGroup

    fmt.Printf("📥 Téléchargement de %d fichiers...\n\n", nbFichiers)

    startTotal := time.Now()

    for i := 1; i <= nbFichiers; i++ {
        wg.Add(1)
        duree := time.Duration(rand.Intn(1000)+500) * time.Millisecond
        go func(id int, duree time.Duration) {
            defer wg.Done()
            result <- telechargerAvecErreur(id, duree)
        }(i, duree)
    }

    go func() {
        wg.Wait()
        close(result)
    }()

    reussis := 0
    echecs := 0

    for r := range result {
        if r.Erreur != nil {
            echecs++
            fmt.Printf("❌ Fichier %d: %v\n", r.ID, r.Erreur)
        } else {
            reussis++
            fmt.Printf("✅ Fichier %d téléchargé en %v\n", r.ID, r.Duree.Round(time.Millisecond))
        }
    }

    fmt.Printf("\n=== 📊 Résumé ===\n")
    fmt.Printf("✅ Succès: %d/%d\n", reussis, nbFichiers)
    fmt.Printf("❌ Échecs: %d/%d\n", echecs, nbFichiers)
    fmt.Printf("⏱️  Temps total: %v\n", time.Since(startTotal).Round(time.Millisecond))
}
```

### Tests à effectuer

```bash
# Exécution simple
go run download.go

# Avec plus de fichiers
# Modifiez nbFichiers := 20

# Observez le temps total vs temps cumulé
# Avec 5 fichiers de 1s chacun : séquentiel = 5s, parallèle ≈ 1s
```

### Pour aller plus loin (optionnel)

**1. Limiter le parallélisme avec un channel de jeton**

```go
// Limite à 3 téléchargements simultanés
limiteur := make(chan struct{}, 3)

for i := 1; i <= 20; i++ {
    wg.Add(1)
    go func(id int) {
        limiteur <- struct{}{}        // Prendre un jeton
        defer func() { <-limiteur }() // Rendre le jeton

        // Téléchargement...
    }(i)
}
```

**2. Barre de progression avec `\r`**

```go
termines := 0
for range result {
    termines++
    fmt.Printf("\r📥 Progression: %d/%d", termines, nbFichiers)
}
fmt.Println()
```

**3. Channel de progression séparé**

```go
progress := make(chan int)

go func() {
    for p := range progress {
        fmt.Printf("\rProgression: %d%%", p)
    }
}()

// Dans la goroutine
progress <- 50  // Envoyer la progression
```

---

## Récapitulatif des acquis

À la fin de ce module, vous savez :

| Concept | Syntaxe | Utilité |
|---------|---------|---------|
| Goroutine | `go fonction()` | Exécution parallèle légère |
| WaitGroup | `wg.Add()`, `wg.Done()`, `wg.Wait()` | Attendre la fin des goroutines |
| Channel | `make(chan T)`, `ch <-`, `<-ch` | Communication entre goroutines |
| Channel bufferisé | `make(chan T, n)` | Éviter le blocage d'envoi |
| Select | `select { case ... }` | Attendre sur plusieurs channels |
| Fermeture | `close(ch)` | Signaler la fin des données |

---

## Prochain module (aperçu)

Dans le Module 10, nous découvrirons :
- La lecture et écriture de fichiers
- Les entrées/sorties avec `os` et `bufio`
- Les arguments de ligne de commande (`os.Args`)
- La manipulation de fichiers CSV et JSON

---

**Félicitations !** Vous savez maintenant écrire des programmes Go concurrents, capables d'exécuter plusieurs tâches en parallèle de manière simple et efficace. C'est l'une des plus grandes forces de Go ! 🚀
