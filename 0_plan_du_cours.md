# Plan de cours complet - Langage Go (Golang)

## Module 1 – Introduction et mise en place
**Objectif** : Installer Go, écrire et exécuter son premier programme.

1. Historique et caractéristiques de Go
   1.1. Pourquoi Go a été créé
   1.2. Domaines d’utilisation (backend, CLI, cloud, outils)
2. Installation et configuration
   2.1. Sur Linux / Windows / Mac / Termux (téléphone)
   2.2. Vérification avec `go version`
3. Structure d’un programme Go
   3.1. `package main`
   3.2. `import "fmt"`
   3.3. `func main()`
4. Premier programme : Hello, World!
5. Compilation et exécution
   5.1. `go run` (exécution directe)
   5.2. `go build` (génération d’exécutable)

**TP final** :  
> Écris un programme qui affiche ton prénom et ton âge (valeurs écrites en dur dans le code).  
> Exécute-le avec `go run`, puis génère un exécutable avec `go build` et lance-le.

---

## Module 2 – Types et variables
**Objectif** : Déclarer des variables et manipuler les types de base.

1. Types de base
   1.1. `int`, `float64`, `string`, `bool`
2. Déclaration de variables
   2.1. `var nom type`
   2.2. `:=` (inférence de type)
   2.3. Variables multiples
3. Constantes et `iota`
   3.1. `const`
   3.2. Utilisation simple de `iota` pour des énumérations
4. Conversion de types
   4.1. `int(variable)`, `float64(variable)`, etc.

**TP final** :  
> Crée un mini-calculateur qui :  
> - Déclare deux variables `a` et `b` (ex : `10`, `3`)  
> - Affiche leur somme, différence, produit et quotient (division en `float64`)  
> - Utilise une constante pour le nom du programme.

---

## Module 3 – Contrôle de flux
**Objectif** : Maîtriser les conditions et les boucles.

1. Instructions conditionnelles
   1.1. `if`, `else if`, `else`
   1.2. `switch` simple
2. Boucles
   2.1. `for` classique (comme un `while`)
   2.2. `for range` sur une slice (vue rapidement)
3. Gestion d’erreurs simple
   3.1. Retourner une erreur avec `errors.New()`
   3.2. Vérifier `if err != nil`

**TP final** :  
> Jeu de devinette :  
> - Un nombre secret est fixé dans le code (ex: `42`)  
> - L’utilisateur entre un nombre via `fmt.Scanln()`  
> - Le programme lui dit si c’est trop petit, trop grand ou gagné  
> - Il a droit à 5 essais maximum (boucle `for`).

---

## Module 4 – Fonctions et méthodes
**Objectif** : Structurer le code avec des fonctions réutilisables.

1. Définition de fonctions
   1.1. Paramètres et retour de valeurs
   1.2. Retour multiple (ex: `(int, error)`)
2. Paramètres variadiques (`...int`)
3. Fonctions anonymes et closures
   3.1. Assigner une fonction à une variable
   3.2. Closure simple (compteur)
4. Méthodes sur `struct`
   4.1. Déclarer un `struct`
   4.2. Ajouter une méthode avec un *receiver*

**TP final** :  
> Écris une fonction `IMC(poids float64, taille float64) float64` qui retourne l’IMC.  
> Crée une structure `Personne` avec `Nom`, `Poids`, `Taille`.  
> Ajoute une méthode `AfficherIMC()` à `Personne` qui calcule et affiche l’IMC.  
> Dans `main()`, crée une personne et appelle sa méthode.

---

## Module 5 – Structures et pointeurs
**Objectif** : Manipuler des données complexes et comprendre la mémoire.

1. `struct` avancé
   1.1. Struct imbriqué
2. Pointeurs
   2.1. `*T` et `&variable`
   2.2. Passage par référence vs par valeur
3. Tableaux et slices
   3.1. Différence tableau fixe / slice dynamique
   3.2. `append()`, `len()`, `cap()`
4. Maps
   4.1. Création `make(map[K]V)`
   4.2. Ajout, lecture, suppression

**TP final** :  
> Gestion de contacts :  
> - Structure `Contact` avec `Nom`, `Numero`  
> - Une slice de contacts  
> - Fonctions : `ajouterContact()`, `afficherTous()`  
> - Le programme demande à l’utilisateur d’ajouter 2 contacts, puis les affiche.

