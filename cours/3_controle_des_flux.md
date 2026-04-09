Module 3 – Contrôle de flux

Bienvenue dans ce troisième module ! À la fin de ce module, vous saurez prendre des décisions dans votre programme (conditions), répéter des actions (boucles), et gérer des erreurs simples. Vous écrirez un jeu de devinette complet.

0. Parenthèse – Les packages en Go (imports)

Avant de commencer, comprenons comment Go organise son code. Un package est un ensemble de fonctionnalités regroupées. Pour les utiliser, on les importe.

Les packages que nous utilisons déjà et allons utiliser

```go
package main

import "fmt"      // Pour afficher (Println, Printf) et lire (Scanln)
import "errors"   // Pour créer des erreurs (errors.New)
import "strconv"  // Pour convertir texte ↔ nombres (Atoi, Itoa)
import "math/rand"// Pour générer des nombres aléatoires (optionnel)
import "time"     // Pour le temps (nécessaire avec rand)
```

Version raccourcie (recommandée) :

```go
import (
    "fmt"
    "errors"
    "strconv"
    "math/rand"
    "time"
)
```

Que fait chaque package ?

Package Utilité Exemple
fmt Entrées/sorties basiques fmt.Println("texte") fmt.Scanln(&variable) fmt.Printf("%.2f", prix)
errors Créer des erreurs errors.New("message d'erreur")
strconv Convertir string ↔ nombres strconv.Atoi("42") (string → int) strconv.Itoa(42) (int → string) strconv.ParseFloat("3.14", 64)
math/rand Nombres aléatoires rand.Intn(100) (0 à 99)
time Temps (pour initialiser le hasard) time.Now().UnixNano()

Pourquoi strconv existe ?

En Go, on ne peut pas faire int("42") comme dans certains langages. Il faut passer par strconv.Atoi() (ASCII to Integer).

```go
// ❌ Ceci ne fonctionne PAS
nombre := int("42")  // Erreur de compilation !

// ✅ Il faut faire :
import "strconv"
nombre, err := strconv.Atoi("42")  // nombre = 42, err = nil
```

⚠️ Atoi retourne toujours deux valeurs : le nombre converti ET une éventuelle erreur.

```go
valeur, err := strconv.Atoi("abc")
if err != nil {
    fmt.Println("Erreur :", err)  // Erreur : strconv.Atoi: parsing "abc": invalid syntax
}
```

Où trouver la documentation ?

· En ligne : pkg.go.dev (ex: pkg.go.dev/fmt)
· Dans le terminal : go doc fmt.Println
· Votre éditeur de code affiche souvent les informations au survol

💡 Conseil : Vous n'avez pas besoin de mémoriser tous les packages. Apprenez à chercher la documentation quand vous en avez besoin.

---

1. Instructions conditionnelles

Les conditions permettent d'exécuter du code seulement si une certaine situation est vraie.

1.1. if, else if, else

La structure if est la plus basique. Elle s'utilise sans parenthèses autour de la condition.

```go
age := 18

if age >= 18 {
    fmt.Println("Vous êtes majeur")
}
```

Avec une alternative :

```go
age := 16

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

💡 Particularité de Go : On peut exécuter une petite instruction avant la condition :

```go
// age n'existe qu'à l'intérieur du if/else
if age := 20; age >= 18 {
    fmt.Println("Majeur")
}
// age n'est plus accessible ici
```

1.2. switch – alternative élégante au if/else chaîné

Quand vous avez beaucoup de cas à vérifier, switch est plus lisible.

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

⚠️ Particularité importante : En Go, les case n'ont pas besoin de break . Le switch s'arrête automatiquement après le premier cas correspondant.

```go
// Switch sans expression = if/else amélioré
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

2. Boucles

En Go, il n'y a qu'un seul mot-clé pour les boucles : for . Pas de while ni de do...while.

2.1. for classique – comme en C/Java

```go
// for initialisation; condition; post-iteration
for i := 0; i < 5; i++ {
    fmt.Println("Tour numéro", i)
}
// Affiche 0, 1, 2, 3, 4
```

Le for comme un while :

```go
// Tant que compteur < 3
compteur := 0
for compteur < 3 {
    fmt.Println("Encore...")
    compteur++
}
```

Boucle infinie (utile avec break) :

```go
for {
    fmt.Println("Tourne en rond")
    break   // sans ça, c'est infini !
}
```

break et continue :

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue   // saute i=3, passe au suivant
    }
    if i == 7 {
        break      // s'arrête complètement à i=7
    }
    fmt.Println(i) // 0,1,2,4,5,6
}
```

2.2. for range – pour parcourir des collections

range permet d'itérer facilement sur des slices, tableaux, maps ou chaînes.

```go
// Avec une slice
nombres := []int{10, 20, 30, 40}

for index, valeur := range nombres {
    fmt.Println("Position", index, "=", valeur)
}
```

```go
// Si l'index ne vous intéresse pas, utilisez _
for _, valeur := range nombres {
    fmt.Println("Valeur :", valeur)
}
```

💡 Nous verrons range plus en détail dans le module sur les collections.

3. Gestion d'erreurs simple

En Go, la gestion d'erreur est explicite. Pas d'exceptions comme en Python/Java. Une fonction "qui peut échouer" retourne souvent un résultat et une erreur.

3.1. Créer une erreur avec errors.New()

```go
import "errors"

