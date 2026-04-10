# Module 6 – Packages et modularité

Bienvenue dans ce sixième module ! À la fin de ce module, vous saurez organiser votre code en plusieurs fichiers, créer vos propres packages réutilisables, comprendre l'exportation avec les majuscules, et structurer un projet Go professionnel.

---

## 0. Pourquoi les packages ?

Jusqu'ici, tout votre code était dans un seul fichier `main.go`. Pour un petit programme, ça va. Mais dès que votre projet grandit :

- **Difficile à naviguer** (1000+ lignes dans un fichier)
- **Difficile à réutiliser** (copier/coller de code entre projets)
- **Difficile à tester** (tout est mélangé)
- **Difficile à partager** (pas de frontière claire entre les fonctionnalités)

Les packages résolvent ces problèmes :

```
❌ Sans packages (tout dans main.go) :
main.go (2000 lignes)
├── fonctions métier
├── gestion base de données
├── affichage UI
├── lecture fichiers
└── calculs divers

✅ Avec packages (code organisé) :
main.go (50 lignes)
├── contact/
│   └── contact.go (gestion des contacts)
├── storage/
│   └── storage.go (sauvegarde fichiers)
├── ui/
│   └── display.go (affichage)
└── utils/
    └── validator.go (validation)
```

---

## 1. Qu'est-ce qu'un package ?

En Go, un **package** est un ensemble de fichiers Go regroupés dans un même dossier. Tous les fichiers d'un même dossier partagent le même nom de package.

### 1.1. Structure d'un package

```
monprojet/
├── go.mod
├── main.go          (package main)
├── contact/
│   └── contact.go   (package contact)
└── utils/
    └── string.go    (package utils)
```

### 1.2. Le package `main` est spécial

```go
// main.go - POINT D'ENTRÉE OBLIGATOIRE
package main

import "fmt"

func main() {
    fmt.Println("Le programme commence ici")
}
```

- `package main` est le seul package qui produit un **exécutable**
- Il doit contenir une fonction `main()`
- Les autres packages produisent du **code réutilisable** (bibliothèques)

### 1.3. Créer un package réutilisable

```go
// contact/contact.go
package contact  // Pas "main" !

type Contact struct {
    Name  string
    Phone string
}

func New(name, phone string) Contact {
    return Contact{Name: name, Phone: phone}
}

func (c Contact) Display() string {
    return c.Name + " - " + c.Phone
}
```

---

## 2. Exporter avec la majuscule

En Go, la **casse** détermine ce qui est visible depuis l'extérieur du package.

### 2.1. Règle fondamentale

| Casse | Visibilité | Exemple |
|-------|------------|---------|
| **Majuscule** (public) | Visible depuis l'extérieur | `Contact`, `New()`, `Display()` |
| **Minuscule** (private) | Visible seulement dans le package | `formatPhone()`, `validateName()` |

```go
// contact/contact.go
package contact

// EXPORTÉ (visible par main.go)
type Contact struct {     // Majuscule → exporté
    Name  string          // Majuscule → exporté
    Phone string          // Majuscule → exporté
}

// EXPORTÉ
func New(name, phone string) Contact {
    return Contact{Name: name, Phone: phone}
}

// EXPORTÉ
func (c Contact) Display() string {
    return c.Name + " - " + formatPhone(c.Phone)  // appel interne
}

// NON EXPORTÉ (visible seulement dans contact/)
func formatPhone(phone string) string {
    // nettoyage du numéro
    return strings.ReplaceAll(phone, " ", "")
}
```

### 2.2. Utiliser un package exporté

```go
// main.go
package main

import (
    "fmt"
    "monprojet/contact"  // import du package contact
)

func main() {
    // ✅ Visible (majuscule)
    c := contact.New("Alice", "06 12 34 56 78")
    fmt.Println(c.Display())  // Alice - 0612345678

    // ❌ NON visible (minuscule)
    // contact.formatPhone("0612345678")  // Erreur de compilation !
}
```

💡 **Piège fréquent :** Les champs d'une structure doivent aussi être en majuscule pour être accessibles depuis l'extérieur !

