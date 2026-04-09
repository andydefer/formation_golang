
Module 2 – Types et variables

Bienvenue dans ce deuxième module !
À la fin de ce module, vous saurez déclarer des variables, utiliser les types de base de Go, créer des constantes, convertir des types, et vous écrirez un mini-calculateur fonctionnel.

---

1. Types de base

En Go, chaque donnée a un type. Le type indique à l’ordinateur :

· quelle taille de mémoire réserver
· quelles opérations sont autorisées (+, -, /, *, etc.)

1.1. int – les nombres entiers

int représente un nombre entier (positif ou négatif).

```go
var age int = 28
var temperature int = -5
```

💡 Go a aussi des variantes plus précises : int8, int16, int32, int64 (pour économiser de la mémoire). Mais pour débuter, int est parfait.

1.2. float64 – les nombres décimaux

float64 sert à stocker des nombres avec une virgule.

```go
var prix float64 = 19.99
var pi float64 = 3.14159
```

⚠️ Piège fréquent : En Go, 5 / 2 donne 2 (division entière). Pour obtenir 2.5, il faut écrire 5.0 / 2 ou utiliser des float64.

1.3. string – le texte

string représente une chaîne de caractères (texte), toujours entre guillemets doubles.

```go
var nom string = "Aline"
var message string = "Bonjour tout le monde !"
```

1.4. bool – les booléens (vrai / faux)

bool ne prend que deux valeurs : true ou false.

```go
var estMajeur bool = true
var estPret bool = false
```

✅ Les booléens sont très utiles dans les conditions (if, for).

---

2. Déclaration des variables

Il existe plusieurs façons de déclarer une variable en Go. Chacune a son usage.

2.1. var nom type – déclaration complète

C’est la forme la plus explicite.

```go
var prenom string
prenom = "Aline"

var age int = 28
```

On peut aussi déclarer sans initialiser : Go mettra une valeur zéro :

· 0 pour les nombres
· "" (chaîne vide) pour les string
· false pour les bool

```go
var x int        // x = 0
var texte string // texte = ""
var actif bool   // actif = false
```

2.2. := – inférence de type (déclaration courte)

C’est la façon la plus pratique et la plus utilisée en Go.
Go devine le type automatiquement à partir de la valeur.

```go
prenom := "Aline"   // Go comprend que c'est un string
age := 28           // Go comprend que c'est un int
prix := 19.99       // Go comprend que c'est un float64
estMajeur := true   // Go comprend que c'est un bool
```

⚠️ Règle importante : := ne peut être utilisé qu’à l’intérieur d’une fonction (pas au niveau global).
À l’extérieur d’une fonction, utilisez var.

2.3. Variables multiples

On peut déclarer plusieurs variables sur une seule ligne.

```go
// Avec var
var a, b, c int = 1, 2, 3

// Avec :=
x, y := 10, "bonjour"

// Mélange de types
nom, age, actif := "Aline", 28, true
```

💡 Conseil : La déclaration multiple est très pratique pour les retours de fonctions (ex: resultat, err := maFonction()).

---

3. Constantes et iota

3.1. const – une valeur qui ne change jamais

Une constante se déclare avec const. Elle ne peut pas être modifiée après sa création.

```go
const Pi = 3.14159
const NomProgramme = "MonCalculateur"

// Déclaration multiple
const (
    HeuresParJour = 24
    JoursParSemaine = 7
)
```

✅ Quand utiliser const ?
Pour des valeurs qui ne varient jamais : nombre de secondes dans une minute, taux de TVA, nom d’un programme, etc.

3.2. iota – pour créer des énumérations simples

iota est un générateur automatique de constantes consécutives.
Il est réinitialisé à 0 à chaque nouveau bloc const.

```go
const (
    Lundi = iota     // 0
    Mardi            // 1
    Mercredi         // 2
    Jeudi            // 3
    Vendredi         // 4
)

const (
    Un = iota + 1    // 1
    Deux             // 2
    Trois            // 3
)
```

💡 iota évite d’écrire manuellement 0, 1, 2, 3... pour des énumérations.

---

4. Conversion de types

En Go, on ne peut pas mélanger deux types différents sans conversion explicite.