---

## Module 6 – Packages et modularité
**Objectif** : Organiser son code en plusieurs fichiers / packages.

1. Packages standards
   1.1. `fmt`, `os`, `strings`, `math/rand`
2. Création d’un package personnalisé
   2.1. Structure d’un package (dossier + fichiers)
   2.2. Exporter avec majuscule
3. Import et alias
   3.1. `import "monpackage"`
   3.2. Alias : `import m "monpackage"`

**TP final** :  
> Crée un dossier `mathutil` avec un fichier `mathutil.go` :  
> - Une fonction `Addition(a, b int) int`  
> - Une fonction `Multiplication(a, b int) int`  
> Dans `main.go`, importe ton package et affiche `Addition(5,3)` et `Multiplication(4,2)`.

---

## Module 7 – Gestion des erreurs (niveau 1)
**Objectif** : Rendre les programmes robustes face aux erreurs.

1. Convention d’erreur en Go
   1.1. Retourner `error`
   1.2. Vérifier `if err != nil`
2. `panic` et `recover`
   2.1. Quand utiliser `panic` (cas exceptionnel)
   2.2. `defer` + `recover` pour rattraper une panique
3. Logging simple
   3.1. `log.Println()` et `log.Fatal()`

**TP final** :  
> Programme qui lit un fichier texte :  
> - Demande le nom du fichier à l’utilisateur  
> - Essaie d’ouvrir le fichier avec `os.Open`  
> - Si le fichier n’existe pas, affiche une erreur claire (pas de `panic`)  
> - Si le fichier existe, affiche son contenu.

---

## Module 8 – Interfaces et polymorphisme
**Objectif** : Comprendre et utiliser les interfaces implicites.

1. Déclaration d’une interface
   1.1. Définir un ensemble de méthodes
2. Implémentation implicite
   2.1. Un type implémente une interface s’il a les méthodes
3. Interface vide `interface{}` / `any`
4. Assertion de type et type switch
   4.1. `v, ok := x.(T)`
   4.2. `switch x.(type)`

**TP final** :  
> Interface `Animal` avec `Parler() string`.  
> Deux structs : `Chien` et `Chat`.  
> Implémente `Parler()` pour chacun.  
> Crée une slice d’`Animal` et fais parler chaque animal dans une boucle.

---

## Module 9 – Concurrence de base
**Objectif** : Lancer des fonctions en parallèle avec goroutines et channels.

1. Goroutines
   1.1. Lancer une fonction avec `go`
   1.2. Attention : le programme principal ne les attend pas
2. Channels (non bufferisés)
   2.1. Créer un canal avec `make(chan T)`
   2.2. Envoyer `ch <- valeur`
   2.3. Recevoir `<-ch`
3. `select` et synchronisation simple
4. `sync.WaitGroup` pour attendre plusieurs goroutines

**TP final** :  
> Simule le téléchargement de 3 fichiers :  
> - Une fonction `telecharger(nom string, duree int)` qui attend `duree` secondes (time.Sleep) puis affiche "Téléchargé : X"  
> - Lance chaque téléchargement dans une goroutine  
> - Utilise `WaitGroup` pour attendre la fin des 3.

---

## Module 10 – Entrées/Sorties et fichiers
**Objectif** : Lire/écrire des fichiers et gérer des arguments CLI.

1. Lecture/écriture de fichiers texte
   1.1. `os.ReadFile()` / `os.WriteFile()`
   1.2. `bufio.Scanner` pour lire ligne par ligne
2. Fichiers CSV
   2.1. Package `encoding/csv`
3. Arguments en ligne de commande
   3.1. `os.Args`
4. Buffers et scanner

**TP final** :  
> Programme qui stocke des notes dans un fichier CSV :  
> - Lit `notes.csv` au démarrage (s’il existe)  
> - Demande à l’utilisateur une note (texte)  
> - Ajoute la note au fichier  
> - Affiche toutes les notes.

---

## Module 11 – Concurrence avancée
**Objectif** : Maîtriser canaux bufferisés et synchronisation fine.

1. Canaux bufferisés
   1.1. `make(chan T, capacité)`
2. Fermeture de canal
   2.1. `close(ch)`
   2.2. Tester avec `v, ok := <-ch`
3. `for range` sur un canal
4. `sync.Mutex` et `sync.RWMutex`

