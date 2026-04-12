# Module 8 – Interfaces et polymorphisme

Bienvenue dans ce huitième module ! À la fin de ce module, vous saurez comprendre et utiliser les interfaces, le cœur de la programmation orientée objet en Go. Vous maîtriserez le polymorphisme implicite, l'assertion de type, et le pattern "accept interfaces, return structs".

---

## 0. Pourquoi les interfaces ?

Dans les modules précédents, nous avons vu les `struct` pour regrouper des données. Mais comment écrire du code **générique** qui fonctionne avec différents types ?

**Problème :** Sans interfaces, on est obligé de dupliquer le code :

```go
// ❌ Sans interfaces - code dupliqué
func AfficherChien(c Chien) {
    fmt.Println(c.Nom, "dit:", c.Aboyer())
}

func AfficherChat(c Chat) {
    fmt.Println(c.Nom, "dit:", c.Miauler())
}
```

**Solution :** Avec une interface, on peut écrire du code qui fonctionne avec **tout type** qui sait "parler" :

```go
// ✅ Avec interface - code unique
type Parleur interface {
    Parler() string
}

func Afficher(p Parleur) {
    fmt.Println(p.Parler())
}
```

Les interfaces sont la porte d'entrée du **polymorphisme** en Go.

---

## 1. Qu'est-ce qu'une interface ?

Une **interface** est un ensemble de signatures de méthodes. Elle définit **ce qu'un type doit savoir faire**, sans préciser comment.

### 1.1. Syntaxe de base

```go
// Définition d'une interface
type NomDeInterface interface {
    Methode1(param type) typeRetour
    Methode2() typeRetour
}
```

### 1.2. Premier exemple concret

```go
package main

import "fmt"

// Définition de l'interface
type Forme interface {
    Aire() float64
    Perimetre() float64
}

// Structure Rectangle
type Rectangle struct {
    Largeur float64
    Hauteur float64
}

// Le Rectangle implémente Forme (méthode Aire)
func (r Rectangle) Aire() float64 {
    return r.Largeur * r.Hauteur
}

// Le Rectangle implémente Forme (méthode Perimetre)
func (r Rectangle) Perimetre() float64 {
    return 2 * (r.Largeur + r.Hauteur)
}

// Structure Cercle
type Cercle struct {
    Rayon float64
}

// Le Cercle implémente Forme
func (c Cercle) Aire() float64 {
    return 3.14159 * c.Rayon * c.Rayon
}

// Le Cercle implémente Forme
func (c Cercle) Perimetre() float64 {
    return 2 * 3.14159 * c.Rayon
}

// Fonction qui fonctionne avec TOUTE forme
func AfficherDetails(f Forme) {
    fmt.Printf("Aire: %.2f, Périmètre: %.2f\n", f.Aire(), f.Perimetre())
}

func main() {
    r := Rectangle{Largeur: 10, Hauteur: 5}
    c := Cercle{Rayon: 7}

    AfficherDetails(r) // Aire: 50.00, Périmètre: 30.00
    AfficherDetails(c) // Aire: 153.94, Périmètre: 43.98
}
```

💡 **Ce qui est magique :** `Rectangle` et `Cercle` n'ont **jamais dit** qu'ils implémentent `Forme`. Go le détecte automatiquement !

---

## 2. Les interfaces implicites (le secret de Go)

Contrairement à Java ou C#, en Go **on ne déclare pas** qu'un type implémente une interface.

**Règle d'or :** Un type implémente une interface s'il possède **toutes les méthodes** de l'interface.

```go
// Nulle part on écrit "implements Parleur"
type Chat struct {
    Nom string
}

func (c Chat) Parler() string {
    return "Miaou"
}

// Chat implémente automatiquement Parleur
type Parleur interface {
    Parler() string
}

// Donc ceci fonctionne
func faireParler(p Parleur) {
    fmt.Println(p.Parler())
}

func main() {
    monChat := Chat{Nom: "Felix"}
    faireParler(monChat) // Miaou
}
```

### 2.1. Avantages de l'implémentation implicite

- **Découplage total** : Un package peut définir une interface, et un autre package peut l'implémenter sans savoir qu'elle existe
- **Tests simplifiés** : On peut créer des mocks facilement
- **Évolution souple** : On peut ajouter des méthodes sans casser les interfaces existantes