❌ Ceci ne fonctionne PAS :

```go
var a int = 10
var b float64 = 3.5
var c = a + b  // Erreur ! (int + float64 interdit)
```

✅ Il faut convertir :

```go
var a int = 10
var b float64 = 3.5

// Convertir int en float64
var somme1 float64 = float64(a) + b   // 13.5

// Convertir float64 en int (perd la virgule)
var somme2 int = a + int(b)           // 13 (car 3.5 → 3)
```

Conversions courantes

```go
var x int = 42
var y float64 = float64(x)   // int → float64
var z int = int(y)           // float64 → int (troncature)

var texte string = "123"
// Pour convertir string → int, il faut un package spécial (strconv)
// Nous verrons cela plus tard.
```

⚠️ Piège : Convertir un float64 en int perd les décimales (troncature, pas d’arrondi).
Exemple : int(3.999) donne 3, pas 4.

---

TP final – Module 2

Énoncé

Créez un mini-calculateur qui :

1. Déclare deux variables a et b avec les valeurs 10 et 3
2. Affiche leur somme, différence, produit et quotient (division en float64)
3. Utilise une constante pour stocker le nom du programme

Exemple de sortie attendue :

```
=== MonMiniCalculateur ===
a = 10, b = 3
Somme : 10 + 3 = 13
Différence : 10 - 3 = 7
Produit : 10 * 3 = 30
Quotient : 10 / 3 = 3.3333333333333335
```

---

Étapes à suivre

Étape 1 – Créer le fichier

Créez un dossier module2 et un fichier calculateur.go.

Étape 2 – Structure de base

```go
package main

import "fmt"

func main() {
    // Votre code ici
}
```

Étape 3 – Déclarer les variables et la constante

```go
const NomProgramme = "MonMiniCalculateur"

a := 10
b := 3
```

Étape 4 – Calculs

```go
somme := a + b
difference := a - b
produit := a * b
quotient := float64(a) / float64(b)  // Conversion importante !
```

Étape 5 – Affichage

```go
fmt.Println("===", NomProgramme, "===")
fmt.Println("a =", a, "b =", b)
fmt.Println("Somme :", a, "+", b, "=", somme)
fmt.Println("Différence :", a, "-", b, "=", difference)
fmt.Println("Produit :", a, "*", b, "=", produit)
fmt.Println("Quotient :", a, "/", b, "=", quotient)
```

---

Code complet corrigé

```go
package main

import "fmt"

func main() {
    const NomProgramme = "MonMiniCalculateur"

    a := 10
    b := 3

    somme := a + b
    difference := a - b
    produit := a * b
    quotient := float64(a) / float64(b)

    fmt.Println("===", NomProgramme, "===")
    fmt.Println("a =", a, "b =", b)
    fmt.Println("Somme :", a, "+", b, "=", somme)
    fmt.Println("Différence :", a, "-", b, "=", difference)
    fmt.Println("Produit :", a, "*", b, "=", produit)
    fmt.Println("Quotient :", a, "/", b, "=", quotient)
}
```

---

Comment tester votre programme

1. Exécution directe :
   ```bash
   go run calculateur.go
   ```
2. Génération d’un exécutable :
   ```bash
   go build calculateur.go
   ./calculateur        # Linux/Mac
   calculateur.exe      # Windows
   ```
3. Modifiez les valeurs :
   · Changez a := 10 en a := 15
   · Changez b := 3 en b := 4
   · Re-exécutez et vérifiez que tous les résultats sont mis à jour
4. Testez les conversions :
   · Supprimez float64(a) et laissez a / b → observez l’erreur de compilation
   · Comprenez pourquoi Go vous oblige à être explicite

---

Pour aller plus loin (optionnel)

· Ajoutez le reste de la division (%), ex: 10 % 3 = 1
· Demandez les valeurs à l’utilisateur avec fmt.Scanln() au lieu de les écrire en dur
· Gérez la division par zéro (si b devient 0)

---

Félicitations ! Vous maîtrisez maintenant les types, les variables, les constantes et les conversions en Go.

➡️ Dans le Module 3, nous apprendrons à contrôler le flux d’un programme avec if, switch et les boucles for.