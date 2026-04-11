# Module 3 – Contrôle de flux

Bienvenue dans ce troisième module ! À la fin de ce module, vous saurez prendre des décisions dans votre programme (conditions), répéter des actions (boucles), et gérer des erreurs simples. Vous écrirez un jeu de devinette complet.

---

## 0. Les packages en Go

Avant de commencer, comprenons comment Go organise son code. Un **package** est un ensemble de fonctionnalités regroupées. Pour les utiliser, on les importe.

### Packages couramment utilisés

```go
import (
    "fmt"       // Entrées/sorties (Println, Scanln, Printf)
    "errors"    // Création d'erreurs (errors.New)
    "strconv"   // Conversion texte ↔ nombres (Atoi, Itoa)
    "math/rand" // Nombres aléatoires (rand.Intn)
    "time"      // Temps (pour initialiser le hasard)
)
```

| Package | Utilité | Exemple |
|---------|---------|---------|
| `fmt` | Entrées/sorties basiques | `fmt.Println("texte")`, `fmt.Scanln(&var)` |
| `errors` | Créer des erreurs | `errors.New("message")` |
| `strconv` | Convertir string ↔ nombres | `strconv.Atoi("42")`, `strconv.Itoa(42)` |
| `math/rand` | Nombres aléatoires | `rand.Intn(100)` (0 à 99) |
| `time` | Temps (pour le hasard) | `time.Now().UnixNano()` |

### Pourquoi `strconv` existe ?

En Go, on ne peut pas faire `int("42")` comme dans certains langages :

```go
// ❌ Ne fonctionne pas
nombre := int("42")  // Erreur de compilation !

// ✅ Il faut utiliser strconv
import "strconv"
nombre, err := strconv.Atoi("42")  // nombre = 42, err = nil
```

⚠️ **Important :** `strconv.Atoi()` retourne **toujours deux valeurs** : le nombre converti ET une éventuelle erreur.

```go
valeur, err := strconv.Atoi("abc")
if err != nil {
    fmt.Println("Erreur :", err)  // Erreur : invalid syntax
}
```

### Où trouver la documentation ?

- En ligne : [pkg.go.dev](https://pkg.go.dev)
- Dans le terminal : `go doc fmt.Println`
- Votre éditeur de code (affichage au survol)

💡 **Conseil :** Vous n'avez pas besoin de mémoriser tous les packages. Apprenez à chercher la documentation quand vous en avez besoin.

---

## 1. Instructions conditionnelles

Les conditions permettent d'exécuter du code seulement si une certaine situation est vraie.

### 1.1. `if`, `else if`, `else`

La structure `if` s'utilise **sans parenthèses** autour de la condition.

```go
age := 18

if age >= 18 {
    fmt.Println("Vous êtes majeur")
} else {
    fmt.Println("Vous êtes mineur")
}
```

Avec plusieurs cas :

```go
note := 85

if note >= 90 {
    fmt.Println("Excellent")
} else if note >= 70 {
    fmt.Println("Bien")
} else if note >= 50 {
    fmt.Println("Passable")
} else {
    fmt.Println("Insuffisant")
}
```

💡 **Particularité de Go :** On peut exécuter une instruction avant la condition :

```go
if age := 20; age >= 18 {
    fmt.Println("Majeur")
}
// age n'est plus accessible ici
```

### 1.2. `switch` – pour les cas multiples

Quand vous avez beaucoup de cas à vérifier, `switch` est plus lisible.

```go
jour := "mardi"

switch jour {
case "lundi":
    fmt.Println("Début de semaine")
case "mardi", "mercredi", "jeudi":
    fmt.Println("Milieu de semaine")
case "vendredi":
    fmt.Println("Enfin vendredi !")
default:
    fmt.Println("Week-end")
}
```

⚠️ **Particularité importante :** En Go, les `case` n'ont pas besoin de `break`. Le `switch` s'arrête automatiquement après le premier cas correspondant.

**Switch sans expression** (équivalent à `if`/`else` amélioré) :

```go
note := 85

switch {
case note >= 90:
    fmt.Println("A")
case note >= 80:
    fmt.Println("B")
case note >= 70:
    fmt.Println("C")
default:
    fmt.Println("D ou moins")
}
```

---

## 2. Boucles

En Go, il n'y a qu'un seul mot-clé pour les boucles : **`for`**. Pas de `while` ni de `do...while`.

### 2.1. `for` classique

```go
for i := 0; i < 5; i++ {
    fmt.Println("Tour numéro", i)
}
// Affiche 0, 1, 2, 3, 4
```

### 2.2. `for` comme un `while`

```go
compteur := 0
for compteur < 3 {
    fmt.Println("Encore...")
    compteur++
}
```

### 2.3. Boucle infinie

```go
for {
    fmt.Println("Tourne en rond")
    break   // sans break, c'est infini !
}
```

### 2.4. `break` et `continue`

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue   // Saute i=3, passe au suivant
    }
    if i == 7 {
        break      // S'arrête complètement à i=7
    }
    fmt.Println(i) // 0,1,2,4,5,6
}
```

### 2.5. `for range` – parcourir des collections

`range` permet d'itérer facilement sur des slices, tableaux, maps ou chaînes.

```go
nombres := []int{10, 20, 30, 40}