```go
// ❌ Ne fonctionne PAS
type Contact struct {
    name  string  // minuscule → inaccessible
    phone string  // minuscule → inaccessible
}

// Dans main.go :
c := contact.Contact{Name: "Alice"}  // Erreur : unknown field Name
```

---

## 3. Importer des packages

### 3.1. Syntaxe d'importation

```go
import "fmt"                    // Package standard
import "monprojet/contact"      // Package local
import "github.com/user/pkg"    // Package tiers
```

**Version recommandée (multi-imports) :**

```go
import (
    "fmt"
    "strings"
    "monprojet/contact"
    "monprojet/utils"
)
```

### 3.2. Alias d'importation

Quand deux packages ont le même nom ou pour raccourcir :

```go
import (
    "fmt"
    myfmt "mypackage/fmt"        // Évite le conflit
    c "monprojet/contact"        // Alias court
)

func main() {
    c.New("Alice", "0612345678") // Utilisation avec alias
}
```

### 3.3. Import "dot" (à éviter)

```go
import . "fmt"  // Importe tout dans l'espace global

func main() {
    Println("Hello")  // Pas besoin de fmt.Println
}
```

⚠️ **Déconseillé** : Cela pollue l'espace de noms et rend le code moins lisible.

### 3.4. Import "blank" (pour les effets de bord)

```go
import _ "image/png"  // Enregistre le format PNG sans l'utiliser
```

Utile pour les packages qui ont un effet de bord à l'initialisation (ex: drivers de base de données).

---

## 4. Organisation d'un projet Go

### 4.1. Structure simple (petit projet)

```
monprojet/
├── go.mod
├── main.go
├── contact/
│   └── contact.go
└── utils/
    └── utils.go
```

### 4.2. Structure professionnelle (moyen/gros projet)

```
monprojet/
├── go.mod
├── go.sum
├── cmd/                    # Exécutables
│   ├── api/
│   │   └── main.go
│   └── cli/
│       └── main.go
├── internal/               # Code privé (non importable)
│   └── auth/
│       └── auth.go
├── pkg/                    # Code public (importable)
│   ├── contact/
│   │   └── contact.go
│   └── storage/
│       └── storage.go
└── test/                   # Tests externes
    └── integration_test.go
```

