# Module 5 – Structures et pointeurs

Bienvenue dans ce cinquième module ! À la fin de ce module, vous saurez manipuler des données complexes avec les structures avancées, comprendre les pointeurs et la mémoire, et utiliser les collections dynamiques (slices et maps).

---

## 0. Rappel : La mémoire en Go

Avant de comprendre les pointeurs, il faut comprendre comment Go stocke les données.

```go
// Quand vous écrivez :
age := 25

// En mémoire :
// [age] -> 25 (stocké à un endroit précis)
```

Chaque variable a :
- **Une valeur** (ce qu'elle contient)
- **Une adresse** (où elle se trouve en mémoire)

```go
age := 25
fmt.Println(age)      // 25 (la valeur)
fmt.Println(&age)     // 0xc0000140a0 (l'adresse, exemple)
```

Le symbole `&` donne l'adresse mémoire. Le symbole `*` donne la valeur à une adresse.

---

## 1. Structures avancées (`struct`)

### 1.1. Struct imbriqué

Une structure peut contenir d'autres structures.

```go
// Sous-structure
type Adresse struct {
    Rue      string
    Ville    string
    CodePostal string
}

// Structure principale qui contient Adresse
type Personne struct {
    Nom     string
    Age     int
    Adresse Adresse  // Struct imbriqué
}

func main() {
    p := Personne{
        Nom: "Alice",
        Age: 30,
        Adresse: Adresse{
            Rue:      "123 Rue de Paris",
            Ville:    "Lyon",
            CodePostal: "69001",
        },
    }

    // Accès aux champs imbriqués
    fmt.Println(p.Nom)              // Alice
    fmt.Println(p.Adresse.Ville)    // Lyon
    fmt.Println(p.Adresse.Rue)      // 123 Rue de Paris
}
```

### 1.2. Struct avec tags (pour JSON, validation, etc.)

Les tags sont des métadonnées attachées aux champs.

```go
import "encoding/json"

type Personne struct {
    Nom  string `json:"nom"`      // En JSON, s'appellera "nom"
    Age  int    `json:"age"`
    Email string `json:"email,omitempty"` // Omis si vide
}

func main() {
    p := Personne{Nom: "Bob", Age: 25}

    // Convertir en JSON
    jsonData, _ := json.Marshal(p)
    fmt.Println(string(jsonData)) // {"nom":"Bob","age":25}
}
```

💡 Les tags sont très utilisés avec les bases de données et les API.

### 1.3. Struct anonyme (ponctuel)

Pour une utilisation unique, pas besoin de nommer la structure.

```go
func main() {
    // Structure anonyme
    point := struct {
        X, Y int
    }{10, 20}

    fmt.Println(point.X, point.Y)  // 10 20
}
```

---

## 2. Pointeurs

### 2.1. Qu'est-ce qu'un pointeur ?

Un pointeur est une variable qui contient **l'adresse mémoire** d'une autre variable.

```go
var x int = 42
var p *int = &x   // p pointe vers x

fmt.Println(x)    // 42 (valeur de x)
fmt.Println(&x)   // 0xc0000140a0 (adresse de x)
fmt.Println(p)    // 0xc0000140a0 (adresse pointée)
fmt.Println(*p)   // 42 (valeur à l'adresse pointée)
```

**Schéma visuel :**

```
Variable x : [42] à l'adresse 0xc0000140a0
Pointeur p : [0xc0000140a0] (contient l'adresse de x)

Quand on écrit *p, on va à l'adresse 0xc0000140a0 et on lit 42
```

### 2.2. Créer et utiliser des pointeurs

```go
// Déclaration
var ptr *int        // pointeur vers int (nil par défaut)
fmt.Println(ptr)    // <nil>

// Créer un pointeur vers une nouvelle valeur
nombre := 100
ptr = &nombre

// Modifier la valeur via le pointeur
*ptr = 200
fmt.Println(nombre)  // 200 (la variable originale a changé !)
```

### 2.3. `new()` – Créer un pointeur vers une valeur zéro

```go
p := new(int)   // p est un *int, pointant vers 0
fmt.Println(*p) // 0

*p = 42
fmt.Println(*p) // 42
```

### 2.4. Passage par valeur vs par référence

**Passage par valeur** (par défaut) : Go copie la valeur.

```go
func doubler(v int) {
    v = v * 2
    fmt.Println("Dans fonction:", v)  // 20
}

func main() {
    x := 10
    doubler(x)
    fmt.Println("Dans main:", x)  // 10 (pas changé !)
}
```

**Passage par référence** (avec pointeur) : on modifie l'original.

```go
func doubler(v *int) {
    *v = *v * 2
    fmt.Println("Dans fonction:", *v)  // 20
}

func main() {
    x := 10
    doubler(&x)
    fmt.Println("Dans main:", x)  // 20 (modifié !)
}
```

### 2.5. Pointeurs sur structures

C'est très courant en Go de passer des structures par pointeur.

```go
type Personne struct {
    Nom string
    Age int
}

// Receiver par pointeur (modifie l'original)
func (p *Personne) FeterAnniversaire() {
    p.Age++
}

// Receiver par valeur (ne modifie pas l'original)
func (p Personne) Afficher() {
    fmt.Printf("%s a %d ans\n", p.Nom, p.Age)
}

func main() {
    alice := Personne{Nom: "Alice", Age: 30}

    alice.Afficher()              // Alice a 30 ans
    alice.FeterAnniversaire()
    alice.Afficher()              // Alice a 31 ans
}
```

⚠️ **Règle d'or :** Si une méthode doit modifier la structure → receiver par pointeur (`*Personne`). Sinon, receiver par valeur suffit.

---

## 3. Tableaux et Slices

### 3.1. Tableaux (taille fixe)

Le tableau a une taille **déterminée à la compilation** et ne peut pas changer.

```go
// Déclaration
var nombres [5]int           // tableau de 5 entiers (tous à 0)
couleurs := [3]string{"rouge", "vert", "bleu"}

// Accès
nombres[0] = 10
fmt.Println(nombres[0])  // 10
fmt.Println(len(nombres)) // 5 (taille fixe)

// ❌ Impossible d'ajouter un 6ème élément
// nombres[5] = 20  // Erreur de compilation !
```

**Quand utiliser un tableau ?** Presque jamais. Les slices sont plus flexibles.

### 3.2. Slices (taille dynamique)

Une slice est une vue dynamique sur un tableau. C'est ce qu'on utilise 99% du temps.

```go
// Créer une slice (plusieurs façons)
var vide []int                    // slice vide (nil)
nombres := []int{10, 20, 30}      // avec initialisation
makeslice := make([]int, 3)       // avec make (3 éléments à 0)

// Ajouter des éléments avec append
nombres = append(nombres, 40)     // [10, 20, 30, 40]
nombres = append(nombres, 50, 60) // [10, 20, 30, 40, 50, 60]

// Concaténer deux slices
autres := []int{70, 80}
nombres = append(nombres, autres...) // ... pour "déballer"

// Longueur et capacité
fmt.Println(len(nombres))  // 8 (nombre d'éléments)
fmt.Println(cap(nombres))  // capacité (taille sous-jacente, >= len)
```

### 3.3. len() et cap() – Comprendre la capacité

```go
slice := make([]int, 3, 5)  // len=3, cap=5
fmt.Println(slice)          // [0, 0, 0]
fmt.Println(len(slice))     // 3
fmt.Println(cap(slice))     // 5

slice = append(slice, 10)   // len=4, cap=5
slice = append(slice, 20)   // len=5, cap=5
slice = append(slice, 30)   // len=6, cap=10 (double automatiquement)
```

💡 **Pourquoi la capacité ?** Pour éviter de réallouer trop souvent. Go double la capacité quand c'est nécessaire.

### 3.4. Découpage (slicing)

```go
nombres := []int{0, 1, 2, 3, 4, 5}

a := nombres[1:4]   // [1, 2, 3] (index 1 à 3)
b := nombres[:3]    // [0, 1, 2] (début à 2)
c := nombres[3:]    // [3, 4, 5] (3 à fin)
d := nombres[:]     // copie superficielle (tous les éléments)
```

⚠️ **Attention :** Un slice ne copie pas les données, il crée une vue. Modifier un slice modifie l'original !

```go
original := []int{1, 2, 3}
copieVue := original[0:2]
copieVue[0] = 99
fmt.Println(original)  // [99, 2, 3] (modifié !)

// Pour une vraie copie :
vraieCopie := make([]int, len(original))
copy(vraieCopie, original)
```

### 3.5. Parcourir une slice avec range

```go
fruits := []string{"pomme", "banane", "cerise"}

// Avec index et valeur
for i, fruit := range fruits {
    fmt.Printf("%d: %s\n", i, fruit)
}

// Juste la valeur
for _, fruit := range fruits {
    fmt.Println(fruit)
}
```

---

## 4. Maps (dictionnaires)

Une map associe des **clés** à des **valeurs** (comme un dictionnaire).

### 4.1. Créer une map

```go
// Méthode 1 : make()
ages := make(map[string]int)

// Méthode 2 : littérale
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
    "Charlie": 92,
}

// Méthode 3 : var (nil map, attention !)
var vide map[string]int  // nil map, ne peut pas recevoir d'éléments
```

### 4.2. Ajouter, lire, supprimer

```go
ages := make(map[string]int)

// Ajouter ou modifier
ages["Alice"] = 30
ages["Bob"] = 25
fmt.Println(ages)  // map[Alice:30 Bob:25]

// Lire une valeur
fmt.Println(ages["Alice"])   // 30

// Lire avec vérification d'existence
valeur, existe := ages["Charlie"]
if existe {
    fmt.Println("Charlie a", valeur)
} else {
    fmt.Println("Charlie n'existe pas")
}

// Supprimer
delete(ages, "Bob")
fmt.Println(ages)  // map[Alice:30]
```

### 4.3. Parcourir une map

```go
scores := map[string]int{
    "Alice": 95,
    "Bob": 87,
    "Charlie": 92,
}

for nom, score := range scores {
    fmt.Printf("%s a eu %d\n", nom, score)
}
```

⚠️ **L'ordre n'est pas garanti !** Les maps ne sont pas ordonnées.

### 4.4. Map de slices ou de structures

```go
// Map de slices
etudiantsParCours := map[string][]string{
    "Maths":   {"Alice", "Bob"},
    "Physique": {"Charlie"},
}
etudiantsParCours["Maths"] = append(etudiantsParCours["Maths"], "David")

// Map de structures
type Contact struct {
    Email string
    Phone string
}

annuaire := map[string]Contact{
    "Alice": {Email: "alice@email.com", Phone: "0612345678"},
}
```

---

## TP final – Module 5 – Gestion de contacts

### Énoncé

Créez un programme de gestion de contacts avec :

1. **Structure `Contact`** avec les champs : `Nom` (string) et `Numero` (string)
2. **Une slice de contacts** pour stocker tous les contacts
3. **Une fonction `ajouterContact()`** qui demande à l'utilisateur les infos et ajoute le contact à la slice
4. **Une fonction `afficherTous()`** qui affiche tous les contacts
5. **Dans `main()`** : demander à l'utilisateur d'ajouter 2 contacts, puis les afficher

### Exemple de sortie attendue :

```
=== Gestion de contacts ===

--- Ajout du contact 1/2 ---
Nom : Alice
Numéro : 0612345678
✓ Contact ajouté !

--- Ajout du contact 2/2 ---
Nom : Bob
Numéro : 0787654321
✓ Contact ajouté !

=== Liste des contacts ===
1. Alice - 0612345678
2. Bob - 0787654321
```

### Étapes à suivre

**Étape 1 – Créer le fichier**

Créez un dossier `module5` et un fichier `contacts.go`.

**Étape 2 – Structure de base et imports**

```go
package main

import "fmt"

// Structure Contact
type Contact struct {
    Nom    string
    Numero string
}

func main() {
    // Votre code ici
}
```

**Étape 3 – Écrire la fonction `ajouterContact()`**

```go
func ajouterContact(contacts []Contact) []Contact {
    var nom, numero string

    fmt.Print("Nom : ")
    fmt.Scanln(&nom)
    fmt.Print("Numéro : ")
    fmt.Scanln(&numero)

    nouveau := Contact{Nom: nom, Numero: numero}
    contacts = append(contacts, nouveau)

    fmt.Println("✓ Contact ajouté !")
    return contacts
}
```

**Étape 4 – Écrire la fonction `afficherTous()`**

```go
func afficherTous(contacts []Contact) {
    if len(contacts) == 0 {
        fmt.Println("Aucun contact.")
        return
    }

    fmt.Println("=== Liste des contacts ===")
    for i, contact := range contacts {
        fmt.Printf("%d. %s - %s\n", i+1, contact.Nom, contact.Numero)
    }
}
```

**Étape 5 – Écrire le main**

```go
func main() {
    contacts := []Contact{}  // Slice vide

    fmt.Println("=== Gestion de contacts ===\n")

    // Ajouter 2 contacts
    for i := 1; i <= 2; i++ {
        fmt.Printf("--- Ajout du contact %d/2 ---\n", i)
        contacts = ajouterContact(contacts)
        fmt.Println()
    }

    // Afficher tous les contacts
    afficherTous(contacts)
}
```

### À vous de jouer !

Essayez de coder la solution par vous-même avant de regarder la proposition ci-dessous.

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

// Structure Contact
type Contact struct {
    Nom    string
    Numero string
}

// Ajouter un contact (retourne la nouvelle slice)
func ajouterContact(contacts []Contact) []Contact {
    var nom, numero string

    fmt.Print("Nom : ")
    fmt.Scanln(&nom)
    fmt.Print("Numéro : ")
    fmt.Scanln(&numero)

    nouveau := Contact{
        Nom:    nom,
        Numero: numero,
    }

    contacts = append(contacts, nouveau)
    fmt.Println("✓ Contact ajouté !")

    return contacts
}

// Afficher tous les contacts
func afficherTous(contacts []Contact) {
    if len(contacts) == 0 {
        fmt.Println("Aucun contact.")
        return
    }

    fmt.Println("\n=== Liste des contacts ===")
    for i, contact := range contacts {
        fmt.Printf("%d. %s - %s\n", i+1, contact.Nom, contact.Numero)
    }
}

func main() {
    contacts := []Contact{}  // Slice initiale vide

    fmt.Println("=== Gestion de contacts ===\n")

    // Ajouter 2 contacts
    for i := 1; i <= 2; i++ {
        fmt.Printf("--- Ajout du contact %d/2 ---\n", i)
        contacts = ajouterContact(contacts)
        fmt.Println()
    }

    // Afficher tous les contacts
    afficherTous(contacts)
}
```

### Version améliorée (avec pointeur)

```go
// Version avec pointeur sur slice (moins de copie)
func ajouterContact(contacts *[]Contact) {
    var nom, numero string

    fmt.Print("Nom : ")
    fmt.Scanln(&nom)
    fmt.Print("Numéro : ")
    fmt.Scanln(&numero)

    nouveau := Contact{Nom: nom, Numero: numero}
    *contacts = append(*contacts, nouveau)

    fmt.Println("✓ Contact ajouté !")
}

func main() {
    contacts := []Contact{}

    for i := 1; i <= 2; i++ {
        fmt.Printf("--- Ajout du contact %d/2 ---\n", i)
        ajouterContact(&contacts)  // Passage par pointeur
        fmt.Println()
    }

    afficherTous(contacts)
}
```

### Version avec map (optionnel)

```go
// Version avec map (recherche par nom)
type Annuaire struct {
    contacts map[string]string  // nom -> numéro
}

func (a *Annuaire) Ajouter(nom, numero string) {
    if a.contacts == nil {
        a.contacts = make(map[string]string)
    }
    a.contacts[nom] = numero
}

func (a Annuaire) Afficher() {
    fmt.Println("\n=== Annuaire ===")
    for nom, numero := range a.contacts {
        fmt.Printf("%s : %s\n", nom, numero)
    }
}

func (a Annuaire) Chercher(nom string) (string, bool) {
    numero, existe := a.contacts[nom]
    return numero, existe
}
```

### Tests à effectuer

```bash
# Exécution
go run contacts.go

# Testez les cas :
# ✅ Ajout de 2 contacts
# ✅ Vérifiez l'affichage
# ✅ Essayez avec des noms composés (utilisez bufio si problème)
```

### Pour aller plus loin (optionnel)

1. **Rechercher un contact** par nom
2. **Supprimer un contact**
3. **Sauvegarder dans un fichier** (module 6)
4. **Menu interactif** avec options multiples

```go
// Exemple de menu interactif
func main() {
    contacts := []Contact{}

    for {
        fmt.Println("\n1. Ajouter un contact")
        fmt.Println("2. Afficher tous les contacts")
        fmt.Println("3. Quitter")
        fmt.Print("Choix : ")

        var choix int
        fmt.Scanln(&choix)

        switch choix {
        case 1:
            contacts = ajouterContact(contacts)
        case 2:
            afficherTous(contacts)
        case 3:
            fmt.Println("Au revoir !")
            return
        default:
            fmt.Println("Choix invalide")
        }
    }
}
```

---

## Récapitulatif des acquis

À la fin de ce module, vous savez :

- ✅ Créer des structures imbriquées
- ✅ Comprendre et utiliser les pointeurs (`*`, `&`)
- ✅ Différence entre passage par valeur et par référence
- ✅ Utiliser des slices (dynamiques) plutôt que des tableaux (fixes)
- ✅ Manipuler des maps (dictionnaires clé-valeur)
- ✅ Ajouter des éléments avec `append()`
- ✅ Connaître `len()` et `cap()`

---

## Prochain module (aperçu)

Dans le Module 6, nous découvrirons :
- Les interfaces
- La gestion avancée des erreurs
- La manipulation de fichiers
- Les goroutines (concurrence légère)

---

**Félicitations !** Vous maîtrisez maintenant les structures de données essentielles de Go. 🚀