for index, valeur := range nombres {
    fmt.Println("Position", index, "=", valeur)
}

// Si l'index ne vous intéresse pas, utilisez _
for _, valeur := range nombres {
    fmt.Println("Valeur :", valeur)
}
```

💡 Nous verrons `range` plus en détail dans le module sur les collections.

---

## 3. Gestion d'erreurs

En Go, la gestion d'erreur est **explicite**. Pas d'exceptions comme en Python ou Java. Une fonction "qui peut échouer" retourne souvent un résultat et une erreur.

### 3.1. Créer une erreur avec `errors.New()`

```go
import "errors"

func diviser(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division par zéro impossible")
    }
    return a / b, nil   // nil = pas d'erreur
}
```

### 3.2. Vérifier avec `if err != nil`

C'est le **motif le plus fréquent** en Go :

```go
resultat, err := diviser(10, 0)
if err != nil {
    fmt.Println("Erreur :", err)
} else {
    fmt.Println("Résultat :", resultat)
}
```

✅ **À retenir :**
- Si `err != nil` → gérez l'erreur
- Si `err == nil` → tout va bien, utilisez le résultat

```go
// Exemple avec conversion string → int
nombreStr := "42"
valeur, err := strconv.Atoi(nombreStr)
if err != nil {
    fmt.Println("Impossible de convertir :", err)
} else {
    fmt.Println("Valeur convertie :", valeur)
}
```

---

## TP final – Jeu de devinette

### Énoncé

Créez un jeu où l'ordinateur choisit un nombre secret (fixé dans le code). Le joueur doit le deviner en 5 essais maximum.

**Règles :**
- Nombre secret : 42 (vous pourrez changer plus tard)
- Le programme indique "trop petit", "trop grand" ou "gagné"
- Limite de 5 tentatives
- Si échec après 5 essais, afficher le nombre secret

**Exemple de sortie :**

```
=== Jeu de Devinette ===
Devinez le nombre secret (entre 1 et 100) : 50
Trop grand ! Il reste 4 essais.
Devinez le nombre secret (entre 1 et 100) : 25
Trop petit ! Il reste 3 essais.
Devinez le nombre secret (entre 1 et 100) : 42
Bravo ! Vous avez trouvé en 3 essais.
```

### Étapes à suivre

**Étape 1 – Créer le fichier**

Créez un dossier `module3` et un fichier `devinette.go`.

**Étape 2 – Structure de base**

```go
package main

import "fmt"