```go
// Package animaux (ne connaît pas l'interface Parleur)
package animaux

type Chien struct {
    Nom string
}

func (c Chien) Aboyer() string {
    return "Wouf"
}

// Package main (définit l'interface localement)
package main

type Parleur interface {
    Parler() string
}

// On peut adapter Chien à Parleur !
func (c Chien) Parler() string {
    return c.Aboyer()
}
```

---

## 3. L'interface vide `interface{}`

L'interface vide ne déclare **aucune méthode**. Donc **TOUS les types** l'implémentent.

```go
var nimporteQuoi interface{}

nimporteQuoi = 42
nimporteQuoi = "hello"
nimporteQuoi = Rectangle{10, 5}
nimporteQuoi = true
```

### 3.1. Utilité : fonctions génériques (avant les génériques)

```go
// Affiche n'importe quelle valeur
func Afficher(v interface{}) {
    fmt.Printf("Valeur: %v, Type: %T\n", v, v)
}

func main() {
    Afficher(42)          // Valeur: 42, Type: int
    Afficher("Bonjour")   // Valeur: Bonjour, Type: string
    Afficher(3.14)        // Valeur: 3.14, Type: float64
}
```

### 3.2. Limites : on ne peut pas appeler de méthodes

```go
func Mauvais(v interface{}) {
    // v.Parler() // ERREUR ! interface{} n'a pas de méthode Parler
}
```

Pour utiliser la valeur, il faut d'abord **retrouver son type** (voir section 5).

---

## 4. Les interfaces embarquées (composition)

Go permet de **composer** des interfaces entre elles.

```go
type Lecteur interface {
    Lire(p []byte) (int, error)
}

type Ecrivain interface {
    Ecrire(p []byte) (int, error)
}

// Interface composée (embarquée)
type LecteurEcrivain interface {
    Lecteur
    Ecrivain
    // On peut ajouter d'autres méthodes
    Fermer() error
}
```

### 4.1. Exemple concret avec `io.Reader` et `io.Writer`

Le package standard utilise intensément ce pattern :

```go
// Package io
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// Interface composée
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

---

## 5. Assertion de type (type assertion)

Quand on a une interface, on peut vouloir **retrouver le type concret** sous-jacent.

### 5.1. Syntaxe de base

```go
var i interface{} = "hello"

// Assertion simple (panic si le type ne correspond pas)
s := i.(string)
fmt.Println(s) // hello

// Assertion avec vérification (sécurisée)
s, ok := i.(string)
if ok {
    fmt.Println("C'est une string:", s)
} else {
    fmt.Println("Ce n'est pas une string")
}
```

### 5.2. Exemple complet

```go
func Decrire(v interface{}) {
    // Assertion sécurisée
    if s, ok := v.(string); ok {
        fmt.Printf("C'est une string de longueur %d: %q\n", len(s), s)
        return
    }
    if i, ok := v.(int); ok {
        fmt.Printf("C'est un entier: %d\n", i)
        return
    }
    if b, ok := v.(bool); ok {
        fmt.Printf("C'est un booléen: %t\n", b)
        return
    }
    fmt.Printf("Type inconnu: %T\n", v)
}

func main() {
    Decrire("Bonjour")  // C'est une string de longueur 7: "Bonjour"
    Decrire(42)         // C'est un entier: 42
    Decrire(true)       // C'est un booléen: true
    Decrire(3.14)       // Type inconnu: float64
}
```

---

## 6. Type switch (interrogation de type)

Le `type switch` est une forme spéciale de `switch` pour les interfaces.

### 6.1. Syntaxe

```go
func Decrire(v interface{}) {
    switch valeur := v.(type) {
    case string:
        fmt.Printf("String: %q (longueur: %d)\n", valeur, len(valeur))
    case int:
        fmt.Printf("Int: %d\n", valeur)
    case bool:
        fmt.Printf("Bool: %t\n", valeur)
    case nil:
        fmt.Println("nil")
    default:
        fmt.Printf("Type inconnu: %T\n", valeur)
    }
}
```

### 6.2. Exemple : analyseur de données variées

```go
func AnalyserDonnees(data interface{}) {
    switch v := data.(type) {
    case []interface{}:
        fmt.Println("C'est une slice de", len(v), "éléments")
        for i, elem := range v {
            fmt.Printf("  Élément %d: ", i)
            AnalyserDonnees(elem) // Récursion
        }
    case map[string]interface{}:
        fmt.Println("C'est une map avec", len(v), "clés")
        for cle, val := range v {
            fmt.Printf("  %s: ", cle)
            AnalyserDonnees(val)
        }
    default:
        fmt.Printf("Valeur simple: %v (%T)\n", v, v)
    }
}

