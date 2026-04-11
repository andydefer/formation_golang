# Module 7 – Gestion des erreurs (niveau 1)

Bienvenue dans ce septième module ! À la fin de ce module, vous saurez gérer les erreurs comme un vrai professionnel Go, comprendre `panic`, `recover`, `defer`, et mettre en place du logging structuré.

---

## 0. Rappel : Les erreurs en Go

En Go, il n'y a **pas d'exceptions** comme en Python, Java ou C#. Go utilise une approche explicite :

```go
// Python (exceptions)
try:
    result = division(10, 0)
except ZeroDivisionError:
    print("Erreur !")

// Go (erreur explicite)
result, err := division(10, 0)
if err != nil {
    fmt.Println("Erreur !")
}
```

**Philosophie Go :** Les erreurs sont des valeurs comme les autres. On les traite immédiatement, là où elles se produisent.

---

## 1. Le pattern `if err != nil`

C'est le pattern le plus fréquent en Go. Vous allez l'écrire des centaines de fois.

### 1.1. Structure de base

```go
package main

import (
    "errors"
    "fmt"
)

// Fonction qui retourne une erreur
func diviser(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division par zéro impossible")
    }
    return a / b, nil
}

func main() {
    resultat, err := diviser(10, 0)
    if err != nil {
        fmt.Println("Erreur :", err)
        return  // On sort de la fonction
    }
    fmt.Println("Résultat :", resultat)
}
```

### 1.2. Vérification inline (idiomatique)

```go
// ✅ Style Go idiomatique
if resultat, err := diviser(10, 0); err != nil {
    fmt.Println("Erreur :", err)
    return
} else {
    fmt.Println("Résultat :", resultat)
}
```

### 1.3. Regrouper les vérifications

```go
func processFichier(nom string) error {
    data, err := lireFichier(nom)
    if err != nil {
        return fmt.Errorf("lecture échouée: %w", err)
    }

    resultat, err := traiterData(data)
    if err != nil {
        return fmt.Errorf("traitement échoué: %w", err)
    }

    err = sauvegarder(resultat)
    if err != nil {
        return fmt.Errorf("sauvegarde échouée: %w", err)
    }

    return nil
}
```

---

## 2. Créer des erreurs

### 2.1. `errors.New()` – erreur simple

```go
import "errors"

var ErrDivisionParZero = errors.New("division par zéro impossible")
var ErrFichierIntrouvable = errors.New("fichier non trouvé")

func diviser(a, b float64) (float64, error) {
    if b == 0 {
        return 0, ErrDivisionParZero
    }
    return a / b, nil
}
```

**Bonnes pratiques :**
- Nommez les variables d'erreur avec le préfixe `Err`
- Déclarez les erreurs communes au niveau du package
- Commentez chaque erreur exportée

```go
// Package mathutil fournit des fonctions mathématiques
package mathutil

// ErrDivisionParZero est retourné quand on divise par zéro
var ErrDivisionParZero = errors.New("mathutil: division par zéro")
```

### 2.2. `fmt.Errorf()` – erreur formatée

```go
import "fmt"

func lireFichier(nom string) ([]byte, error) {
    if nom == "" {
        return nil, fmt.Errorf("nom de fichier invalide: %s est vide", nom)
    }
    // ...
}
```

### 2.3. `%w` – wrapping d'erreur (Go 1.13+)

Le wrapping permet d'**emballer** une erreur dans une autre, tout en conservant l'erreur originale.

```go
func ouvrirBaseDeDonnees(path string) error {
    err := ouvrirFichier(path)
    if err != nil {
        // %w = wrap (conserve l'erreur originale)
        return fmt.Errorf("ouverture base de données échouée: %w", err)
    }
    return nil
}

func main() {
    err := ouvrirBaseDeDonnees("/data/db")
    if err != nil {
        // Vérifier si l'erreur contient ErrFichierIntrouvable
        if errors.Is(err, ErrFichierIntrouvable) {
            fmt.Println("Le fichier n'existe pas")
        }
        fmt.Println("Erreur:", err)
    }
}
```

### 2.4. Erreurs personnalisées avec `struct`

Pour les erreurs complexes (avec contexte) :

```go
type ValidationError struct {
    Champ   string
    Valeur  interface{}
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation échouée pour %s (valeur: %v): %s",
        e.Champ, e.Valeur, e.Message)
}

func validerAge(age int) error {
    if age < 0 {
        return ValidationError{
            Champ:   "age",
            Valeur:  age,
            Message: "l'âge ne peut pas être négatif",
        }
    }
    if age > 150 {
        return ValidationError{
            Champ:   "age",
            Valeur:  age,
            Message: "l'âge est trop grand",
        }
    }
    return nil
}
```