**TP final** :  
> Producteur‑consommateur :  
> - Canal bufferisé de taille 3  
> - Une goroutine produit 5 nombres (0 à 4)  
> - Deux goroutines consomment et affichent ce qu’elles reçoivent  
> - Utilise `WaitGroup` et ferme le canal quand la production est finie.

---

## Module 12 – Package `context`
**Objectif** : Annuler des opérations longues ou imposer un timeout.

1. Création d’un contexte
   1.1. `context.Background()`
   1.2. `context.WithCancel()`, `context.WithTimeout()`
2. Annulation manuelle avec `cancel()`
3. Timeout et deadline
4. Passage de valeurs (notion simple)

**TP final** :  
> Simule une recherche longue (3 secondes de `time.Sleep`).  
> L’utilisateur entre un timeout en secondes.  
> Lance la recherche dans une goroutine.  
> Si le timeout est dépassé, annule et affiche "Trop long".

---

## Module 13 – Web et API
**Objectif** : Créer un serveur HTTP et une API REST simple.

1. Serveur HTTP avec `net/http`
   1.1. `http.HandleFunc()`
   1.2. `http.ListenAndServe()`
2. Routes et handlers
3. JSON
   3.1. `json.Marshal()` / `json.Unmarshal()`
   3.2. Struct tags (`json:"nom"`)

**TP final** :  
> API REST pour gérer une liste de tâches (en mémoire) :  
> - `GET /tasks` → liste des tâches  
> - `POST /tasks` → ajouter une tâche (JSON reçu)  
> - `GET /tasks/{id}` → détail d’une tâche  
> - `DELETE /tasks/{id}` → supprimer

---

## Module 14 – Tests et bonnes pratiques (niveau 1)
**Objectif** : Écrire des tests unitaires et benchmarks.

1. Package `testing`
   1.1. Fichiers `_test.go`
   1.2. Fonctions `TestXxx(t *testing.T)`
2. Benchmarks
   2.1. `BenchmarkXxx(b *testing.B)`
3. Documentation GoDoc
   3.1. Commentaires avant une fonction / package

**TP final** :  
> Ajoute des tests unitaires au mini‑calculateur (Module 2) :  
> - Teste `Addition`, `Multiplication`  
> - Vérifie que 2+2=4, 3×4=12

---

## Module 15 – Génériques (Go 1.18+)
**Objectif** : Écrire des fonctions et structures réutilisables sans interface vide.

1. Syntaxe `[T any]`
2. Contraintes avec interfaces
3. Fonction générique simple (`Max`, `Somme`)
4. Structure générique (`Pile[T]`)

**TP final** :  
> Implémente une fonction `InverseSlice[T any](s []T) []T` qui retourne une slice inversée.  
> Teste avec `[]int`, `[]string`, `[]float64`.

---

## Module 16 – Bonnes pratiques, outils et organisation de projet
**Objectif** : Structurer un projet Go professionnel.

1. Organisation de dossiers
   1.1. `cmd/`, `internal/`, `pkg/`
2. Outils essentiels
   2.1. `go fmt`, `go vet`, `go mod tidy`
3. Compilation croisée
   3.1. `GOOS=linux GOARCH=amd64 go build`
4. Profiling basique (`pprof`)

**TP final** :  
> Restructure le gestionnaire de contacts (Module 5) :  
> - Place `main.go` dans `cmd/contacts/`  
> - Place les fonctions métier dans `pkg/contacts/`  
> - Formate le code avec `go fmt` et exécute `go vet`

---

## Projet final intégrateur (après le module 16)

**Capacités requises** :  
Tous les modules précédents (variables, fonctions, erreurs, fichiers, interfaces, concurrence, context, web, tests, génériques, organisation).

**Énoncé** :  
> Développe un **outil CLI de téléchargement parallèle** avec :  
> - Lecture d’une liste d’URLs depuis un fichier CSV  
> - Téléchargement concurrent des fichiers (goroutines + channel de progression)  
> - Annulation globale possible via `context` (timeout ou Ctrl+C)  
> - Affichage d’une barre de progression simple  
> - Sauvegarde d’un log des téléchargements (succès/échec)  
> - Tests unitaires sur la fonction qui écrit le log  
> - Organisation du projet en `cmd/`, `pkg/`

---


