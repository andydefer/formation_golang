# Plan de cours complet – Langage Go (Golang)

## Module 1 – Introduction et mise en place
**Objectif** : Installer Go, écrire et exécuter son premier programme.

1. Historique et caractéristiques de Go
   - 1.1. Pourquoi Go a été créé
   - 1.2. Domaines d'utilisation (backend, CLI, cloud, outils)
2. Installation et configuration
   - 2.1. Sur Linux / Windows / Mac / Termux (téléphone)
   - 2.2. Vérification avec `go version`
3. Structure d'un programme Go
   - 3.1. `package main`
   - 3.2. `import "fmt"`
   - 3.3. `func main()`
4. Premier programme : Hello, World!
5. Compilation et exécution
   - 5.1. `go run` (exécution directe)
   - 5.2. `go build` (génération d'exécutable)

**TP final** :
> Écris un programme qui affiche ton prénom et ton âge (valeurs écrites en dur dans le code).
> Exécute-le avec `go run`, puis génère un exécutable avec `go build` et lance-le.

---

## Module 2 – Types et variables
**Objectif** : Déclarer des variables et manipuler les types de base.

1. Types de base
   - 1.1. `int` – nombres entiers
   - 1.2. `float64` – nombres décimaux
   - 1.3. `string` – texte
   - 1.4. `bool` – vrai/faux
2. Déclaration de variables
   - 2.1. `var nom type` (déclaration complète)
   - 2.2. `:=` (inférence de type, déclaration courte)
   - 2.3. Variables multiples
3. Constantes et `iota`
   - 3.1. `const` – valeurs immuables
   - 3.2. `iota` – énumérations simples
4. Conversion de types
   - 4.1. `int(variable)`, `float64(variable)`
   - 4.2. Piège de la division entière

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
   - 1.1. `if`, `else if`, `else`
   - 1.2. `switch` (avec et sans expression)
2. Boucles
   - 2.1. `for` classique (initialisation; condition; post-itération)
   - 2.2. `for` comme un `while`
   - 2.3. Boucle infinie
   - 2.4. `break` et `continue`
   - 2.5. `for range` (aperçu sur les slices)
3. Gestion d'erreurs simple
   - 3.1. Créer une erreur avec `errors.New()`
   - 3.2. Vérifier `if err != nil` (pattern standard)
   - 3.3. Conversion `string` → `int` avec `strconv.Atoi()`

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
   - 1.1. Paramètres et retour de valeurs
   - 1.2. Retour multiple (ex: `(int, error)`)
   - 1.3. Retour nommé
2. Paramètres variadiques (`...int`)
3. Fonctions anonymes et closures
   - 3.1. Assigner une fonction à une variable
   - 3.2. Closure simple (compteur)
4. Structures (`struct`) et méthodes
   - 4.1. Déclarer un `struct` avec des champs
   - 4.2. Ajouter une méthode avec un *receiver* (par valeur ou par pointeur)

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
   - 2.1. Struct imbriqué (une structure dans une autre)
   - 2.2. Tags JSON (`json:"nom"`)
   - 2.3. Struct anonyme (utilisation ponctuelle)
3. Pointeurs
   - 3.1. `*T` et `&variable` (adresses mémoire)
   - 3.2. Passage par valeur vs par référence
   - 3.3. `new()` pour créer un pointeur vers une valeur zéro
   - 3.4. Pointeurs sur structures et receivers
4. Tableaux et slices
   - 4.1. Tableaux (taille fixe) – usage rare
   - 4.2. Slices (taille dynamique) – usage quotidien
   - 4.3. `append()` pour ajouter des éléments
   - 4.4. `len()` et `cap()` (longueur et capacité)
   - 4.5. Découpage (slicing) : `[debut:fin]`
   - 4.6. `copy()` pour une vraie copie
5. Maps (dictionnaires)
   - 5.1. Création avec `make(map[K]V)` ou littérale
   - 5.2. Ajout, lecture, suppression (`delete()`)
   - 5.3. Vérification d'existence avec `valeur, ok := map["clé"]`
   - 5.4. Parcourir une map avec `for range`

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
   - 1.1. Structure d'un package (dossier + fichiers)
   - 1.2. Le package `main` (point d'entrée)
   - 1.3. Créer un package réutilisable
2. Exporter avec la majuscule
   - 2.1. Règle fondamentale (majuscule = public, minuscule = private)
   - 2.2. Exportation des fonctions, types, champs de struct
   - 2.3. Pièges fréquents (champs de struct oubliés)
3. Importer des packages
   - 3.1. Syntaxe d'importation
   - 3.2. Alias d'importation
   - 3.3. Import "dot" (à éviter)
   - 3.4. Import "blank" (`_`) pour les effets de bord
4. Organisation d'un projet Go
   - 4.1. Structure simple (petit projet)
   - 4.2. Structure professionnelle (`cmd/`, `internal/`, `pkg/`)
   - 4.3. Le fichier `go.mod` (module, dépendances)
   - 4.4. Commandes utiles : `go mod init`, `go get`, `go mod tidy`
5. Packages standards couramment utilisés
   - 5.1. `fmt`, `os`, `io`, `bufio`
   - 5.2. `strings`, `strconv`, `errors`
   - 5.3. `time`, `math`, `math/rand`
   - 5.4. `encoding/json`, `net/http`, `sync`, `context`, `testing`

**TP final** :
> Création d'un package `mathutil` :
> - Fonction `Addition(a, b int) int`
> - Fonction `Multiplication(a, b int) int`
> - Fonction `Factorielle(n int) int`
> - Fonction `EstPair(n int) bool`
> - Dans `main.go`, importer et utiliser le package

---

## Module 7 – Gestion des erreurs (niveau 1)
**Objectif** : Maîtriser la philosophie Go de gestion des erreurs, utiliser `defer`, `panic`, `recover` et mettre en place du logging professionnel.

0. Rappel : Pas d'exceptions en Go
   - Comparaison avec Python/Java (try/catch)
   - Philosophie : les erreurs sont des valeurs
1. Le pattern `if err != nil`
   - 1.1. Vérification immédiate des erreurs
   - 1.2. Regroupement des vérifications
   - 1.3. Style idiomatique (vérification inline)
2. Création d'erreurs
   - 2.1. `errors.New()` – erreur simple
   - 2.2. `fmt.Errorf()` – erreur formatée
   - 2.3. Wrapping avec `%w` (Go 1.13+)
   - 2.4. Erreurs personnalisées avec `struct`
   - 2.5. Bonnes pratiques (préfixe `Err`, exportation)
3. Fonctions essentielles du package `errors`
   - 3.1. `errors.Is()` – vérifier le type d'erreur
   - 3.2. `errors.As()` – extraire une erreur personnalisée
   - 3.3. `errors.Unwrap()` – déballer une erreur
4. `defer` – exécution différée
   - 4.1. Fermeture de ressources (fichiers, connexions)
   - 4.2. Ordre d'exécution LIFO (Last In, First Out)
   - 4.3. Piège classique : évaluation des arguments
   - 4.4. Utilisation avec les mutex (déverrouillage automatique)
5. `panic` et `recover`
   - 5.1. Quand utiliser `panic` (erreurs irrécupérables)
   - 5.2. `panic` – arrêt brutal du programme
   - 5.3. `recover` – reprendre le contrôle après un panic
   - 5.4. Cas d'usage typique : serveur HTTP avec recovery
   - 5.5. Ce qu'il ne faut PAS faire avec panic
6. Logging avec le package `log`
   - 6.1. Utilisation basique (`log.Println`, `log.Printf`, `log.Fatal`, `log.Panic`)
   - 6.2. Configuration du logger (`log.SetFlags`, `log.SetPrefix`, `log.SetOutput`)
   - 6.3. Niveaux de log avec loggers personnalisés
   - 6.4. Écriture des logs dans un fichier
7. Bonnes pratiques pour des programmes robustes
   - 7.1. Toujours traiter les erreurs (ne jamais utiliser `_`)
   - 7.2. Ne pas utiliser `panic` pour les erreurs normales
   - 7.3. Enrichir les erreurs avec du contexte (`fmt.Errorf`)
   - 7.4. Définir les erreurs de package en tête de fichier
   - 7.5. Valider les entrées utilisateur systématiquement

**TP final** :
> Calculateur robuste avec logging :
> - Package `calculator` avec opérations de base (+, -, *, /)
> - Gestion des erreurs : division par zéro, opérateur invalide
> - Validation des entrées utilisateur (nombres valides)
> - Logging des erreurs dans un fichier `app.log`
> - Protection avec `defer` + `recover` pour éviter les crashs
> - Formatage des logs avec date, heure et fichier source

---

## Module 8 – Interfaces et polymorphisme
**Objectif** : Comprendre et utiliser les interfaces, le cœur de la programmation orientée objet en Go.

0. Pourquoi les interfaces ?
   - Problème du code dupliqué
   - Solution : code générique avec interfaces
1. Qu'est-ce qu'une interface ?
   - 1.1. Définition d'une interface (ensemble de signatures de méthodes)
   - 1.2. Syntaxe de base
   - 1.3. Exemple concret : `Forme` avec `Aire()` et `Perimetre()`
2. Les interfaces implicites (le secret de Go)
   - 2.1. Pas de mot-clé `implements`
   - 2.2. Détection automatique par le compilateur
   - 2.3. Avantages : découplage, tests simplifiés, évolution souple
3. L'interface vide `interface{}`
   - 3.1. Tous les types l'implémentent
   - 3.2. Utilité : fonctions génériques (avant Go 1.18)
   - 3.3. Limites : impossible d'appeler des méthodes directement
4. Les interfaces embarquées (composition)
   - 4.1. Composition d'interfaces entre elles
   - 4.2. Exemple : `io.Reader`, `io.Writer`, `io.ReadWriteCloser`
5. Assertion de type (type assertion)
   - 5.1. Syntaxe : `x.(T)` – récupérer le type concret
   - 5.2. Assertion simple (panic si échec)
   - 5.3. Assertion avec vérification : `valeur, ok := x.(T)`
6. Type switch (interrogation de type)
   - 6.1. Syntaxe : `switch v := x.(type)`
   - 6.2. Exemple : analyseur de données variées
   - 6.3. Cas `default` pour les types non gérés
7. Pattern "accept interfaces, return structs"
   - 7.1. La règle d'or de la philosophie Go
   - 7.2. Pourquoi accepter des interfaces en paramètre
   - 7.3. Pourquoi retourner des structs concrets
   - 7.4. Exemple concret : constructeur et fonctions génériques
8. Interfaces et pointeurs vs valeurs
   - 8.1. Règle : méthode sur pointeur ≠ méthode sur valeur
   - 8.2. Tableau récapitulatif (receiver valeur vs pointeur)
   - 8.3. Piège classique à éviter
9. Interfaces du package standard à connaître
   - 9.1. `fmt.Stringer` – contrôle l'affichage avec `String()`
   - 9.2. `error` – l'interface d'erreur (déjà vue)
   - 9.3. `io.Reader` et `io.Writer` – E/S fondamentales
   - 9.4. `sort.Interface` – tri personnalisé (Len, Less, Swap)

**TP final** :
> Système de paiement avec interfaces :
> - Interface `Payeur` avec méthodes `Payer(montant float64) error` et `GetNom() string`
> - Implémentations : `CarteBancaire`, `PayPal`, `Crypto` (avec frais de 2%)
> - Fonction `EffectuerPaiement(p Payeur, montant float64)` générique
> - Fonction `AfficherInfos(p interface{})` avec type switch pour détails spécifiques
> - Gestion des soldes et des erreurs (solde insuffisant)

---

## Module 9 – Concurrence de base
**Objectif** : Comprendre et maîtriser la concurrence en Go avec les goroutines, les channels et le `sync.WaitGroup`.

0. Pourquoi la concurrence ?
   - Problème : les tâches lentes bloquent tout (téléchargements, requêtes API, calculs)
   - Solution : exécution parallèle pour meilleures performances
   - Concurrence ≠ Parallélisme (structure vs exécution)
1. Les goroutines (le "go" léger)
   - 1.1. Qu'est-ce qu'une goroutine ? (fonction qui s'exécute simultanément)
   - 1.2. Syntaxe : `go maFonction()`
   - 1.3. Les goroutines sont légères (~2 Ko vs 1 Mo pour un thread)
   - 1.4. Piège classique : le programme principal se termine trop tôt
2. `sync.WaitGroup` – attendre que les goroutines finissent
   - 2.1. Les trois méthodes : `Add(n)`, `Done()`, `Wait()`
   - 2.2. Pattern standard avec boucle `for`
   - 2.3. Piège classique : capture de la variable de boucle
   - 2.4. Passage par paramètre vs closure
3. Les channels (canaux) – communiquer entre goroutines
   - 3.1. Philosophie : "Don't communicate by sharing memory, share memory by communicating"
   - 3.2. Création : `make(chan T)`
   - 3.3. Envoi : `ch <- valeur`
   - 3.4. Réception : `valeur := <-ch`
   - 3.5. Channels non bufferisés (synchrones) – l'envoi bloque jusqu'à réception
   - 3.6. Channels bufferisés – `make(chan T, n)` – envoi bloque seulement si buffer plein
4. Exemples concrets
   - 4.1. Producteur / Consommateur
   - 4.2. Calcul parallèle (somme des carrés)
   - 4.3. Notation de direction : `chan<- T` (envoi uniquement), `<-chan T` (réception uniquement)
5. Le `select` – attendre sur plusieurs channels
   - 5.1. Syntaxe : `select { case <-ch1: ... case <-ch2: ... }`
   - 5.2. `select` avec timeout (`time.After`)
   - 5.3. `select` avec `default` (non bloquant)
6. Fermeture des channels
   - 6.1. Pourquoi fermer ? (`close(ch)`) – indique la fin des données
   - 6.2. Vérifier si un channel est fermé : `v, ok := <-ch`
   - 6.3. `for range` sur channel (s'arrête automatiquement à la fermeture)
   - 6.4. Règle d'or : c'est l'expéditeur qui ferme, jamais le récepteur

**TP final** :
> Téléchargement parallèle :
> - Simulation de téléchargement avec `time.Sleep()`
> - Lancement de plusieurs téléchargements en parallèle avec goroutines
> - Collecte des résultats via channel
> - Affichage en temps réel de la progression
> - Calcul du temps total gagné par rapport à l'exécution séquentielle

---

## Module 10 – Entrées/Sorties et fichiers
**Objectif** : Lire et écrire des fichiers, utiliser les arguments de ligne de commande, et manipuler les flux d'entrée/sortie standards.

1. Les flux standards (stdin, stdout, stderr)
   - 1.1. `os.Stdin`, `os.Stdout`, `os.Stderr`
   - 1.2. Redirection en ligne de commande (`>`, `2>`, `|`)
2. Lire depuis l'entrée standard
   - 2.1. `fmt.Scan()` – basique
   - 2.2. `bufio.Scanner` – recommandé
3. Arguments de ligne de commande
   - 3.1. `os.Args` – accès brut
   - 3.2. Package `flag` – solution professionnelle
4. Lire un fichier
   - 4.1. `os.ReadFile()` – tout le fichier
   - 4.2. `os.Open()` + `bufio.Scanner` – ligne par ligne
   - 4.3. `defer file.Close()` – toujours fermer
5. Écrire dans un fichier
   - 5.1. `os.WriteFile()` – tout d'un coup
   - 5.2. `os.Create()` + `bufio.Writer` – écriture bufferisée
   - 5.3. Ajouter à la fin (`os.O_APPEND`)
6. Manipulation des chemins (`path/filepath`)
   - 6.1. `filepath.Join()`, `Dir()`, `Base()`, `Ext()`
   - 6.2. `filepath.Walk()` – parcourir un répertoire
7. Vérifications sur les fichiers
   - 7.1. `os.Stat()` – métadonnées
   - 7.2. Vérifier l'existence (`os.IsNotExist()`)

**TP final** :
> Analyseur de logs :
> - Filtrage par niveau (INFO, WARN, ERROR) avec flag `-level`
> - Statistiques avec flag `-stats`
> - Sortie vers fichier avec `-output`

---

## Module 11 – Concurrence avancée
**Objectif** : Protéger les données partagées avec les mutex, éviter les conditions de course, utiliser les patterns avancés de concurrence (worker pool, fan-in/fan-out, pipeline, rate limiting, circuit breaker).

0. Pourquoi la concurrence avancée ?
   - Problème : les données partagées non protégées (conditions de course)
   - Solutions de synchronisation (mutex, RWMutex, Once, Cond)
1. `sync.Mutex` – exclusion mutuelle
   - 1.1. Utilisation basique (Lock/Unlock)
   - 1.2. Règles d'utilisation et pièges à éviter
   - 1.3. Ne jamais copier un mutex
2. `sync.RWMutex` – lecture/écriture
   - 2.1. `RLock()` / `RUnlock()` pour les lectures
   - 2.2. `Lock()` / `Unlock()` pour les écritures
   - 2.3. Quand utiliser RWMutex (lectures fréquentes)
3. `sync.Once` – initialisation unique
   - 3.1. Pattern du singleton
   - 3.2. Initialisation de ressources coûteuses
4. `sync.Map` – map concurrente
   - 4.1. `Store()`, `Load()`, `Delete()`, `Range()`
   - 4.2. Quand l'utiliser (et quand l'éviter)
5. `sync.Cond` – attente conditionnelle
   - 5.1. `Wait()`, `Signal()`, `Broadcast()`
   - 5.2. Producteur-consommateur avec Cond
6. Détection des conditions de course
   - 6.1. `go run -race` – détecteur intégré
   - 6.2. Interpréter les rapports de race
7. Éviter les deadlocks (interblocages)
   - 7.1. Qu'est-ce qu'un deadlock ?
   - 7.2. Outils de détection
   - 7.3. Bonnes pratiques anti-deadlock
8. Patterns avancés de concurrence
   - 8.1. Worker Pool (pool de workers)
   - 8.2. Fan-Out / Fan-In (distribution et agrégation)
   - 8.3. Pipeline (chaînage de traitements)
   - 8.4. Rate Limiting (limitation de débit)
   - 8.5. Circuit Breaker (protection contre les défaillances)

**TP final** :
> Cache concurrent avec expiration :
> - Structure thread-safe avec `sync.RWMutex`
> - Expiration automatique des entrées (TTL)
> - Nettoyage périodique en arrière-plan
> - Statistiques (hits/miss)
> - Rafraîchissement du TTL à la lecture

---

## Module 12 – Package `context`
**Objectif** : Maîtriser le package `context` pour gérer l'annulation, les timeouts, et la propagation de valeurs à travers la chaîne d'appels.

0. Pourquoi le package `context` ?
   - Problème : comment arrêter une goroutine ?
   - Solutions : annulation manuelle, timeout, deadline
1. Créer des contextes
   - 1.1. `context.Background()` et `context.TODO()` (contextes racines)
   - 1.2. Contextes dérivés : `WithCancel`, `WithTimeout`, `WithDeadline`, `WithValue`
   - 1.3. Règle d'or : toujours appeler `cancel()` (defer)
2. Annuler des goroutines
   - 2.1. Pattern basique avec `ctx.Done()`
   - 2.2. Propagation d'annulation en cascade
   - 2.3. Annulation avec `select` dans une boucle
3. Timeout et Deadline
   - 3.1. `WithTimeout` – timeout relatif
   - 3.2. `WithDeadline` – deadline absolue
   - 3.3. Combinaison de timeouts
4. Propager des valeurs avec `WithValue`
   - 4.1. Utilisation basique
   - 4.2. Règles importantes (type personnalisé pour les clés)
   - 4.3. Pattern du trace ID
5. `context` avec les packages standards
   - 5.1. `net/http` – serveur et client
   - 5.2. `database/sql` – requêtes avec contexte
6. Erreurs et vérifications
   - 6.1. `ctx.Err()` : `DeadlineExceeded` ou `Canceled`
   - 6.2. `errors.Is()` pour identifier la cause
   - 6.3. Vérifier l'existence d'une deadline avec `ctx.Deadline()`
7. Patterns avancés
   - 7.1. Groupe de goroutines annulables
   - 7.2. Timeout personnalisé avec channel
   - 7.3. Contexte avec valeur de requête (exemple complet)
8. Pièges et bonnes pratiques
   - 8.1. Oublier d'appeler `cancel()` → fuite mémoire
   - 8.2. Passer un contexte `nil` → panic
   - 8.3. Stocker des mutex dans le contexte
   - 8.4. Règles récapitulatives

**TP final** :
> Téléchargement parallèle annulable :
> - Timeout global avec flag `-timeout`
> - Annulation manuelle par appui sur `Entrée`
> - Propagation de l'annulation à tous les téléchargements
> - Affichage de la progression en temps réel

---

## Module 13 – Web et API
**Objectif** : Créer des serveurs HTTP, des API REST, manipuler du JSON, utiliser les middleware, et comprendre les bases du développement web avec Go.

0. Pourquoi le web avec Go ?
   - Simplicité (pas de framework requis)
   - Performance et concurrence native
   - Déploiement en un seul binaire
1. Le package `net/http` – Les bases
   - 1.1. Premier serveur HTTP (`http.HandleFunc` + `http.ListenAndServe`)
   - 1.2. Comprendre `http.ResponseWriter` et `*http.Request`
   - 1.3. Différentes façons d'écrire une réponse
2. Routes et méthodes HTTP
   - 2.1. Détection des méthodes (`r.Method`)
   - 2.2. Routes paramétrées (sans framework)
   - 2.3. Structure plus propre avec `http.ServeMux`
3. JSON – La base des API modernes
   - 3.1. Encoder (Marshal) – Go → JSON
   - 3.2. Décoder (Unmarshal) – JSON → Go
   - 3.3. Utiliser `json.NewEncoder` (plus efficace)
   - 3.4. Tags JSON avancés (`omitempty`, `-`, `string`)
4. API REST complète
   - 4.1. Structure d'une API (modèle, stockage, handlers)
   - 4.2. Endpoints CRUD (GET, POST, PUT, DELETE)
   - 4.3. Tests avec `curl`
5. Middleware
   - 5.1. Qu'est-ce qu'un middleware ?
   - 5.2. Structure d'un middleware
   - 5.3. Middleware de logging
   - 5.4. Middleware d'authentification
   - 5.5. Middleware CORS
   - 5.6. Middleware de récupération (recovery)
   - 5.7. Chaînage de middleware
6. Servir des fichiers statiques
   - 6.1. Serveur de fichiers simple (`http.FileServer`)
   - 6.2. Single Page Application (SPA)
7. Serveur HTTP avancé
   - 7.1. Configuration personnalisée (`http.Server`)
   - 7.2. Arrêt gracieux (graceful shutdown)
   - 7.3. HTTPS / TLS
8. Client HTTP
   - 8.1. Requêtes simples (`http.Get`, `http.Post`)
   - 8.2. Client configurable (`http.Client`)
   - 8.3. Requête avec contexte et timeout

**TP final** :
> API de gestion de tâches (Todo API) :
> - CRUD complet sur les tâches
> - Middleware : logging + authentification basique
> - Filtrage par statut (complétées/non complétées)
> - Endpoint `PATCH /todos/{id}/toggle` pour basculer le statut

---

## Modules suivants (aperçu)

| Module | Titre | Contenu principal |
|--------|-------|-------------------|
| 14 | Tests et bonnes pratiques | `testing`, table-driven tests, benchmarks, coverage |
| 15 | Génériques (Go 1.18+) | `[T any]`, contraintes, types paramétrés |
| 16 | Organisation de projet | `cmd/`, `pkg/`, `internal/`, compilation croisée, CI/CD |

---

## Projet final intégrateur (après le module 16)

**Capacités requises** :
Tous les modules précédents (variables, fonctions, erreurs, fichiers, interfaces, concurrence de base, concurrence avancée, context, web, tests, génériques, organisation).

**Énoncé** :
> Développe un **outil CLI de téléchargement parallèle** avec :
> - Lecture d'une liste d'URLs depuis un fichier CSV (Module 10)
> - Téléchargement concurrent des fichiers (Module 9)
> - Protection des compteurs avec mutex (Module 11)
> - Annulation globale possible via `context` (Module 12)
> - Affichage d'une barre de progression simple (Module 9)
> - Sauvegarde d'un log des téléchargements (succès/échec) (Module 7)
> - Tests unitaires sur la fonction qui écrit le log (Module 14)
> - Organisation du projet en `cmd/`, `pkg/` (Module 16)

---

**Fin du plan de cours**