func main() {
    data := []interface{}{
        42,
        "hello",
        map[string]interface{}{
            "nom": "Alice",
            "age": 30,
        },
        []interface{}{1, 2, 3},
    }
    AnalyserDonnees(data)
}
```

---

## 7. Pattern "accept interfaces, return structs"

C'est une des phrases les plus célèbres de la philosophie Go.

### 7.1. La règle

> **"Accept interfaces, return structs"**
>
> *(Acceptez des interfaces, retournez des structs)*

**Explication :**
- En **paramètre** : acceptez le type le plus large possible (interface)
- En **retour** : retournez le type le plus spécifique possible (struct)

### 7.2. Pourquoi ?

```go
// ❌ Mauvais : retourner une interface
func NouveauChien(nom string) Animal {  // Animal est une interface
    return &Chien{Nom: nom}
}

// ✅ Bon : retourner un struct
func NouveauChien(nom string) *Chien {
    return &Chien{Nom: nom}
}
```

**Raisons :**
- Flexibilité pour l'appelant (il peut décider d'utiliser l'interface ou pas)
- Les méthodes spécifiques au struct restent accessibles
- Meilleure performance (pas de lookup d'interface)

### 7.3. Exemple concret

```go
package storage

// Interface exposée (ce que l'utilisateur peut faire)
type Storage interface {
    Save(data []byte) error
    Load(id string) ([]byte, error)
}

// Struct concret (non exporté)
type fileStorage struct {
    path string
}

// Méthodes spécifiques (non accessibles via l'interface)
func (fs *fileStorage) SetPath(path string) {
    fs.path = path
}

// Constructeur : retourne l'interface (ou le struct ?)
func NewFileStorage(path string) *fileStorage {  // ← Retourne le struct !
    return &fileStorage{path: path}
}

// Mais on peut aussi proposer une fonction qui accepte l'interface
func SauvegarderTout(storage Storage, data [][]byte) error {
    for _, d := range data {
        if err := storage.Save(d); err != nil {
            return err
        }
    }
    return nil
}
```

---

## 8. Interfaces et pointeurs vs valeurs

**Piège classique :** Les méthodes sur pointeur vs sur valeur.

```go
type Compteur struct {
    valeur int
}

// Méthode avec receiver POINTEUR
func (c *Compteur) Incrementer() {
    c.valeur++
}

// Méthode avec receiver VALEUR
func (c Compteur) Valeur() int {
    return c.valeur
}

type Incrementeur interface {
    Incrementer()
}

func main() {
    c := Compteur{}

    // ✅ OK : un pointeur implémente l'interface
    var inc Incrementeur = &c
    inc.Incrementer()

    // ❌ ERREUR : une valeur n'implémente PAS l'interface
    // var inc2 Incrementeur = c  // Ne compile pas !
}
```

**Règle à retenir :**

| Receiver | Valeur implémente ? | Pointeur implémente ? |
|----------|-------------------|----------------------|
| `func (t T)` | ✅ OUI | ✅ OUI |
| `func (t *T)` | ❌ NON | ✅ OUI |

💡 **Astuce :** Pour éviter la confusion, utilisez systématiquement des **receivers par pointeur** pour les méthodes qui modifient l'objet.

---

## 9. Interfaces du package standard à connaître

### 9.1. `fmt.Stringer`

La plus célèbre : contrôle comment un type est affiché par `fmt.Println`.

```go
type Stringer interface {
    String() string
}

// Implémentation personnalisée
type Personne struct {
    Nom string
    Age int
}

func (p Personne) String() string {
    return fmt.Sprintf("%s (%d ans)", p.Nom, p.Age)
}

func main() {
    p := Personne{"Alice", 30}
    fmt.Println(p) // Alice (30 ans) - au lieu de {Alice 30}
}
```

### 9.2. `error`

L'interface d'erreur (déjà utilisée au module 7) :

```go
type error interface {
    Error() string
}

// On peut créer ses propres erreurs
type ErreurValidation struct {
    Champ string
    Valeur interface{}
}

func (e ErreurValidation) Error() string {
    return fmt.Sprintf("validation échouée pour %s: %v", e.Champ, e.Valeur)
}
```

### 9.3. `io.Reader` et `io.Writer`

Les interfaces fondamentales pour les E/S :

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Exemple : tout ce qui est lisible (fichier, buffer, connexion réseau)
func Copier(dst io.Writer, src io.Reader) (int64, error) {
    return io.Copy(dst, src)  // Fonction standard
}
```

