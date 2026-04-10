# Module 4 – Fonctions

Bienvenue dans ce quatrième module ! À la fin de ce module, vous saurez structurer votre code avec des fonctions réutilisables, comprendre les retours multiples, les fonctions anonymes, et même créer vos propres méthodes sur des structures.

---

## 0. Pourquoi les fonctions ?

Jusqu'ici, tout votre code était dans `main()`. Pour un petit programme, ça va. Mais dès que votre programme grandit :

- **Difficile à lire** (100 lignes dans main)
- **Code dupliqué** (même calcul à plusieurs endroits)
- **Difficile à tester**

Les fonctions résolvent ces problèmes :

```go
// ❌ Sans fonction : code dupliqué
prix1 := 100 * 1.20  // TVA 20%
prix2 := 200 * 1.20
prix3 := 150 * 1.20

// ✅ Avec fonction : on écrit le calcul une fois
func ajouterTVA(prix float64) float64 {
    return prix * 1.20
}
prix1 := ajouterTVA(100)
prix2 := ajouterTVA(200)
prix3 := ajouterTVA(150)
```

---

## 1. Définition de fonctions

### 1.1. Syntaxe de base

```go
func nomDeLaFonction(parametres) typeDeRetour {
    // code
    return valeur
}
```

**Exemple simple :**

```go
func direBonjour() {
    fmt.Println("Bonjour !")
}

func addition(a int, b int) int {
    return a + b
}

// Quand les paramètres ont le même type, on peut les grouper
func soustraction(a, b int) int {
    return a - b
}

func main() {
    direBonjour()                    // Bonjour !
    resultat := addition(5, 3)       // 8
    fmt.Println(resultat)
}
```

### 1.2. Retour multiple (très courant en Go)

Une fonction peut retourner plusieurs valeurs. C'est la manière standard de gérer les erreurs.

```go
import "errors"

func diviser(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division par zéro impossible")
    }
    return a / b, nil  // nil = pas d'erreur
}

func main() {
    resultat, err := diviser(10, 2)
    if err != nil {
        fmt.Println("Erreur :", err)
    } else {
        fmt.Println("Résultat :", resultat)  // 5
    }

    // Test avec division par zéro
    resultat, err = diviser(10, 0)
    if err != nil {
        fmt.Println("Erreur :", err)  // Erreur : division par zéro impossible
    }
}
```

💡 **Vous reconnaissez ce pattern ?** C'est exactement ce que fait `strconv.Atoi()` !

```go
nombre, err := strconv.Atoi("42")   // Retourne (int, error)
```

### 1.3. Retour nommé

Go permet de nommer les valeurs de retour. Utile pour la documentation.

```go
func calculerCercle(rayon float64) (aire, perimetre float64) {
    aire = 3.14 * rayon * rayon
    perimetre = 2 * 3.14 * rayon
    return  // return nu (retourne aire et perimetre)
}

func main() {
    a, p := calculerCercle(5)
    fmt.Printf("Aire: %.2f, Périmètre: %.2f\n", a, p)
}
```

⚠️ À utiliser avec modération. Le retour explicite est souvent plus lisible.

---

## 2. Paramètres variadiques (`...`)

Une fonction peut accepter un nombre variable d'arguments.

```go
func somme(nombres ...int) int {
    total := 0
    for _, n := range nombres {
        total += n
    }
    return total
}

func main() {
    fmt.Println(somme(1, 2))       // 3
    fmt.Println(somme(1, 2, 3))    // 6
    fmt.Println(somme(1, 2, 3, 4)) // 10
    fmt.Println(somme())           // 0 (zéro argument fonctionne aussi)
}
```

**Astuce :** Passer une slice à une fonction variadique

```go
nombres := []int{1, 2, 3, 4, 5}
resultat := somme(nombres...)  // Les trois petits points "déballent" la slice
fmt.Println(resultat)          // 15
```

---

## 3. Fonctions anonymes et closures

### 3.1. Assigner une fonction à une variable

En Go, les fonctions sont des **valeurs** comme les autres. On peut les stocker dans des variables.

```go
func main() {
    // Fonction anonyme assignée à une variable
    carre := func(x int) int {
        return x * x
    }

    fmt.Println(carre(5))  // 25

    // On peut aussi passer une fonction en paramètre
    appliquer := func(f func(int) int, valeur int) int {
        return f(valeur)
    }

    resultat := appliquer(carre, 4)
    fmt.Println(resultat)  // 16
}
```

### 3.2. Closure (fermeture)

Une closure est une fonction qui "capture" des variables de son environnement.