func diviser(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division par zéro impossible")
    }
    return a / b, nil   // nil = pas d'erreur
}
```

3.2. Vérifier if err != nil

C'est le motif le plus fréquent en Go :

```go
resultat, err := diviser(10, 0)
if err != nil {
    fmt.Println("Erreur :", err)
} else {
    fmt.Println("Résultat :", resultat)
}
```

✅ À retenir : Si err != nil, gérez l'erreur. Si err == nil, tout va bien.

```go
// Exemple avec conversion string → int (package strconv)
import "strconv"

nombreStr := "42"
valeur, err := strconv.Atoi(nombreStr)
if err != nil {
    fmt.Println("Impossible de convertir :", err)
} else {
    fmt.Println("Valeur convertie :", valeur)
}
```

---

TP final – Module 3 – Jeu de devinette

Énoncé

Créez un jeu où l'ordinateur choisit un nombre secret (fixé dans le code). Le joueur doit le deviner en 5 essais maximum.

Règles :

· Nombre secret : 42 (vous pourrez changer plus tard)
· Le programme indique "trop petit", "trop grand" ou "gagné"
· Limite de 5 tentatives
· Si échec après 5 essais, afficher le nombre secret

Exemple de sortie attendue :

```
=== Jeu de Devinette ===
Devinez le nombre secret (entre 1 et 100) : 50
Trop grand ! Il reste 4 essais.
Devinez le nombre secret (entre 1 et 100) : 25
Trop petit ! Il reste 3 essais.
Devinez le nombre secret (entre 1 et 100) : 42
Bravo ! Vous avez trouvé en 3 essais.
```

Étapes à suivre

Étape 1 – Créer le fichier

Créez un dossier module3 et un fichier devinette.go.

Étape 2 – Structure de base avec les imports

```go
package main

import "fmt"   // Pour Println, Printf, Scanln

func main() {
    // Votre code ici
}
```

Étape 3 – Déclarez les constantes

Déclarez une constante pour le nombre secret et une pour le nombre maximum d'essais.

Étape 4 – Créez une variable pour suivre si le joueur a gagné

Une variable bool sera utile pour savoir si le joueur a trouvé avant la fin des essais.

Étape 5 – Écrivez une boucle for pour les 5 essais

Deux approches possibles :

· Un for i := 1; i <= 5; i++ où i représente le numéro de l'essai
· Un for essaisRestants := 5; essaisRestants > 0; essaisRestants--

Étape 6 – À chaque tour : lisez la tentative de l'utilisateur avec fmt.Scanln()

```go
var tentative int
fmt.Print("Devinez le nombre : ")
fmt.Scanln(&tentative)
```

Étape 7 – Comparez avec if et affichez l'indice

· Si égal → gagne = true et break (on sort de la boucle)
· Si trop petit → affichez "Trop petit !"
· Si trop grand → affichez "Trop grand !"
· N'oubliez pas d'afficher le nombre d'essais restants

Étape 8 – Après la boucle, affichez le résultat final

· Si gagne est true → message de victoire
· Sinon → message de défaite avec révélation du nombre secret

À vous de jouer !

Essayez de coder la solution par vous-même avant de regarder la proposition ci-dessous.

L'important est que votre programme produise le comportement attendu. Il n'y a pas une seule façon correcte de le faire. Tant que votre jeu fonctionne comme décrit, c'est une bonne solution.

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

Proposition de correction

Voici une solution possible. La vôtre peut être différente et tout à fait valable !

```go
package main

import (
    "fmt"
)

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

Variante possible avec une autre structure de boucle :

```go
// Version avec i pour compter les essais
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

Comment tester votre programme

Exécution directe :

```bash
go run devinette.go
```

Testez les différents cas :

· ✅ Devinez correctement du premier coup
· ✅ Devinez après plusieurs essais
· ✅ Épuisez les 5 essais sans trouver
· ✅ Changez NombreSecret en 10 et rejouez
· ✅ Changez MaxEssais en 3 pour plus de difficulté

Pour aller plus loin (optionnel)

1. Générer un nombre aléatoire (nécessite les packages math/rand et time)

```go
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())  // Initialise le hasard
    NombreSecret := rand.Intn(100) + 1 // entre 1 et 100
    // ... reste du code identique
}
```

2. Valider que l'entrée est bien un nombre (avec strconv)

```go
import (
    "fmt"
    "strconv"
)

// Pour lire une ligne complète et la convertir
var input string
fmt.Print("Devinez le nombre : ")
fmt.Scanln(&input)
tentative, err := strconv.Atoi(input)
if err != nil {
    fmt.Println("Veuillez entrer un nombre valide !")
    continue  // On recommence l'essai
}
```

3. Ajouter un mode "difficile" avec moins d'essais

4. Proposer de rejouer à la fin avec une boucle for qui englobe tout le jeu

5. Ajouter un indice après 3 essais ("Le nombre est pair" ou "Le nombre est multiple de 3")

---

Félicitations !

Vous maîtrisez maintenant les conditions (if, switch), les boucles (for sous toutes ses formes), et la gestion d'erreur basique en Go.

Récap des acquis :

· ✅ Comprendre les packages et les importer (fmt, errors, strconv...)
· ✅ Prendre des décisions avec if/else if/else
· ✅ Utiliser switch pour des cas multiples
· ✅ Créer des boucles avec for
· ✅ Gérer des erreurs simples avec if err != nil

➡️ Dans le Module 4, nous découvrirons les fonctions : créer ses propres blocs réutilisables, retourner des valeurs multiples, et comprendre les pointeurs.