### 9.4. `sort.Interface`

Pour trier n'importe quelle collection :

```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

// Exemple : trier des personnes par âge
type Personnes []Personne

func (p Personnes) Len() int           { return len(p) }
func (p Personnes) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p Personnes) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
    personnes := Personnes{
        {"Bob", 30},
        {"Alice", 25},
        {"Charlie", 35},
    }
    sort.Sort(personnes)
    fmt.Println(personnes) // Alice (25), Bob (30), Charlie (35)
}
```

---

## TP final – Module 8 – Système de paiement avec interfaces

### Énoncé

Créez un système de paiement qui supporte différents modes de paiement (Carte bancaire, PayPal, Cryptomonnaie) en utilisant les interfaces.

**Partie 1 – Définir l'interface `Payeur`**

```go
type Payeur interface {
    Payer(montant float64) error
    GetNom() string
}
```

**Partie 2 – Implémenter trois modes de paiement**

1. **CarteBancaire** : champs `numero`, `dateExp`, `titulaire`. Méthode `Payer()` → déduit du solde (solde initial 1000€)
2. **PayPal** : champs `email`, `solde`. Méthode `Payer()` → vérifie solde suffisant
3. **Crypto** : champs `adresse`, `solde`. Méthode `Payer()` → avec frais de 2%

**Partie 3 – Fonction de paiement générique**

```go
func EffectuerPaiement(p Payeur, montant float64)
```

Affiche le résultat du paiement (succès/échec)

**Partie 4 – Fonction d'information avec type switch**

```go
func AfficherInfos(p interface{})
```

Utilise un `type switch` pour afficher des détails spécifiques à chaque mode.

**Partie 5 – Dans `main()`**

- Créer une instance de chaque mode de paiement
- Tenter un paiement de 150€ avec chacun
- Afficher les infos détaillées de chaque

### Exemple de sortie attendue

```
=== Tests de paiement ===

Carte bancaire (John Doe): Paiement de 150.00€ effectué. Nouveau solde: 850.00€
PayPal (alice@email.com): Paiement de 150.00€ effectué. Nouveau solde: 350.00€
Crypto (0xABC123): Paiement de 150.00€ + frais 3.00€ = 153.00€. Nouveau solde: 847.00€

=== Informations détaillées ===

[CB] Titulaire: John Doe, Numéro: ****1234
[PayPal] Email: alice@email.com, Solde: 350.00€
[Crypto] Adresse: 0xABC123, Solde: 847.00€
```

### Structure attendue

```
module8/
├── go.mod
├── main.go
└── paiement/
    └── paiement.go
```

---

### Étapes à suivre

**Étape 1 – Initialiser le module**

```bash
mkdir -p module8/paiement
cd module8
go mod init module8
```

**Étape 2 – Créer le package `paiement`**

```go
// paiement/paiement.go
package paiement

import "fmt"

// Payeur est l'interface que tous les modes de paiement doivent implémenter
type Payeur interface {
    Payer(montant float64) error
    GetNom() string
}

// CarteBancaire
type CarteBancaire struct {
    Titulaire string
    Numero    string
    DateExp   string
    Solde     float64
}

func (c *CarteBancaire) Payer(montant float64) error {
    if c.Solde < montant {
        return fmt.Errorf("solde insuffisant: %.2f€ < %.2f€", c.Solde, montant)
    }
    c.Solde -= montant
    return nil
}

func (c *CarteBancaire) GetNom() string {
    return c.Titulaire
}

// PayPal
type PayPal struct {
    Email string
    Solde float64
}

func (p *PayPal) Payer(montant float64) error {
    if p.Solde < montant {
        return fmt.Errorf("solde insuffisant: %.2f€ < %.2f€", p.Solde, montant)
    }
    p.Solde -= montant
    return nil
}

func (p *PayPal) GetNom() string {
    return p.Email
}

// Crypto
type Crypto struct {
    Adresse string
    Solde   float64
}

func (c *Crypto) Payer(montant float64) error {
    frais := montant * 0.02
    total := montant + frais
    if c.Solde < total {
        return fmt.Errorf("solde insuffisant: %.2f€ < %.2f€ (montant + frais)", c.Solde, total)
    }
    c.Solde -= total
    return nil
}

func (c *Crypto) GetNom() string {
    return c.Adresse
}
```

