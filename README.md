# 🚀 Formation Go - Cours complet du langage Golang

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Complet-brightgreen)]()

Bienvenue dans ce cours complet d'apprentissage du langage Go (Golang). Ce dépôt contient un plan de cours structuré, des modules progressifs et des travaux pratiques pour maîtriser Go, de zéro à un niveau avancé.

---

## 📋 Table des matières

- [À propos](#-à-propos)
- [Prérequis](#-prérequis)
- [Structure du cours](#-structure-du-cours)
- [Modules (1 à 13)](#-modules-1-à-13)
- [Modules suivants (aperçu)](#-modules-suivants-aperçu)
- [Projet final](#-projet-final)
- [Installation](#-installation)
- [Comment utiliser ce cours](#-comment-utiliser-ce-cours)
- [TP par module](#-tp-par-module)
- [Ressources complémentaires](#-ressources-complémentaires)

---

## 🎯 À propos

Ce cours est conçu pour vous apprendre Go de manière progressive et pratique. Chaque module combine :

- **Concepts théoriques** expliqués simplement
- **Exemples concrets** et bonnes pratiques
- **Travaux pratiques (TP)** avec corrigés
- **Pièges à éviter** et astuces

**Contenu actuel :** Modules 1 à 13 (fondamentaux, approfondissement, modularité, gestion d'erreurs, interfaces, concurrence de base, entrées/sorties, concurrence avancée, context, web et API)
**À venir :** Modules 14 à 16

---

## 📚 Prérequis

- Connaissances de base en programmation (variables, fonctions, boucles)
- Un ordinateur avec accès à Internet
- Aucune connaissance préalable de Go n'est requise

---

## 🏗️ Structure du cours

Le cours est découpé en **16 modules progressifs** + un **projet final intégrateur** :

```
formation_golang/
├── README.md                          # Ce fichier
├── 01_introduction_et_mise_en_place.md
├── 02_types_et_variables.md
├── 03_controle_de_flux.md
├── 04_fonctions_et_methodes.md
├── 05_structures_et_pointeurs.md
├── 06_packages_et_modularite.md
├── 07_gestion_des_erreurs.md
├── 08_interfaces_et_polymorphisme.md
├── 09_concurrence_base.md
├── 10_entrees_sorties_fichiers.md
├── 11_concurrence_avancee.md
├── 12_package_context.md              ✅ Ajouté
├── 13_web_et_api.md                   ✅ Ajouté
├── 14_tests_et_bonnes_pratiques.md    📝 À venir
├── 15_generiques.md                   📝 À venir
└── 16_organisation_projet.md          📝 À venir
```

---

## 📖 Modules (1 à 13)

### Module 1 – Introduction et mise en place
**Objectif** : Installer Go, écrire et exécuter son premier programme.

| Sous-partie | Description |
|-------------|-------------|
| 1.1 | Historique et caractéristiques de Go |
| 1.2 | Domaines d'utilisation (backend, CLI, cloud, outils) |
| 2 | Installation et configuration (Linux/Windows/Mac/Termux) |
| 3 | Structure d'un programme Go (`package main`, `import`, `func main()`) |
| 4 | Premier programme : Hello, World! |
| 5 | Compilation et exécution (`go run` vs `go build`) |

**TP** : Afficher prénom et âge → exécuter avec `go run`, puis `go build`

---

### Module 2 – Types et variables
**Objectif** : Déclarer des variables et manipuler les types de base.

| Sous-partie | Description |
|-------------|-------------|
| 1 | Types de base : `int`, `float64`, `string`, `bool` |
| 2 | Déclaration de variables (`var`, `:=`, multiples) |
| 3 | Constantes (`const`) et `iota` pour les énumérations |
| 4 | Conversion de types (`int(x)`, `float64(x)`) |

**Pièges à éviter** : Division entière (`5/2 = 2`), conversion `float64` → `int` (troncature)

**TP** : Mini-calculateur (somme, différence, produit, quotient)

---

### Module 3 – Contrôle de flux
**Objectif** : Maîtriser les conditions, les boucles et la gestion d'erreurs simple.

| Sous-partie | Description |
|-------------|-------------|
| 0 | Packages standards : `fmt`, `errors`, `strconv`, `math/rand`, `time` |
| 1 | Conditions : `if`, `else if`, `else`, `switch` |
| 2 | Boucles : `for` classique, `for` comme `while`, `break`, `continue`, `for range` |
| 3 | Gestion d'erreurs : `errors.New()`, `if err != nil`, `strconv.Atoi()` |

**TP** : Jeu de devinette (nombre secret, 5 essais, indices "trop petit/grand")

---

### Module 4 – Fonctions et méthodes
**Objectif** : Structurer le code avec des fonctions réutilisables.

| Sous-partie | Description |
|-------------|-------------|
| 1 | Définition de fonctions (paramètres, retour, retour multiple, retour nommé) |
| 2 | Paramètres variadiques (`...int`) |
| 3 | Fonctions anonymes et closures |
| 4 | Structures (`struct`) et méthodes (receiver par valeur ou par pointeur) |

**TP** : Calcul d'IMC (fonction `IMC()`, structure `Personne`, méthode `AfficherIMC()`)

---

### Module 5 – Structures et pointeurs
**Objectif** : Manipuler des données complexes avec les structures avancées, les pointeurs et la mémoire.

| Sous-partie | Description |
|-------------|-------------|
| 1 | Structures avancées (struct imbriqué, tags JSON, struct anonyme) |
| 2 | Pointeurs (`*T`, `&`, passage par valeur vs référence, `new()`) |
| 3 | Tableaux (fixes) vs Slices (dynamiques) – `append()`, `len()`, `cap()`, slicing |
| 4 | Maps (dictionnaires) – création, ajout, lecture, suppression, vérification |

**TP** : Gestion de contacts (structure `Contact`, slice, `ajouterContact()`, `afficherTous()`)

---

### Module 6 – Packages et modularité
**Objectif** : Organiser son code en plusieurs fichiers et créer des packages réutilisables.

| Sous-partie | Description |
|-------------|-------------|
| 0 | Pourquoi les packages ? (problèmes de code non structuré) |
| 1 | Qu'est-ce qu'un package ? (structure, package main vs autres) |
| 2 | Exporter avec la majuscule (public vs private, règles de casse) |
| 3 | Importer des packages (syntaxe, alias, dot, blank import) |
| 4 | Organisation d'un projet Go (structure simple vs professionnelle, `cmd/`, `internal/`, `pkg/`) |
| 5 | Packages standards couramment utilisés (fmt, os, strings, json, http, etc.) |

**TP** : Création d'un package `mathutil` avec fonctions `Addition`, `Multiplication`, `Factorielle`, `EstPair`

---

### Module 7 – Gestion des erreurs (niveau 1)
**Objectif** : Maîtriser la philosophie Go de gestion des erreurs, utiliser `defer`, `panic`, `recover` et mettre en place du logging.

| Sous-partie | Description |
|-------------|-------------|
| 0 | Rappel : Pas d'exceptions en Go, les erreurs sont des valeurs |
| 1 | Le pattern `if err != nil` (vérification immédiate) |
| 2 | Création d'erreurs : `errors.New()`, `fmt.Errorf()`, wrapping avec `%w` |
| 3 | Erreurs personnalisées avec `struct` |
| 4 | Fonctions essentielles : `errors.Is()`, `errors.As()`, `errors.Unwrap()` |
| 5 | `defer` – exécution différée (fermeture de ressources, LIFO) |
| 6 | `panic` et `recover` (quand les utiliser, récupération) |
| 7 | Logging avec le package `log` (configuration, niveaux, fichier de log) |
| 8 | Bonnes pratiques pour des programmes robustes |

**TP** : Calculateur robuste avec validation des entrées, gestion d'erreurs complète, logging dans un fichier et protection contre les panic

---

### Module 8 – Interfaces et polymorphisme
**Objectif** : Comprendre et utiliser les interfaces, le cœur de la programmation orientée objet en Go.

| Sous-partie | Description |
|-------------|-------------|
| 0 | Pourquoi les interfaces ? (problème de code dupliqué, solution générique) |
| 1 | Qu'est-ce qu'une interface ? (définition, syntaxe, exemple concret) |
| 2 | Interfaces implicites (pas de `implements`, Go détecte automatiquement) |
| 3 | L'interface vide `interface{}` (tous les types l'implémentent) |
| 4 | Interfaces embarquées (composition d'interfaces) |
| 5 | Assertion de type (`x.(T)` avec ou sans vérification) |
| 6 | Type switch (`switch v := x.(type)`) pour interroger le type |
| 7 | Pattern "accept interfaces, return structs" (philosophie Go) |
| 8 | Interfaces et pointeurs vs valeurs (pièges à connaître) |
| 9 | Interfaces du package standard : `Stringer`, `error`, `io.Reader/Writer`, `sort.Interface` |

**TP** : Système de paiement avec interfaces
- Interface `Payeur` (méthodes `Payer()` et `GetNom()`)
- Implémentation : `CarteBancaire`, `PayPal`, `Crypto`
- Fonction générique `EffectuerPaiement(p Payeur, montant float64)`
- Type switch dans `AfficherInfos()` pour détails spécifiques

---

### Module 9 – Concurrence de base
**Objectif** : Comprendre et maîtriser la concurrence en Go avec les goroutines, les channels et le `sync.WaitGroup`.

| Sous-partie | Description |
|-------------|-------------|
| 0 | Pourquoi la concurrence ? (tâches lentes, parallélisme vs concurrence) |
| 1 | Les goroutines (le "go" léger) – syntaxe, légèreté (~2 Ko), piège de terminaison |
| 2 | `sync.WaitGroup` – `Add()`, `Done()`, `Wait()` pour synchroniser |
| 3 | Piège classique : capture de variable de boucle dans les goroutines |
| 4 | Les channels (canaux) – `make(chan T)`, envoi `<-`, réception `<-` |
| 5 | Channels non bufferisés (synchrones) vs bufferisés |
| 6 | Notation de direction : `chan<- T` (envoi), `<-chan T` (réception) |
| 7 | Producteur / Consommateur (exemple concret) |
| 8 | Calcul parallèle (somme des carrés) |
| 9 | Le `select` – attendre sur plusieurs channels |
| 10 | Timeout avec `time.After` et `select` |
| 11 | Fermeture des channels (`close()`), détection de fermeture (`v, ok := <-ch`) |
| 12 | `for range` sur channel (s'arrête automatiquement à la fermeture) |

**Pièges à éviter** :
- Programme principal qui se termine avant les goroutines
- Capture de variable de boucle (`for i := range` avec closure)
- Envoi sur channel fermé (panic)
- Oubli de fermeture d'un channel (récepteur bloqué indéfiniment)

**TP** : Téléchargement parallèle
- Simulation de téléchargement avec `time.Sleep()`
- Lancement de plusieurs téléchargements en parallèle avec goroutines
- Collecte des résultats via channel
- Affichage en temps réel de la progression
- Calcul du temps total gagné par rapport à l'exécution séquentielle

---

### Module 10 – Entrées/Sorties et fichiers
**Objectif** : Lire et écrire des fichiers, utiliser les arguments de ligne de commande, et manipuler les flux d'entrée/sortie standards.

| Sous-partie | Description |
|-------------|-------------|
| 0 | Pourquoi les E/S ? (fichiers, arguments CLI, redirections) |
| 1 | Les flux standards (`os.Stdin`, `os.Stdout`, `os.Stderr`) |
| 2 | Lire depuis l'entrée standard (`fmt.Scan`, `bufio.Scanner`) |
| 3 | Arguments de ligne de commande (`os.Args`, package `flag`) |
| 4 | Lire un fichier (`os.ReadFile`, `os.Open` + `bufio.Scanner`) |
| 5 | Écrire dans un fichier (`os.WriteFile`, `os.Create`, `bufio.Writer`) |
| 6 | Manipulation des chemins (`path/filepath`) |
| 7 | Vérifications et informations sur les fichiers (`os.Stat`) |
| 8 | Redirection et pipes (`exec.Command`) |

**Pièges à éviter** :
- Oublier `defer file.Close()` (fuite de mémoire)
- Ignorer les erreurs de fermeture de fichier
- Utiliser `os.ReadFile()` sur des fichiers énormes
- Ne pas gérer les erreurs de parsing des flags

**TP** : Analyseur de logs
- Fichier de log avec niveaux (INFO, WARN, ERROR)
- Filtrage par niveau via flag `-level`
- Statistiques avec flag `-stats`
- Sortie vers fichier avec `-output`

---

### Module 11 – Concurrence avancée
**Objectif** : Protéger les données partagées avec les mutex, éviter les conditions de course, utiliser les patterns avancés de concurrence.

| Sous-partie | Description |
|-------------|-------------|
| 0 | Pourquoi la concurrence avancée ? (conditions de course, données non protégées) |
| 1 | `sync.Mutex` – exclusion mutuelle (Lock/Unlock, règles d'utilisation) |
| 2 | `sync.RWMutex` – lecture/écriture (RLock/RUnlock, performances) |
| 3 | `sync.Once` – initialisation unique (singleton, ressources coûteuses) |
| 4 | `sync.Map` – map concurrente (Store, Load, Delete, Range) |
| 5 | `sync.Cond` – attente conditionnelle (Wait, Signal, Broadcast) |
| 6 | Détection des conditions de course (`go run -race`) |
| 7 | Éviter les deadlocks (interblocages, bonnes pratiques) |
| 8 | Patterns avancés : Worker Pool, Fan-Out/Fan-In, Pipeline, Rate Limiting, Circuit Breaker |

**Pièges à éviter** :
- Copier un mutex (toujours passer par pointeur)
- Oublier `defer Unlock()` après `Lock()`
- Deadlock par acquisition dans le mauvais ordre

**TP** : Cache concurrent avec expiration
- Structure thread-safe avec `sync.RWMutex`
- Expiration automatique des entrées (TTL)
- Nettoyage périodique en arrière-plan
- Statistiques (hits/miss)
- Rafraîchissement du TTL à la lecture

---

### Module 12 – Package `context` 🆕
**Objectif** : Maîtriser le package `context` pour gérer l'annulation, les timeouts, et la propagation de valeurs.

| Sous-partie | Description |
|-------------|-------------|
| 0 | Pourquoi le package `context` ? (arrêter des goroutines proprement) |
| 1 | Créer des contextes (`Background`, `TODO`, `WithCancel`, `WithTimeout`, `WithDeadline`, `WithValue`) |
| 2 | Annuler des goroutines avec `ctx.Done()` |
| 3 | Timeout et Deadline |
| 4 | Propager des valeurs avec `WithValue` (trace ID, authentification) |
| 5 | `context` avec `net/http` et `database/sql` |
| 6 | Erreurs et vérifications (`ctx.Err()` : `DeadlineExceeded` / `Canceled`) |
| 7 | Patterns avancés (groupe de goroutines annulables) |
| 8 | Pièges et bonnes pratiques (toujours `defer cancel()`) |

**TP** : Téléchargement parallèle annulable
- Timeout global avec flag `-timeout`
- Annulation manuelle par appui sur `Entrée`
- Propagation de l'annulation à tous les téléchargements
- Affichage de la progression en temps réel

---

### Module 13 – Web et API 🆕
**Objectif** : Créer des serveurs HTTP, des API REST, manipuler du JSON, utiliser les middleware.

| Sous-partie | Description |
|-------------|-------------|
| 0 | Pourquoi le web avec Go ? (simplicité, performance, concurrence) |
| 1 | Le package `net/http` – premier serveur, `ResponseWriter`, `Request` |
| 2 | Routes et méthodes HTTP (`r.Method`, `http.ServeMux`) |
| 3 | JSON – `Marshal`/`Unmarshal`, `json.NewEncoder`, tags JSON |
| 4 | API REST complète (CRUD avec stockage mémoire) |
| 5 | Middleware (logging, auth, CORS, recovery, chaînage) |
| 6 | Servir des fichiers statiques (`http.FileServer`) |
| 7 | Serveur HTTP avancé (configuration, graceful shutdown, HTTPS) |
| 8 | Client HTTP (`http.Client`, requêtes avec contexte) |

**TP** : API de gestion de tâches (Todo API)
- CRUD complet sur les tâches
- Middleware : logging + authentification basique
- Filtrage par statut (complétées/non complétées)
- Endpoint `PATCH /todos/{id}/toggle` pour basculer le statut

---

## 📖 Modules suivants (aperçu)

| Module | Titre | Description |
|--------|-------|-------------|
| 14 | Tests et bonnes pratiques | `testing`, table-driven tests, benchmarks, coverage |
| 15 | Génériques (Go 1.18+) | `[T any]`, contraintes, types paramétrés |
| 16 | Organisation de projet | `cmd/`, `pkg/`, `internal/`, compilation croisée, CI/CD |

---

## 🎓 Projet final

**Outil CLI de téléchargement parallèle**

À la fin du cours (module 16), vous développerez un outil complet qui combine toutes les compétences acquises :

- ✅ Lecture d'URLs depuis un fichier CSV (Module 10)
- ✅ Téléchargement concurrent de fichiers (Modules 9 et 11)
- ✅ Protection des compteurs avec mutex (Module 11)
- ✅ Annulation globale avec `context` (Module 12)
- ✅ Barre de progression avec channels (Module 9)
- ✅ Logs des téléchargements (Module 7)
- ✅ Tests unitaires (Module 14)
- ✅ Organisation professionnelle (Module 16)

---

## 🔧 Installation

### Installation de Go

**Linux (Ubuntu/Debian) :**
```bash
sudo apt update
sudo apt install golang-go
```

**Mac :**
```bash
brew install go
```

**Windows :**
Téléchargez l'installateur `.msi` depuis [go.dev/dl](https://go.dev/dl/)

**Termux (Android) :**
```bash
pkg update
pkg install golang
```

### Vérification
```bash
go version
# Output: go version go1.22.0 linux/amd64
```

### Cloner ce dépôt
```bash
git clone https://github.com/andydefer/formation_golang.git
cd formation_golang
```

---

## 💡 Comment utiliser ce cours

1. **Par module** : Suivez l'ordre recommandé (Module 1 → 13, puis la suite)
2. **Pratiquez** : Faites chaque TP **sans regarder la correction** d'abord
3. **Expérimentez** : Modifiez les exemples, testez vos idées
4. **Compilez** : Utilisez `go run` pour tester, `go build` pour produire des exécutables

### Commandes utiles

```bash
# Exécuter un programme
go run mon_fichier.go

# Détecter les conditions de course
go run -race mon_fichier.go

# Détecter les deadlocks (context)
go run mon_fichier.go

# Générer un exécutable
go build mon_fichier.go

# Formater le code automatiquement
go fmt ./...

# Vérifier les erreurs potentielles
go vet ./...

# Initialiser un module
go mod init monmodule

# Ajouter une dépendance
go get github.com/user/package

# Nettoyer les dépendances
go mod tidy

# Lancer les tests (modules 14+)
go test ./...
```

---

## 📚 TP par module (Modules 1 à 13)

| Module | TP | Concepts clés |
|--------|-----|----------------|
| **1** | Afficher prénom et âge | `package main`, `fmt.Println()`, `go run` / `go build` |
| **2** | Mini-calculateur | Variables, constantes, conversion de types (`float64()`) |
| **3** | Jeu de devinette | `if`/`else`, boucle `for`, `fmt.Scanln()`, `break` |
| **4** | Calcul d'IMC | Fonctions, `struct`, méthodes, receiver |
| **5** | Gestion de contacts | Slices, `append()`, maps, pointeurs |
| **6** | Package `mathutil` | Création de package, exportation (majuscule), `go mod init` |
| **7** | Calculateur robuste | Gestion d'erreurs, `defer`, logging, `panic`/`recover` |
| **8** | Système de paiement | Interfaces, implémentation implicite, type switch, polymorphisme |
| **9** | Téléchargement parallèle | Goroutines, `WaitGroup`, channels, `select`, parallélisme |
| **10** | Analyseur de logs | `flag`, `bufio.Scanner`, fichiers, `os.Stdout`/`os.Stderr` |
| **11** | Cache concurrent avec expiration | `sync.RWMutex`, `sync.Once`, TTL, nettoyage automatique |
| **12** | Téléchargement parallèle annulable | `context.WithTimeout`, `context.WithCancel`, propagation, annulation |
| **13** | API Todo (REST) | `net/http`, JSON, middleware, CRUD, routes |

### Exemple de réalisation (Module 12 - Context)

```go
// Téléchargement avec annulation
func telecharger(ctx context.Context, id int, duree time.Duration) error {
    for progression := 0; progression <= 100; progression += 10 {
        select {
        case <-ctx.Done():
            return fmt.Errorf("téléchargement %d annulé: %w", id, ctx.Err())
        case <-time.After(duree / 10):
            fmt.Printf("Fichier %d: %d%%\n", id, progression)
        }
    }
    return nil
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Annulation manuelle
    go func() {
        fmt.Scanln()
        cancel()
    }()

    err := telecharger(ctx, 1, 10*time.Second)
    if err != nil {
        fmt.Println("Erreur:", err)
    }
}
```

### Exemple de réalisation (Module 13 - API Todo)

```go
// Middleware d'authentification
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        apiKey := r.Header.Get("X-API-Key")
        if apiKey != "secret-key" {
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(map[string]string{"error": "Non autorisé"})
            return
        }
        next(w, r)
    }
}

// Handler GET /todos
func getTodos(w http.ResponseWriter, r *http.Request) {
    todosMu.RLock()
    defer todosMu.RUnlock()

    var todoList []Todo
    for _, todo := range todos {
        todoList = append(todoList, todo)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todoList)
}
```

---

## 🔗 Ressources complémentaires

### Documentation officielle
- [Documentation Go](https://go.dev/doc/)
- [Tour of Go (interactif)](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Package documentation](https://pkg.go.dev/)
- [Race Detector](https://go.dev/doc/articles/race_detector)
- [Package context](https://pkg.go.dev/context)
- [Package net/http](https://pkg.go.dev/net/http)

### Livres recommandés
- "The Go Programming Language" - Alan A. A. Donovan & Brian W. Kernighan
- "Go in Action" - William Kennedy
- "Concurrency in Go" - Katherine Cox-Buday

### Outils recommandés
- **VS Code** + extension Go
- **GoLand** (JetBrains)
- **Vim/Neovim** + vim-go
- **Acode** (sur Android avec Termux)
- **curl** pour tester les API

### Communauté
- [Gophers France (Slack)](https://gophers-france.slack.com/)
- [r/golang](https://www.reddit.com/r/golang/)
- [Go Discord](https://discord.gg/golang)

---

## 📈 Progression recommandée

```
Semaine 1   : Module 1 (Introduction) + Module 2 (Types et variables)
Semaine 2   : Module 3 (Contrôle de flux)
Semaine 3   : Module 4 (Fonctions et méthodes)
Semaine 4   : Module 5 (Structures et pointeurs)
Semaine 5   : Module 6 (Packages et modularité)
Semaine 6   : Module 7 (Gestion des erreurs)
Semaine 7   : Module 8 (Interfaces et polymorphisme)
Semaine 8   : Module 9 (Concurrence de base)
Semaine 9   : Module 10 (Entrées/Sorties et fichiers)
Semaine 10  : Module 11 (Concurrence avancée)
Semaine 11  : Module 12 (Package context) 🆕
Semaine 12  : Module 13 (Web et API) 🆕
Semaine 13  : Modules 14-16 (Tests à Organisation) + Projet final
```

---

## 🤝 Contribution

Les corrections, suggestions et améliorations sont les bienvenues !

1. Forkez le projet
2. Créez votre branche (`git checkout -b amelioration/ma-modification`)
3. Committez vos changements (`git commit -m 'feat: ajout de quelque chose'`)
4. Poussez vers la branche (`git push origin amelioration/ma-modification`)
5. Ouvrez une Pull Request

---

## 📝 Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.

---

## ✨ Remerciements

- L'équipe Go pour ce langage exceptionnel
- La communauté Go francophone
- Tous les contributeurs et apprenants

---

**Bon apprentissage ! 🐹✨**

N'hésitez pas à ouvrir une issue si vous avez des questions ou des suggestions.

*Dernière mise à jour : Avril 2026*