| Dossier | Usage |
|---------|-------|
| `cmd/` | Points d'entrée (exécutables) |
| `internal/` | Code **privé** au projet (Go l'interdit en import externe) |
| `pkg/` | Code **public** réutilisable par d'autres projets |
| `test/` | Tests d'intégration |

💡 **Règle d'or :** `internal` est un mot-clé spécial. Personne en dehors de votre module ne peut importer un package dans `internal/`.

### 4.3. Le fichier `go.mod`

Le fichier `go.mod` définit votre module (l'équivalent du `package.json` en Node.js ou `requirements.txt` en Python).

```bash
# Initialiser un module
go mod init monprojet
```

```go
// go.mod
module monprojet

go 1.22.0

require (
    github.com/gorilla/mux v1.8.0
)
```

| Commande | Utilité |
|----------|---------|
| `go mod init <nom>` | Crée un nouveau module |
| `go get <package>` | Ajoute une dépendance |
| `go mod tidy` | Nettoie les dépendances inutilisées |
| `go mod vendor` | Copie les dépendances dans `vendor/` |

---

## 5. Packages standards couramment utilisés

| Package | Utilité |
|---------|---------|
| `fmt` | Entrées/sorties formatées |
| `os` | Fonctions système (fichiers, env, arguments) |
| `io` | Interfaces d'entrée/sortie |
| `bufio` | Lecture/écriture bufferisée |
| `strings` | Manipulation de chaînes |
| `strconv` | Conversion string ↔ nombres |
| `errors` | Création et manipulation d'erreurs |
| `time` | Dates, durées, temporisation |
| `math` | Fonctions mathématiques |
| `math/rand` | Nombres aléatoires |
| `encoding/json` | JSON (sérialisation/désérialisation) |
| `net/http` | Serveur et client HTTP |
| `sync` | Concurrence (Mutex, WaitGroup) |
| `context` | Contexte (timeout, annulation) |
| `testing` | Tests unitaires |

---

## TP final – Module 6 – Création d'un package mathutil

### Énoncé

Créez un package réutilisable `mathutil` qui contient des fonctions mathématiques, puis utilisez-le dans `main.go`.

**Partie 1 – Créer le package `mathutil`**

Créez un dossier `mathutil` avec un fichier `mathutil.go` contenant :

1. **Fonction `Addition(a, b int) int`** qui retourne la somme
2. **Fonction `Multiplication(a, b int) int`** qui retourne le produit
3. **Fonction `Factorielle(n int) int`** qui retourne n! (0! = 1)
4. **Fonction `EstPair(n int) bool`** qui retourne true si n est pair

**Partie 2 – Utiliser le package**

Dans `main.go`, importez `mathutil` et :

1. Affichez `Addition(5, 3)` → 8
2. Affichez `Multiplication(4, 2)` → 8
3. Affichez `Factorielle(5)` → 120
4. Affichez `EstPair(7)` → false

### Exemple de sortie attendue :

```
=== Tests du package mathutil ===

Addition(5, 3) = 8
Multiplication(4, 2) = 8
Factorielle(5) = 120
EstPair(7) = false
```

### Structure attendue

```
module6/
├── go.mod
├── main.go
└── mathutil/
    └── mathutil.go
```

### Étapes à suivre

**Étape 1 – Créer la structure du projet**

```bash
mkdir -p module6/mathutil
cd module6
go mod init module6
```

**Étape 2 – Créer le package `mathutil`**

```go
// mathutil/mathutil.go
package mathutil

// Addition retourne la somme de a et b
func Addition(a, b int) int {
    return a + b
}

// Multiplication retourne le produit de a et b
func Multiplication(a, b int) int {
    return a * b
}

// Factorielle retourne n! (produit de 1 à n)
func Factorielle(n int) int {
    if n <= 1 {
        return 1
    }
    return n * Factorielle(n-1)
}

// EstPair retourne true si n est pair
func EstPair(n int) bool {
    return n%2 == 0
}
```

**Étape 3 – Créer le main**

```go
// main.go
package main

import (
    "fmt"
    "module6/mathutil"
)

func main() {
    fmt.Println("=== Tests du package mathutil ===\n")

    fmt.Printf("Addition(5, 3) = %d\n", mathutil.Addition(5, 3))
    fmt.Printf("Multiplication(4, 2) = %d\n", mathutil.Multiplication(4, 2))
    fmt.Printf("Factorielle(5) = %d\n", mathutil.Factorielle(5))
    fmt.Printf("EstPair(7) = %t\n", mathutil.EstPair(7))
}
```

**Étape 4 – Exécuter**

```bash
go run main.go
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

**Fichier `mathutil/mathutil.go` :**

```go
package mathutil

// Addition retourne la somme de a et b
func Addition(a, b int) int {
    return a + b
}

// Multiplication retourne le produit de a et b
func Multiplication(a, b int) int {
    return a * b
}

// Factorielle retourne n! (produit de 1 à n)
// Utilisation d'une boucle plutôt que récursion (plus efficace)
func Factorielle(n int) int {
    if n < 0 {
        return -1 // Erreur pour nombre négatif
    }
    resultat := 1
    for i := 2; i <= n; i++ {
        resultat *= i
    }
    return resultat
}

// EstPair retourne true si n est pair
func EstPair(n int) bool {
    return n%2 == 0
}
```

**Fichier `main.go` :**

```go
package main

import (
    "fmt"
    "module6/mathutil"
)

func main() {
    fmt.Println("=== Tests du package mathutil ===\n")

    fmt.Printf("Addition(5, 3) = %d\n", mathutil.Addition(5, 3))
    fmt.Printf("Multiplication(4, 2) = %d\n", mathutil.Multiplication(4, 2))
    fmt.Printf("Factorielle(5) = %d\n", mathutil.Factorielle(5))
    fmt.Printf("EstPair(7) = %t\n", mathutil.EstPair(7))

    // Tests supplémentaires
    fmt.Println("\n--- Tests supplémentaires ---")
    fmt.Printf("Factorielle(0) = %d\n", mathutil.Factorielle(0)) // 1
    fmt.Printf("EstPair(42) = %t\n", mathutil.EstPair(42))       // true
}
```

**Fichier `go.mod` (généré automatiquement) :**

```go
module module6

go 1.22.0
```

### Version alternative avec tests d'erreur

```go
// mathutil/mathutil.go
package mathutil

import "errors"

// Addition retourne la somme de a et b
func Addition(a, b int) int {
    return a + b
}

// Multiplication retourne le produit de a et b
func Multiplication(a, b int) int {
    return a * b
}

// Factorielle retourne n! et une erreur si n est négatif
func Factorielle(n int) (int, error) {
    if n < 0 {
        return 0, errors.New("factorielle d'un nombre négatif non définie")
    }
    resultat := 1
    for i := 2; i <= n; i++ {
        resultat *= i
    }
    return resultat, nil
}

// EstPair retourne true si n est pair
func EstPair(n int) bool {
    return n%2 == 0
}
```

```go
// main.go (version avec gestion d'erreur)
package main

import (
    "fmt"
    "module6/mathutil"
)

func main() {
    fmt.Println("=== Tests du package mathutil ===\n")

    fmt.Printf("Addition(5, 3) = %d\n", mathutil.Addition(5, 3))
    fmt.Printf("Multiplication(4, 2) = %d\n", mathutil.Multiplication(4, 2))

    // Factorielle avec gestion d'erreur
    if fact, err := mathutil.Factorielle(5); err != nil {
        fmt.Println("Erreur:", err)
    } else {
        fmt.Printf("Factorielle(5) = %d\n", fact)
    }

    fmt.Printf("EstPair(7) = %t\n", mathutil.EstPair(7))
}
```

### Tests à effectuer

```bash
# Exécution directe
go run main.go

# Formatage du code
go fmt ./...

# Vérification des erreurs potentielles
go vet ./...

# Construction de l'exécutable
go build -o mathapp main.go
./mathapp
```

### Pour aller plus loin (optionnel)

**1. Ajouter plus de fonctions au package**

```go
// mathutil/mathutil.go

// Moyenne calcule la moyenne d'une slice d'entiers
func Moyenne(nombres []int) float64 {
    if len(nombres) == 0 {
        return 0
    }
    somme := 0
    for _, n := range nombres {
        somme += n
    }
    return float64(somme) / float64(len(nombres))
}

// Min retourne la plus petite valeur d'une slice
func Min(nombres []int) int {
    if len(nombres) == 0 {
        return 0
    }
    min := nombres[0]
    for _, n := range nombres {
        if n < min {
            min = n
        }
    }
    return min
}
```

**2. Créer un package `stringutil`**

```go
// stringutil/stringutil.go
package stringutil

import "strings"

// Majuscule retourne la chaîne en majuscules
func Majuscule(s string) string {
    return strings.ToUpper(s)
}

// Minuscule retourne la chaîne en minuscules
func Minuscule(s string) string {
    return strings.ToLower(s)
}

// Inverse retourne la chaîne inversée
func Inverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

**3. Utiliser un package tiers**

```bash
go get github.com/fatih/color
```

```go
package main

import (
    "github.com/fatih/color"
    "module6/mathutil"
)

func main() {
    color.Cyan("=== Tests du package mathutil ===\n")

    result := mathutil.Addition(5, 3)
    color.Green("Addition(5, 3) = %d", result)
}
```

---

## Récapitulatif des acquis

À la fin de ce module, vous savez :

- ✅ Comprendre pourquoi organiser son code en packages
- ✅ Créer vos propres packages réutilisables
- ✅ Exporter des fonctions/types avec la **majuscule**
- ✅ Importer des packages avec ou sans alias
- ✅ Structurer un projet Go (simple ou professionnel)
- ✅ Utiliser `go mod init` pour créer un module
- ✅ Connaître les packages standards essentiels

---

## Prochain module (aperçu)

Dans le Module 7, nous découvrirons :
- La gestion avancée des erreurs (`panic`, `recover`, `defer`)
- Le logging avec le package `log`
- Les bonnes pratiques pour des programmes robustes

---

**Félicitations !** Vous savez maintenant organiser votre code comme un vrai projet Go professionnel. 🚀
