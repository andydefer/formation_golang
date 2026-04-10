# Plan de cours complet – Langage Go (Golang)

## Module 1 – Introduction et mise en place
**Objectif** : Installer Go, écrire et exécuter son premier programme.

1. Historique et caractéristiques de Go
   1.1. Pourquoi Go a été créé
   1.2. Domaines d'utilisation (backend, CLI, cloud, outils)
2. Installation et configuration
   2.1. Sur Linux / Windows / Mac / Termux (téléphone)
   2.2. Vérification avec `go version`
3. Structure d'un programme Go
   3.1. `package main`
   3.2. `import "fmt"`
   3.3. `func main()`
4. Premier programme : Hello, World!
5. Compilation et exécution
   5.1. `go run` (exécution directe)
   5.2. `go build` (génération d'exécutable)

**TP final** :
> Écris un programme qui affiche ton prénom et ton âge (valeurs écrites en dur dans le code).
> Exécute-le avec `go run`, puis génère un exécutable avec `go build` et lance-le.

---

## Module 2 – Types et variables
**Objectif** : Déclarer des variables et manipuler les types de base.

1. Types de base
   1.1. `int` – nombres entiers
   1.2. `float64` – nombres décimaux
   1.3. `string` – texte
   1.4. `bool` – vrai/faux
2. Déclaration de variables
   2.1. `var nom type` (déclaration complète)
   2.2. `:=` (inférence de type, déclaration courte)
   2.3. Variables multiples
3. Constantes et `iota`
   3.1. `const` – valeurs immuables
   3.2. `iota` – énumérations simples
4. Conversion de types
   4.1. `int(variable)`, `float64(variable)`
   4.2. Piège de la division entière

**TP final** :
> Crée un mini-calculateur qui :
> - Déclare deux variables `a` et `b` (ex : `10`, `3`)
> - Affiche leur somme, différence, produit et quotient (division en `float64`)
> - Utilise une constante pour le nom du programme.

---

## Module 3 – Contrôle de flux
**Objectif** : Maîtriser les conditions, les boucles et la gestion d'erreurs simple.

0. Les packages en Go (rappel sur `import`)
   - Packages standards : `fmt`, `errors`, `strconv`, `math/rand`, `time`
   - Documentation (`go doc`, pkg.go.dev)
1. Instructions conditionnelles
   1.1. `if`, `else if`, `else`
   1.2. `switch` (avec et sans expression)
2. Boucles
   2.1. `for` classique (initialisation; condition; post-itération)
   2.2. `for` comme un `while`
   2.3. Boucle infinie
   2.4. `break` et `continue`
   2.5. `for range` (aperçu sur les slices)
3. Gestion d'erreurs simple
   3.1. Créer une erreur avec `errors.New()`
   3.2. Vérifier `if err != nil` (pattern standard)
   3.3. Conversion `string` → `int` avec `strconv.Atoi()`

**TP final** :
> Jeu de devinette :
> - Nombre secret fixé dans le code (ex: `42`)
> - L'utilisateur entre un nombre via `fmt.Scanln()`
> - Le programme indique "trop petit", "trop grand" ou "gagné"
> - Limite de 5 essais maximum (boucle `for`)
> - Si échec après 5 essais, afficher le nombre secret

---

## Module 4 – Fonctions et méthodes
**Objectif** : Structurer le code avec des fonctions réutilisables.

1. Définition de fonctions
   1.1. Paramètres et retour de valeurs
   1.2. Retour multiple (ex: `(int, error)`)
   1.3. Retour nommé
2. Paramètres variadiques (`...int`)
3. Fonctions anonymes et closures
   3.1. Assigner une fonction à une variable
   3.2. Closure simple (compteur)
4. Structures (`struct`) et méthodes
   4.1. Déclarer un `struct` avec des champs
   4.2. Ajouter une méthode avec un *receiver* (par valeur ou par pointeur)

**TP final** :
> Calcul d'IMC :
> - Fonction `IMC(poids float64, taille float64) float64`
> - Structure `Personne` avec `Nom`, `Poids`, `Taille`
> - Méthode `AfficherIMC()` qui calcule et affiche l'IMC avec son interprétation
> - Dans `main()`, créer une personne et appeler sa méthode

---

## Module 5 – Structures et pointeurs
**Objectif** : Manipuler des données complexes avec les structures avancées, les pointeurs et la mémoire.

1. Rappel : La mémoire en Go (valeurs et adresses)
2. Structures avancées (`struct`)
   2.1. Struct imbriqué (une structure dans une autre)
   2.2. Tags JSON (`json:"nom"`)
   2.3. Struct anonyme (utilisation ponctuelle)
3. Pointeurs
   3.1. `*T` et `&variable` (adresses mémoire)
   3.2. Passage par valeur vs par référence
   3.3. `new()` pour créer un pointeur vers une valeur zéro
   3.4. Pointeurs sur structures et receivers
4. Tableaux et slices
   4.1. Tableaux (taille fixe) – usage rare
   4.2. Slices (taille dynamique) – usage quotidien
   4.3. `append()` pour ajouter des éléments
   4.4. `len()` et `cap()` (longueur et capacité)
   4.5. Découpage (slicing) : `[debut:fin]`
   4.6. `copy()` pour une vraie copie
5. Maps (dictionnaires)
   5.1. Création avec `make(map[K]V)` ou littérale
   5.2. Ajout, lecture, suppression (`delete()`)
   5.3. Vérification d'existence avec `valeur, ok := map["clé"]`
   5.4. Parcourir une map avec `for range`

**TP final** :
> Gestion de contacts :
> - Structure `Contact` avec `Nom`, `Numero`
> - Une slice de contacts pour stocker les données
> - Fonction `ajouterContact()` qui demande les infos et ajoute à la slice
> - Fonction `afficherTous()` qui affiche tous les contacts
> - Dans `main()`, demander à l'utilisateur d'ajouter 2 contacts, puis les afficher

---

## Module 6 – Packages et modularité
**Objectif** : Organiser son code en plusieurs fichiers et créer des packages réutilisables.

0. Pourquoi les packages ?
   - Problèmes du code non structuré
   - Avantages de la modularité
1. Qu'est-ce qu'un package ?
   1.1. Structure d'un package (dossier + fichiers)
   1.2. Le package `main` (point d'entrée)
   1.3. Créer un package réutilisable
2. Exporter avec la majuscule
   2.1. Règle fondamentale (majuscule = public, minuscule = private)
   2.2. Exportation des fonctions, types, champs de struct
   2.3. Pièges fréquents (champs de struct oubliés)
3. Importer des packages
   3.1. Syntaxe d'importation
   3.2. Alias d'importation
   3.3. Import "dot" (à éviter)
   3.4. Import "blank" (`_`) pour les effets de bord
4. Organisation d'un projet Go
   4.1. Structure simple (petit projet)
   4.2. Structure professionnelle (`cmd/`, `internal/`, `pkg/`)
   4.3. Le fichier `go.mod` (module, dépendances)
   4.4. Commandes utiles : `go mod init`, `go get`, `go mod tidy`
5. Packages standards couramment utilisés
   5.1. `fmt`, `os`, `io`, `bufio`
   5.2. `strings`, `strconv`, `errors`
   5.3. `time`, `math`, `math/rand`
   5.4. `encoding/json`, `net/http`, `sync`, `context`, `testing`

**TP final** :
> Création d'un package `mathutil` :
> - Fonction `Addition(a, b int) int`
> - Fonction `Multiplication(a, b int) int`
> - Fonction `Factorielle(n int) int`
> - Fonction `EstPair(n int) bool`
> - Dans `main.go`, importer et utiliser le package

---

## Modules suivants (aperçu)

| Module | Titre | Contenu principal |
|--------|-------|-------------------|
| 7 | Gestion des erreurs (niveau 1) | `panic`, `recover`, `defer`, logging |
| 8 | Interfaces et polymorphisme | Interfaces implicites, assertion de type |
| 9 | Concurrence de base | Goroutines, channels, `WaitGroup` |
| 10 | Entrées/Sorties et fichiers | Lire/écrire des fichiers, `os.Args` |
| 11 | Concurrence avancée | Canaux bufferisés, `sync.Mutex` |
| 12 | Package `context` | Annulation, timeout |
| 13 | Web et API | Serveur HTTP, JSON, routes |
| 14 | Tests et bonnes pratiques | `testing`, benchmarks |
| 15 | Génériques (Go 1.18+) | `[T any]`, contraintes |
| 16 | Organisation de projet | `cmd/`, `pkg/`, `go fmt`, compilation croisée |

---

## Projet final intégrateur (après le module 16)

**Capacités requises** :
Tous les modules précédents (variables, fonctions, erreurs, fichiers, interfaces, concurrence, context, web, tests, génériques, organisation).

**Énoncé** :
> Développe un **outil CLI de téléchargement parallèle** avec :
> - Lecture d'une liste d'URLs depuis un fichier CSV
> - Téléchargement concurrent des fichiers (goroutines + channel de progression)
> - Annulation globale possible via `context` (timeout ou Ctrl+C)
> - Affichage d'une barre de progression simple
> - Sauvegarde d'un log des téléchargements (succès/échec)
> - Tests unitaires sur la fonction qui écrit le log
> - Organisation du projet en `cmd/`, `pkg/`

---

**Fin du plan de cours**
