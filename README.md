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
- [Modules (1 à 9)](#-modules-1-à-9)
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

**Contenu actuel :** Modules 1 à 9 (fondamentaux, approfondissement, modularité, gestion d'erreurs, interfaces et concurrence de base)
**À venir :** Modules 10 à 16 (avancé à professionnel)

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
├── 09_concurrence_base.md             # ✅ Disponible
├── 10_entrees_sorties_fichiers.md      # À venir
├── 11_concurrence_avancee.md           # À venir
├── 12_package_context.md               # À venir
├── 13_web_et_api.md                    # À venir
├── 14_tests_et_bonnes_pratiques.md     # À venir
├── 15_generiques.md                    # À venir
└── 16_organisation_projet.md           # À venir
```

---

## 📖 Modules (1 à 9)

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

### Module 9 – Concurrence de base 🆕
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

**Exemple de réalisation :**
```go
func telecharger(id int, duree time.Duration, ch chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(duree)
    ch <- fmt.Sprintf("Fichier %d terminé en %v", id, duree)
}

func main() {
    ch := make(chan string, 5)
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go telecharger(i, time.Duration(rand.Intn(1500)+500)*time.Millisecond, ch, &wg)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for msg := range ch {
        fmt.Println(msg)
    }
}
```

---

## 📖 Modules suivants (aperçu)

| Module | Titre | Description |
|--------|-------|-------------|
| 10 | Entrées/Sorties et fichiers | Lire/écrire des fichiers, `os.Args`, `bufio` |
| 11 | Concurrence avancée | Canaux bufferisés avancés, `sync.Mutex`, `sync.RWMutex` |
| 12 | Package `context` | Annulation de goroutines, timeout, propagation |
| 13 | Web et API | Serveur HTTP, JSON, routes, middleware |
| 14 | Tests et bonnes pratiques | `testing`, table-driven tests, benchmarks, coverage |
| 15 | Génériques (Go 1.18+) | `[T any]`, contraintes, types paramétrés |
| 16 | Organisation de projet | `cmd/`, `pkg/`, `internal/`, `go fmt`, compilation croisée, CI/CD |

---

## 🎓 Projet final

**Outil CLI de téléchargement parallèle**

À la fin du cours (module 16), vous développerez un outil complet qui combine toutes les compétences acquises :

- ✅ Lecture d'URLs depuis un fichier CSV
- ✅ Téléchargement concurrent de fichiers (goroutines)
- ✅ Annulation globale avec `context` (timeout ou Ctrl+C)
- ✅ Barre de progression avec channels
- ✅ Logs des téléchargements (succès/échec)
- ✅ Tests unitaires
- ✅ Organisation professionnelle (`cmd/`, `pkg/`)

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

1. **Par module** : Suivez l'ordre recommandé (Module 1 → 9, puis la suite)
2. **Pratiquez** : Faites chaque TP **sans regarder la correction** d'abord
3. **Expérimentez** : Modifiez les exemples, testez vos idées
4. **Compilez** : Utilisez `go run` pour tester, `go build` pour produire des exécutables

### Commandes utiles

```bash
# Exécuter un programme
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

## 📚 TP par module (Modules 1 à 9)

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
| **9** | Téléchargement parallèle 🆕 | Goroutines, `WaitGroup`, channels, `select`, parallélisme |

### Exemple de réalisation (Module 9)

```go
// Téléchargement parallèle avec goroutines et channels
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func telecharger(id int, duree time.Duration, ch chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    start := time.Now()
    time.Sleep(duree)
    ch <- fmt.Sprintf("Fichier %d téléchargé en %v", id, time.Since(start).Round(time.Millisecond))
}

func main() {
    rand.Seed(time.Now().UnixNano())

    const nbFichiers = 5
    ch := make(chan string, nbFichiers)
    var wg sync.WaitGroup

    fmt.Printf("📥 Téléchargement de %d fichiers...\n\n", nbFichiers)
    startTotal := time.Now()

    for i := 1; i <= nbFichiers; i++ {
        wg.Add(1)
        duree := time.Duration(rand.Intn(1500)+500) * time.Millisecond
        go telecharger(i, duree, ch, &wg)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    termines := 0
    for msg := range ch {
        termines++
        fmt.Printf("[%d/%d] %s\n", termines, nbFichiers, msg)
    }

    fmt.Printf("\n✅ Temps total: %v\n", time.Since(startTotal).Round(time.Millisecond))
}
```

---

## 🔗 Ressources complémentaires

### Documentation officielle
- [Documentation Go](https://go.dev/doc/)
- [Tour of Go (interactif)](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Package documentation](https://pkg.go.dev/)

### Livres recommandés
- "The Go Programming Language" - Alan A. A. Donovan & Brian W. Kernighan
- "Go in Action" - William Kennedy

### Outils recommandés
- **VS Code** + extension Go
- **GoLand** (JetBrains)
- **Vim/Neovim** + vim-go
- **Acode** (sur Android avec Termux)

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
Semaine 8   : Module 9 (Concurrence de base) 🆕
Semaine 9   : Modules 10-11 (Fichiers à Concurrence avancée)
Semaine 10  : Modules 12-13 (Context à Web et API)
Semaine 11  : Modules 14-15 (Tests à Génériques)
Semaine 12  : Module 16 (Organisation) + Projet final
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
