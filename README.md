
# 🚀 Formation Go - Cours complet du langage Golang

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Complet-brightgreen)]()

Bienvenue dans ce cours complet d'apprentissage du langage Go (Golang). Ce dépôt contient un plan de cours structuré, des modules progressifs et des travaux pratiques pour maîtriser Go, de zéro à un niveau avancé.

## 📋 Table des matières

- [À propos](#-à-propos)
- [Prérequis](#-prérequis)
- [Structure du cours](#-structure-du-cours)
- [Modules](#-modules)
- [Projet final](#-projet-final)
- [Installation](#-installation)
- [Comment utiliser ce cours](#-comment-utiliser-ce-cours)
- [Ressources complémentaires](#-ressources-complémentaires)

## 🎯 À propos

Ce cours est conçu pour vous apprendre Go de manière progressive et pratique. Chaque module combine :

- **Concepts théoriques** expliqués simplement
- **Exemples concrets** et bonnes pratiques
- **Travaux pratiques** (TP) avec corrigés
- **Pièges à éviter** et astuces

Le cours couvre les fondamentaux jusqu'aux concepts avancés comme la concurrence, les API REST, les tests et les génériques.

## 📚 Prérequis

- Connaissances de base en programmation (variables, fonctions, boucles)
- Un ordinateur avec accès à Internet
- Aucune connaissance préalable de Go n'est requise

## 🏗️ Structure du cours

Le cours est découpé en **16 modules progressifs** + un **projet final intégrateur** :

```
cours/
├── 0_plan_du_cours.md          # Plan détaillé du cours
├── 1_introduction_et_mise_en_place.md
├── 2_types_et_variables.md
├── 3_controle_des_flux.md
├── 4_fonctions_et_methodes.md
├── 5_structures_et_pointeurs.md
├── 6_packages_et_modularite.md
├── 7_gestion_des_erreurs.md
├── 8_interfaces_et_polymorphisme.md
├── 9_concurrence_base.md
├── 10_entrees_sorties_fichiers.md
├── 11_concurrence_avancee.md
├── 12_package_context.md
├── 13_web_et_api.md
├── 14_tests_et_bonnes_pratiques.md
├── 15_generiques.md
└── 16_organisation_projet.md
```

## 📖 Modules

### Fondamentaux (Modules 1-3)
| Module | Description |
|--------|-------------|
| **Module 1** | Introduction, installation et premier programme "Hello, World!" |
| **Module 2** | Types de base, variables, constantes et conversions |
| **Module 3** | Conditions (if/switch) et boucles (for) |

### Approfondissement (Modules 4-7)
| Module | Description |
|--------|-------------|
| **Module 4** | Fonctions, paramètres, retours multiples, méthodes sur struct |
| **Module 5** | Structures, pointeurs, slices, maps |
| **Module 6** | Packages, modularité et organisation du code |
| **Module 7** | Gestion d'erreurs, panic/recover, logging |

### Avancé (Modules 8-12)
| Module | Description |
|--------|-------------|
| **Module 8** | Interfaces, polymorphisme implicite, assertions de type |
| **Module 9** | Goroutines, channels, WaitGroup |
| **Module 10** | Lecture/écriture fichiers, CSV, arguments CLI |
| **Module 11** | Canaux bufferisés, Mutex, patterns de concurrence |
| **Module 12** | Package context, timeouts, annulation |

### Professionnel (Modules 13-16)
| Module | Description |
|--------|-------------|
| **Module 13** | Serveur HTTP, API REST, JSON |
| **Module 14** | Tests unitaires, benchmarks, documentation |
| **Module 15** | Génériques (Go 1.18+) |
| **Module 16** | Organisation de projet, outils, compilation croisée |

## 🎓 Projet final

**Outil CLI de téléchargement parallèle**

À la fin du cours, vous développerez un outil complet qui combine toutes les compétences acquises :

- ✅ Lecture d'URLs depuis un fichier CSV
- ✅ Téléchargement concurrent de fichiers
- ✅ Annulation globale avec context (timeout/Ctrl+C)
- ✅ Barre de progression
- ✅ Logs des téléchargements
- ✅ Tests unitaires
- ✅ Organisation professionnelle (cmd/, pkg/)

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

## 💡 Comment utiliser ce cours

1. **Par module** : Suivez l'ordre recommandé (Module 1 → 16)
2. **Pratiquez** : Faites chaque TP sans regarder la correction
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

# Lancer les tests
go test ./...

# Télécharger les dépendances
go mod tidy
```

## 📚 TP par module

Chaque module contient un TP pratique :

| Module | TP |
|--------|-----|
| 1 | Afficher prénom et âge |
| 2 | Mini-calculateur |
| 3 | Jeu de devinette |
| 4 | Calcul d'IMC avec structure Personne |
| 5 | Gestion de contacts |
| 6 | Création d'un package mathutil |
| 7 | Lecture de fichier avec gestion d'erreur |
| 8 | Interface Animal (Chien/Chat) |
| 9 | Téléchargement parallèle simulé |
| 10 | Stockage de notes en CSV |
| 11 | Producteur-consommateur |
| 12 | Timeout sur recherche longue |
| 13 | API REST de tâches |
| 14 | Tests unitaires sur le calculateur |
| 15 | Fonction générique InverseSlice |
| 16 | Restructuration du gestionnaire de contacts |

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

## 📈 Progression recommandée

```
Semaine 1-2  : Modules 1-3 (Fondamentaux)
Semaine 3-4  : Modules 4-7 (Approfondissement)
Semaine 5-6  : Modules 8-12 (Avancé)
Semaine 7-8  : Modules 13-16 (Professionnel)
Semaine 9-10 : Projet final
```

## 🤝 Contribution

Les corrections, suggestions et améliorations sont les bienvenues !

1. Forkez le projet
2. Créez votre branche (`git checkout -b amelioration/ma-modification`)
3. Committez vos changements (`git commit -m 'feat: ajout de quelque chose'`)
4. Poussez vers la branche (`git push origin amelioration/ma-modification`)
5. Ouvrez une Pull Request

## 📝 Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.

## ✨ Remerciements

- L'équipe Go pour ce langage exceptionnel
- La communauté Go francophone
- Tous les contributeurs et apprenants

---

**Bon apprentissage ! 🐹✨**

N'hésitez pas à ouvrir une issue si vous avez des questions ou des suggestions.

*Dernière mise à jour : Avril 2026*


Ce README.md est conçu pour :

1. **Présenter clairement** le contenu et la structure du cours
2. **Guider l'apprenant** de l'installation à la réalisation du projet final
3. **Être visuellement attractif** avec des badges, tableaux et émojis
4. **Fournir des ressources** pour approfondir
5. **Encourager la pratique** avec la liste des TP par module
6. **Être professionnel** tout en restant accessible aux débutants