---

## 3. Les fonctions essentielles du package `errors`

### 3.1. `errors.New()` – créer une erreur

```go
var ErrNotFound = errors.New("not found")
```

### 3.2. `errors.Is()` – vérifier le type d'erreur

```go
if errors.Is(err, ErrNotFound) {
    // L'erreur EST EXACTEMENT ErrNotFound (ou l'enveloppe)
}
```

### 3.3. `errors.As()` – extraire une erreur personnalisée

```go
var valErr ValidationError
if errors.As(err, &valErr) {
    fmt.Printf("Champ en erreur: %s\n", valErr.Champ)
}
```

### 3.4. `errors.Unwrap()` – déballer une erreur

```go
err := fmt.Errorf("problème: %w", ErrNotFound)
sousErreur := errors.Unwrap(err)
fmt.Println(sousErreur == ErrNotFound) // true
```

---

## 4. `defer` – exécution différée

`defer` permet de programmer une fonction pour qu'elle s'exécute **à la fin** de la fonction courante (même en cas d'erreur ou de `panic`).

### 4.1. Fermeture de ressources

```go
func lireFichier(nom string) (string, error) {
    f, err := os.Open(nom)
    if err != nil {
        return "", err
    }
    defer f.Close()  // Exécuté QUAND MÊME à la sortie de la fonction

    // Lecture du fichier...
    data := make([]byte, 100)
    f.Read(data)
    return string(data), nil
}
```

### 4.2. Ordre d'exécution (LIFO)

```go
func main() {
    defer fmt.Println("1")  // S'exécute en 3ème
    defer fmt.Println("2")  // S'exécute en 2ème
    defer fmt.Println("3")  // S'exécute en 1er
    fmt.Println("début")
}
// Output:
// début
// 3
// 2
// 1
```

### 4.3. Piège classique : évaluation des arguments

Les arguments d'un `defer` sont évalués **immédiatement**, pas quand le defer s'exécute !

```go
// ❌ Piège
func mauvaisDefer() {
    x := 1
    defer fmt.Println(x)  // x = 1 est stocké tout de suite
    x = 2
}
// Output: 1 (pas 2 !)

// ✅ Solution
func bonDefer() {
    x := 1
    defer func() { fmt.Println(x) }()  // closure, x est évalué à l'exécution
    x = 2
}
// Output: 2
```

---

## 5. `panic` et `recover`

### 5.1. Quand utiliser `panic` ?

**Règle d'or :** Utilisez `panic` uniquement pour des erreurs **irrécupérables** qui ne devraient jamais arriver.

```go
// ✅ Cas acceptables pour panic
// - Erreur de programmation (bug)
// - Initialisation impossible (ex: fichier de config manquant au démarrage)
// - Préconditions violées (ex: argument nil alors que requis)

func racineCarree(x float64) float64 {
    if x < 0 {
        panic("racine carrée d'un nombre négatif")  // Erreur de programmation
    }
    return math.Sqrt(x)
}

func init() {
    config, err := lireConfig()
    if err != nil {
        panic(fmt.Sprintf("impossible de charger la config: %v", err))
    }
    // ...
}
```

### 5.2. `panic` – arrêt brutal du programme

```go
func main() {
    fmt.Println("1")
    panic("quelque chose a mal tourné")
    fmt.Println("2")  // Jamais exécuté
}
// Output:
// 1
// panic: quelque chose a mal tourné
// ... stack trace ...
```

### 5.3. `recover` – reprendre le contrôle après un panic

```go
func proteger(fn func()) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Panic récupéré:", r)
        }
    }()
    fn()
}

func main() {
    proteger(func() {
        panic("erreur fatale")
    })
    fmt.Println("Le programme continue !")
}
// Output:
// Panic récupéré: erreur fatale
// Le programme continue !
```

### 5.4. Cas d'usage typique : serveur HTTP

```go
func serveurAvecRecovery(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("PANIC: %v", err)
                http.Error(w, "Erreur interne", http.StatusInternalServerError)
            }
        }()
        handler(w, r)
    }
}
```

---

## 6. Logging avec le package `log`

### 6.1. Utilisation basique

```go
import "log"

func main() {
    log.Println("Message standard")
    log.Printf("Message formaté: %s", "valeur")
    log.Fatal("Erreur fatale")  // log + exit(1)
    log.Panic("Panic!")          // log + panic
}
```

### 6.2. Configuration du logger

```go
func init() {
    // Format: date + heure + fichier:ligne
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

    // Prefix personnalisé
    log.SetPrefix("[MON_APP] ")

    // Sortie vers fichier
    fichier, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(fichier)
}
```

### 6.3. Niveaux de log (avec un logger personnalisé)

```go
// Créer des loggers pour différents niveaux
var (
    InfoLogger  *log.Logger
    WarnLogger  *log.Logger
    ErrorLogger *log.Logger
)

func init() {
    InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
    WarnLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
    InfoLogger.Println("Démarrage de l'application")
    WarnLogger.Println("Configuration non optimale")
    ErrorLogger.Println("Impossible de se connecter à la base de données")
}
```

---

## 7. Bonnes pratiques pour des programmes robustes

### 7.1. Toujours traiter les erreurs

```go
// ❌ Mauvais : ignorer l'erreur
result, _ := diviser(10, 0)

// ✅ Bon : traiter l'erreur
result, err := diviser(10, 0)
if err != nil {
    log.Printf("Erreur: %v", err)
    return
}
```

### 7.2. Ne pas utiliser `panic` pour les erreurs normales

```go
// ❌ Mauvais - une entrée utilisateur erronée ne mérite pas un panic
func lireAge() int {
    var age int
    fmt.Scan(&age)
    if age < 0 {
        panic("âge négatif")  // NON !
    }
    return age
}

// ✅ Bon - retourner une erreur
func lireAge() (int, error) {
    var age int
    fmt.Scan(&age)
    if age < 0 {
        return 0, errors.New("âge négatif invalide")
    }
    return age, nil
}
```

### 7.3. Enrichir les erreurs avec du contexte

```go
// ❌ Pauvre
return err

// ✅ Riche
return fmt.Errorf("impossible de lire le fichier %s: %w", nom, err)
```

### 7.4. Définir les erreurs de package en tête de fichier

```go
// user/user.go
package user

import "errors"

// Erreurs exportées du package
var (
    ErrNotFound      = errors.New("user: utilisateur non trouvé")
    ErrInvalidEmail  = errors.New("user: email invalide")
    ErrDuplicateUser = errors.New("user: utilisateur déjà existant")
)
```

---

## TP final – Module 7 – Calculateur robuste avec logging

### Énoncé

Créez un programme qui lit deux nombres et un opérateur depuis l'utilisateur, effectue l'opération, et gère **toutes** les erreurs possibles.

**Fonctionnalités :**

1. **Lecture des entrées** avec validation
2. **Opérations supportées** : `+`, `-`, `*`, `/`
3. **Gestion des erreurs** :
   - Division par zéro
   - Opérateur invalide
   - Entrée non numérique
4. **Logging** des erreurs dans un fichier
5. **Panic/recover** pour protéger le programme

### Structure attendue

```
module7/
├── go.mod
├── main.go
├── calculator/
│   └── calculator.go
└── app.log          (généré à l'exécution)
```

### Exemple de sortie attendue

```
=== Calculateur robuste ===

Entrez le premier nombre: abc
ERREUR: nombre invalide

Entrez le premier nombre: 10
Entrez le deuxième nombre: 0
Entrez l'opérateur (+, -, *, /): /
ERREUR: division par zéro

Entrez le premier nombre: 10
Entrez le deuxième nombre: 3
Entrez l'opérateur (+, -, *, /): *
Résultat: 10 * 3 = 30

Voulez-vous continuer ? (o/n): n
Au revoir !
```

### Étapes à suivre

**Étape 1 – Initialiser le module**

```bash
mkdir -p module7/calculator
cd module7
go mod init module7
```

**Étape 2 – Créer le package calculator**

```go
// calculator/calculator.go
package calculator

import (
    "errors"
    "fmt"
)

// Erreurs exportées
var (
    ErrDivisionParZero = errors.New("calculator: division par zéro")
    ErrOperateurInvalide = errors.New("calculator: opérateur invalide")
)

// Calcul effectue une opération mathématique basique
func Calcul(a, b float64, operateur string) (float64, error) {
    switch operateur {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        if b == 0 {
            return 0, ErrDivisionParZero
        }
        return a / b, nil
    default:
        return 0, fmt.Errorf("%w: %s", ErrOperateurInvalide, operateur)
    }
}
```

**Étape 3 – Créer le main avec logging**

```go
// main.go
package main

import (
    "bufio"
    "fmt"
    "log"
    "module7/calculator"
    "os"
    "strconv"
    "strings"
)

// Logger personnalisé
var (
    infoLog  *log.Logger
    errorLog *log.Logger
)

func init() {
    // Configuration des logs vers fichier
    fichier, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal("Impossible d'ouvrir le fichier de log:", err)
    }

    infoLog = log.New(fichier, "INFO: ", log.Ldate|log.Ltime)
    errorLog = log.New(fichier, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
    // Recovery global pour éviter le crash
    defer func() {
        if r := recover(); r != nil {
            errorLog.Printf("PANIC récupéré: %v", r)
            fmt.Println("\n❌ Une erreur inattendue s'est produite")
        }
    }()

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("=== Calculateur robuste ===\n")

    for {
        // Lecture du premier nombre
        a, err := lireNombre(scanner, "Entrez le premier nombre: ")
        if err != nil {
            errorLog.Printf("Erreur lecture nombre: %v", err)
            fmt.Println("❌", err)
            continue
        }

        // Lecture du deuxième nombre
        b, err := lireNombre(scanner, "Entrez le deuxième nombre: ")
        if err != nil {
            errorLog.Printf("Erreur lecture nombre: %v", err)
            fmt.Println("❌", err)
            continue
        }

        // Lecture de l'opérateur
        operateur := lireOperateur(scanner)

        // Calcul avec gestion d'erreur
        resultat, err := calculator.Calcul(a, b, operateur)
        if err != nil {
            errorLog.Printf("Erreur calcul %f %s %f: %v", a, operateur, b, err)
            fmt.Printf("❌ Erreur: %v\n", err)
            continue
        }

        // Succès !
        infoLog.Printf("Calcul réussi: %f %s %f = %f", a, operateur, b, resultat)
        fmt.Printf("\n✅ Résultat: %.2f %s %.2f = %.2f\n", a, operateur, b, resultat)

        // Continuer ?
        fmt.Print("\nVoulez-vous continuer ? (o/n): ")
        scanner.Scan()
        reponse := strings.ToLower(strings.TrimSpace(scanner.Text()))
        if reponse != "o" && reponse != "oui" {
            fmt.Println("\nAu revoir !")
            break
        }
        fmt.Println()
    }
}

func lireNombre(scanner *bufio.Scanner, prompt string) (float64, error) {
    fmt.Print(prompt)
    scanner.Scan()
    input := strings.TrimSpace(scanner.Text())

    // Protection contre panic (recover capturera si besoin)
    nombre, err := strconv.ParseFloat(input, 64)
    if err != nil {
        return 0, fmt.Errorf("'%s' n'est pas un nombre valide", input)
    }
    return nombre, nil
}

func lireOperateur(scanner *bufio.Scanner) string {
    for {
        fmt.Print("Entrez l'opérateur (+, -, *, /): ")
        scanner.Scan()
        op := strings.TrimSpace(scanner.Text())

        if op == "+" || op == "-" || op == "*" || op == "/" {
            return op
        }

        errorLog.Printf("Opérateur invalide: %s", op)
        fmt.Println("❌ Opérateur invalide. Utilisez +, -, * ou /")
    }
}
```

**Étape 4 – Exécuter**

```bash
go run main.go
```

**Étape 5 – Vérifier les logs**

```bash
cat app.log
# INFO: 2026/01/15 14:30:22 Calcul réussi: 10.000000 * 3.000000 = 30.000000
# ERROR: 2026/01/15 14:30:15 main.go:xx: Erreur lecture nombre: 'abc' n'est pas un nombre valide
```

---

## Récapitulatif des acquis

À la fin de ce module, vous savez :

- ✅ Gérer les erreurs avec le pattern `if err != nil`
- ✅ Créer des erreurs avec `errors.New()` et `fmt.Errorf()`
- ✅ Utiliser `errors.Is()` et `errors.As()` pour inspecter les erreurs
- ✅ Programmer l'exécution différée avec `defer`
- ✅ Comprendre quand utiliser `panic` vs erreur normale
- ✅ Récupérer un `panic` avec `recover`
- ✅ Mettre en place du logging avec le package `log`
- ✅ Écrire des programmes robustes qui ne plantent jamais

---

## Prochain module (aperçu)

Dans le Module 8, nous découvrirons :
- Les interfaces (le cœur de la POO en Go)
- Le polymorphisme implicite
- L'assertion de type et les `type switches`
- Le pattern "accept interfaces, return structs"

---

**Félicitations !** Vous savez maintenant écrire des programmes Go robustes qui gèrent correctement les erreurs et ne plantent jamais de manière inattendue. 🛡️