```go
func compteur() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    c1 := compteur()
    c2 := compteur()

    fmt.Println(c1())  // 1
    fmt.Println(c1())  // 2
    fmt.Println(c1())  // 3

    fmt.Println(c2())  // 1 (indépendant !)
    fmt.Println(c2())  // 2
}
```

💡 **Utilité des closures :**
- Générateurs (comme le compteur)
- Fonctions de configuration
- Éviter les variables globales

---

## 4. Structures (`struct`) et méthodes

### 4.1. Déclarer une structure

Une `struct` est un regroupement de champs (comme un objet simple).

```go
// Déclaration d'une structure Personne
type Personne struct {
    Nom   string
    Age   int
    Ville string
}

func main() {
    // Créer une personne (plusieurs façons)

    // Méthode 1 : avec les champs dans l'ordre
    p1 := Personne{"Alice", 30, "Paris"}

    // Méthode 2 : avec les noms des champs (recommandée)
    p2 := Personne{
        Nom:   "Bob",
        Age:   25,
        Ville: "Lyon",
    }

    // Méthode 3 : champ par champ
    var p3 Personne
    p3.Nom = "Charlie"
    p3.Age = 35
    p3.Ville = "Marseille"

    // Accéder aux champs avec .
    fmt.Println(p1.Nom)  // Alice
    fmt.Println(p2.Age)  // 25
    fmt.Println(p3.Ville) // Marseille
}
```

### 4.2. Ajouter des méthodes avec un receiver

Une méthode est une fonction attachée à un type (une `struct` par exemple).

```go
type Rectangle struct {
    Largeur, Hauteur float64
}

// Méthode avec receiver par valeur (copie)
func (r Rectangle) Aire() float64 {
    return r.Largeur * r.Hauteur
}

// Méthode avec receiver par pointeur (modifie l'original)
func (r *Rectangle) Agrandir(facteur float64) {
    r.Largeur *= facteur
    r.Hauteur *= facteur
}

func main() {
    rect := Rectangle{Largeur: 10, Hauteur: 5}

    fmt.Println(rect.Aire())  // 50

    rect.Agrandir(2)
    fmt.Println(rect.Largeur) // 20
    fmt.Println(rect.Hauteur) // 10
    fmt.Println(rect.Aire())  // 200
}
```

**Receiver par valeur vs par pointeur :**

| Receiver | Modifie l'original ? | Performance |
|----------|---------------------|-------------|
| `(r Rectangle)` | ❌ Non (copie) | Copie la structure |
| `(r *Rectangle)` | ✅ Oui | Passe une référence |

💡 **Règle simple :** Si vous modifiez la structure ou si elle est grosse → utilisez `*`. Sinon, la valeur suffit.

### 4.3. Exemple complet : Personne avec méthode

```go
type Personne struct {
    Nom  string
    Age  int
}

// Méthode qui affiche les infos
func (p Personne) SePresenter() {
    fmt.Printf("Bonjour, je m'appelle %s et j'ai %d ans.\n", p.Nom, p.Age)
}

// Méthode qui fête son anniversaire (modifie l'âge)
func (p *Personne) Anniversaire() {
    p.Age++
    fmt.Printf("Joyeux anniversaire %s ! Vous avez maintenant %d ans.\n", p.Nom, p.Age)
}

func main() {
    alice := Personne{Nom: "Alice", Age: 30}
    alice.SePresenter()     // Bonjour, je m'appelle Alice et j'ai 30 ans.
    alice.Anniversaire()    // Joyeux anniversaire Alice ! Vous avez maintenant 31 ans.
}
```

---

## TP final – Module 4 – Calcul d'IMC

### Énoncé

Créez un programme qui calcule l'Indice de Masse Corporelle (IMC) en utilisant :

1. **Une fonction** `IMC(poids float64, taille float64) float64` qui retourne l'IMC
2. **Une structure** `Personne` avec les champs : `Nom`, `Poids`, `Taille`
3. **Une méthode** `AfficherIMC()` attachée à `Personne` qui calcule et affiche l'IMC avec une interprétation

**Rappel formule IMC :** `poids / (taille * taille)` avec taille en mètres

**Interprétation :**
- IMC < 18.5 → "Insuffisance pondérale"
- 18.5 ≤ IMC < 25 → "Corpulence normale"
- 25 ≤ IMC < 30 → "Surpoids"
- IMC ≥ 30 → "Obésité"

### Exemple de sortie attendue :

```
=== Calcul d'IMC ===

Personne : Alice
Poids : 65.0 kg
Taille : 1.65 m
IMC : 23.88 - Corpulence normale
```

### Étapes à suivre

**Étape 1 – Créer le fichier**