func main() {
    // Votre code ici
}
```

**Étape 3 – Déclarez les constantes**

```go
const (
    NombreSecret = 42
    MaxEssais    = 5
)
```

**Étape 4 – Créez une variable pour suivre la victoire**

```go
gagne := false
```

**Étape 5 – Écrivez une boucle `for` pour les essais**

Deux approches possibles :
- `for i := 1; i <= MaxEssais; i++`
- `for essaisRestants := MaxEssais; essaisRestants > 0; essaisRestants--`

**Étape 6 – Lisez la tentative de l'utilisateur**

```go
var tentative int
fmt.Print("Devinez le nombre : ")
fmt.Scanln(&tentative)
```

**Étape 7 – Comparez avec `if` et affichez l'indice**

- Si égal → `gagne = true` et `break`
- Si trop petit → "Trop petit !"
- Si trop grand → "Trop grand !"
- Affichez le nombre d'essais restants

**Étape 8 – Affichez le résultat final**

- Si `gagne` est `true` → message de victoire
- Sinon → message de défaite avec révélation du nombre secret

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

import "fmt"

func main() {
    const (
        NombreSecret = 42
        MaxEssais    = 5
    )

    fmt.Println("=== Jeu de Devinette ===")
    gagne := false

    for essaisRestants := MaxEssais; essaisRestants > 0; essaisRestants-- {
        var tentative int
        fmt.Print("Devinez le nombre : ")
        fmt.Scanln(&tentative)

        if tentative == NombreSecret {
            gagne = true
            break
        } else if tentative < NombreSecret {
            fmt.Printf("Trop petit ! Il reste %d essais.\n", essaisRestants-1)
        } else {
            fmt.Printf("Trop grand ! Il reste %d essais.\n", essaisRestants-1)
        }
    }

    if gagne {
        fmt.Println("Bravo ! Vous avez trouvé le nombre secret !")
    } else {
        fmt.Printf("Dommage ! Le nombre secret était %d.\n", NombreSecret)
    }
}
```

**Variante avec compteur d'essais :**

```go
gagne := false
for i := 1; i <= MaxEssais; i++ {
    var tentative int
    fmt.Print("Devinez le nombre : ")
    fmt.Scanln(&tentative)

    if tentative == NombreSecret {
        gagne = true
        fmt.Printf("Bravo ! Vous avez trouvé en %d essais !\n", i)
        break
    } else if tentative < NombreSecret {
        fmt.Printf("Trop petit ! (Essai %d/%d)\n", i, MaxEssais)
    } else {
        fmt.Printf("Trop grand ! (Essai %d/%d)\n", i, MaxEssais)
    }
}

if !gagne {
    fmt.Printf("Dommage ! Le nombre secret était %d.\n", NombreSecret)
}
```

### Tests à effectuer

```bash
go run devinette.go
```

Testez ces cas :
- ✅ Deviner du premier coup
- ✅ Deviner après plusieurs essais
- ✅ Épuiser les 5 essais sans trouver
- ✅ Changer `NombreSecret` en 10
- ✅ Changer `MaxEssais` en 3

### Pour aller plus loin (optionnel)

**1. Nombre aléatoire**

```go
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    nombreSecret := rand.Intn(100) + 1  // Entre 1 et 100
    // ... reste du code identique
}
```

**2. Validation des entrées**

```go
var input string
fmt.Print("Devinez le nombre : ")
fmt.Scanln(&input)
tentative, err := strconv.Atoi(input)
if err != nil {
    fmt.Println("Veuillez entrer un nombre valide !")
    continue
}
```

**3. Mode difficile** (moins d'essais)

**4. Rejouer à la fin** (boucle englobante)

**5. Ajouter des indices** après 3 essais

---

## Récapitulatif des acquis

À la fin de ce module, vous savez :

- ✅ Comprendre les packages et les importer (`fmt`, `errors`, `strconv`...)
- ✅ Prendre des décisions avec `if`/`else if`/`else`
- ✅ Utiliser `switch` pour des cas multiples
- ✅ Créer des boucles avec `for` (classique, while, infini)
- ✅ Utiliser `break` et `continue`
- ✅ Gérer des erreurs simples avec `if err != nil`

---

➡️ **Module 4 :** Nous découvrirons les fonctions, les retours multiples, et les pointeurs.

**Félicitations !** Vous maîtrisez maintenant le contrôle de flux en Go. 🚀