**Étape 3 – Créer le `main.go`**

```go
// main.go
package main

import (
    "fmt"
    "module8/paiement"
)

// EffectuerPaiement fonction générique qui accepte toute interface Payeur
func EffectuerPaiement(p paiement.Payeur, montant float64) {
    fmt.Printf("%s tente un paiement de %.2f€...\n", p.GetNom(), montant)

    err := p.Payer(montant)
    if err != nil {
        fmt.Printf("❌ Échec: %v\n", err)
    } else {
        fmt.Printf("✅ Succès !\n")
    }
}

// AfficherInfos utilise un type switch pour afficher des détails spécifiques
func AfficherInfos(p interface{}) {
    switch v := p.(type) {
    case *paiement.CarteBancaire:
        // Masquer les chiffres du milieu
        numeroMasque := "****" + v.Numero[len(v.Numero)-4:]
        fmt.Printf("[CB] Titulaire: %s, Numéro: %s, Solde: %.2f€\n",
            v.Titulaire, numeroMasque, v.Solde)
    case *paiement.PayPal:
        fmt.Printf("[PayPal] Email: %s, Solde: %.2f€\n", v.Email, v.Solde)
    case *paiement.Crypto:
        fmt.Printf("[Crypto] Adresse: %s, Solde: %.2f€\n", v.Adresse, v.Solde)
    default:
        fmt.Printf("[?] Type inconnu: %T\n", v)
    }
}

func main() {
    fmt.Println("=== Système de paiement ===\n")

    // Création des modes de paiement
    cb := &paiement.CarteBancaire{
        Titulaire: "John Doe",
        Numero:    "1234 5678 9012 3456",
        DateExp:   "12/25",
        Solde:     1000.00,
    }

    pp := &paiement.PayPal{
        Email: "alice@email.com",
        Solde: 500.00,
    }

    crypto := &paiement.Crypto{
        Adresse: "0xABC123",
        Solde:   1000.00,
    }

    // Tests de paiement
    montant := 150.00
    fmt.Printf("--- Paiement de %.2f€ ---\n\n", montant)

    fmt.Println("1. Carte bancaire:")
    EffectuerPaiement(cb, montant)

    fmt.Println("\n2. PayPal:")
    EffectuerPaiement(pp, montant)

    fmt.Println("\n3. Crypto:")
    EffectuerPaiement(crypto, montant)

    // Tentative de paiement avec solde insuffisant
    fmt.Println("\n--- Test solde insuffisant (500€ avec PayPal) ---")
    EffectuerPaiement(pp, 500.00)

    // Affichage des informations détaillées
    fmt.Println("\n=== Informations détaillées ===")
    AfficherInfos(cb)
    AfficherInfos(pp)
    AfficherInfos(crypto)
}
```

**Étape 4 – Exécuter**

```bash
go run main.go
```

### Version alternative avec `Stringer`

```go
// Dans paiement/paiement.go, ajoutez cette méthode
func (c CarteBancaire) String() string {
    return fmt.Sprintf("Carte de %s (****%s)", c.Titulaire, c.Numero[len(c.Numero)-4:])
}

// Dans main.go, vous pouvez alors faire
fmt.Println(cb) // Affiche "Carte de John Doe (****3456)"
```

### Pour aller plus loin

**1. Ajouter un nouveau mode de paiement (virement bancaire)**

```go
type Virement struct {
    IBAN   string
    BIC    string
    Solde  float64
}

func (v *Virement) Payer(montant float64) error {
    if v.Solde < montant {
        return fmt.Errorf("solde insuffisant")
    }
    v.Solde -= montant
    return nil
}

func (v *Virement) GetNom() string {
    return v.IBAN
}
```

**2. Ajouter une fonction de remboursement**

```go
type Remboursable interface {
    Payeur
    Rembourser(montant float64) error
}

// Implémentez Rembourser pour CarteBancaire
func (c *CarteBancaire) Rembourser(montant float64) error {
    c.Solde += montant
    return nil
}
```

---

## Prochain module (aperçu)

Dans le Module 9, nous découvrirons :
- La **concurrence** avec les goroutines (légères et économiques)
- Les **channels** pour communiquer entre goroutines
- Le **WaitGroup** pour synchroniser les tâches parallèles

---

**Félicitations !** Vous maîtrisez maintenant le concept le plus puissant de Go pour l'abstraction et le polymorphisme. Les interfaces vous permettront d'écrire du code générique, testable et élégant. 🎯