Créez un dossier `module4` et un fichier `imc.go`.

**Étape 2 – Structure de base**

```go
package main

import "fmt"

func main() {
    // Votre code ici
}
```

**Étape 3 – Écrire la fonction IMC**

```go
func IMC(poids, taille float64) float64 {
    // Calcule et retourne l'IMC
}
```

**Étape 4 – Déclarer la structure Personne**

```go
type Personne struct {
    Nom   string
    Poids float64 // en kg
    Taille float64 // en mètres
}
```

**Étape 5 – Ajouter la méthode AfficherIMC()**

```go
func (p Personne) AfficherIMC() {
    // Calcule l'IMC avec la fonction IMC()
    // Affiche l'IMC avec 2 décimales
    // Affiche l'interprétation
}
```

**Étape 6 – Dans main(), créer une personne et appeler la méthode**

```go
func main() {
    personne := Personne{
        Nom:    "Alice",
        Poids:  65.0,
        Taille: 1.65,
    }

    personne.AfficherIMC()
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

// Fonction IMC
func IMC(poids, taille float64) float64 {
    return poids / (taille * taille)
}

// Structure Personne
type Personne struct {
    Nom    string
    Poids  float64
    Taille float64
}

// Méthode AfficherIMC
func (p Personne) AfficherIMC() {
    imc := IMC(p.Poids, p.Taille)

    // Interprétation
    var interpretation string
    switch {
    case imc < 18.5:
        interpretation = "Insuffisance pondérale"
    case imc < 25:
        interpretation = "Corpulence normale"
    case imc < 30:
        interpretation = "Surpoids"
    default:
        interpretation = "Obésité"
    }

    // Affichage
    fmt.Println("=== Calcul d'IMC ===")
    fmt.Println()
    fmt.Printf("Personne : %s\n", p.Nom)
    fmt.Printf("Poids : %.1f kg\n", p.Poids)
    fmt.Printf("Taille : %.2f m\n", p.Taille)
    fmt.Printf("IMC : %.2f - %s\n", imc, interpretation)
}

func main() {
    // Création d'une personne
    personne := Personne{
        Nom:    "Alice",
        Poids:  65.0,
        Taille: 1.65,
    }

    personne.AfficherIMC()
}
```

### Variante avec plusieurs personnes

```go
func main() {
    personnes := []Personne{
        {"Alice", 65.0, 1.65},
        {"Bob", 85.0, 1.80},
        {"Charlie", 45.0, 1.70},
        {"Diana", 95.0, 1.60},
    }

    for _, p := range personnes {
        p.AfficherIMC()
        fmt.Println()
    }
}
```

### Tests supplémentaires

```bash
# Exécution
go run imc.go

# Test avec différentes valeurs
# Personne très mince : poids=50, taille=1.70 → IMC ~17.30
# Personne très forte : poids=120, taille=1.70 → IMC ~41.52
```

---

## Pour aller plus loin (optionnel)

### 1. Méthode avec pointeur pour modifier

```go
func (p *Personne) ChangerPoids(nouveauPoids float64) {
    p.Poids = nouveauPoids
    fmt.Printf("%s pèse maintenant %.1f kg\n", p.Nom, p.Poids)
}

// Utilisation
personne.ChangerPoids(70.0)
```

### 2. Validation des entrées

```go
func (p Personne) AfficherIMC() {
    if p.Poids <= 0 || p.Taille <= 0 {
        fmt.Println("Erreur : poids ou taille invalide")
        return
    }
    // ... reste du code
}
```

### 3. Ajouter un champ DateNaissance

```go
import "time"

type Personne struct {
    Nom          string
    Poids        float64
    Taille       float64
    DateNaissance time.Time
}

func (p Personne) Age() int {
    now := time.Now()
    age := now.Year() - p.DateNaissance.Year()
    if now.YearDay() < p.DateNaissance.YearDay() {
        age--
    }
    return age
}
```

---

## Récapitulatif des acquis

À la fin de ce module, vous savez :

- ✅ Définir vos propres fonctions
- ✅ Utiliser le retour multiple (notamment `(valeur, error)`)
- ✅ Créer des fonctions variadiques avec `...`
- ✅ Écrire des fonctions anonymes et des closures
- ✅ Déclarer des structures (`struct`)
- ✅ Ajouter des méthodes avec des receivers (par valeur ou par pointeur)

---

## Prochain module (aperçu)

Dans le Module 5, nous découvrirons :
- Les interfaces
- La gestion avancée des erreurs
- La manipulation de fichiers
- Les goroutines (concurrence légère)

---

**Félicitations !** Vous structurez maintenant votre code comme un vrai développeur Go. 🚀
