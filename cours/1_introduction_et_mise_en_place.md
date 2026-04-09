
Module 1 – Introduction et mise en place

Bienvenue dans ce cours Go !
À la fin de ce module, vous aurez installé Go sur votre machine, vous comprendrez la structure d’un programme Go, et vous écrirez votre premier programme fonctionnel.

---

1. Historique et caractéristiques de Go

1.1. Pourquoi Go a été créé

Go (appelé aussi Golang) a été créé en 2007 chez Google par Robert Griesemer, Rob Pike et Ken Thompson (ce dernier a participé à la création d’Unix et du langage C).

À l’époque, les développeurs Google étaient frustrés par :

· La lenteur des compilations en C++
· La complexité des programmes concurrents
· Le manque d’outils modernes (gestionnaire de dépendances, formatage automatique)

Objectif de Go :

Un langage simple, rapide à compiler, performant en exécution, avec une gestion native de la concurrence.

Caractéristiques principales :

· Typage statique fort → moins d’erreurs inattendues
· Compilation rapide → un gros projet compile en quelques secondes
· Garbage collector efficace → pas de gestion manuelle de la mémoire
· Goroutines → des “micro‑threads” très légers pour le parallélisme
· Syntaxe simple et minimaliste → on peut lire et comprendre un programme Go rapidement

💡 Piège à éviter : Ne confondez pas Go et Python. Go est compilé, statiquement typé, et très efficace pour les systèmes backend. Python est interprété, dynamique, plus lent.

1.2. Domaines d’utilisation

Go est utilisé partout dans l’industrie, notamment pour :

Domaine Exemples concrets
Backend / API Docker, Kubernetes, Traefik
CLI (outils en ligne de commande) Hugo (générateur de sites), Terraform
Cloud et microservices Uber, Dropbox, Netflix (certains services)
Outils réseau Serveurs, proxies, bases de temps réel
DevOps et SRE Prometheus, Etcd, Consul

✅ Si vous voulez faire du web, des API, ou des outils système, Go est un excellent choix.

---

2. Installation et configuration

2.1. Sur Linux / Windows / Mac / Termux (téléphone)

Linux (Ubuntu/Debian)

```bash
sudo apt update
sudo apt install golang-go
```

Ou installez la dernière version depuis le site officiel :

```bash
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
```

Windows

· Téléchargez le fichier .msi depuis go.dev/dl
· Lancez l’installateur (case à cocher : "Add to PATH")
· Redémarrez votre terminal

Mac

```bash
brew install go
```

ou téléchargez le .pkg officiel.

Termux (Android)

```bash
pkg update
pkg install golang
```

📱 Astuce : Avec Termux + Acode (éditeur de code), vous pouvez coder en Go sur votre téléphone.

2.2. Vérification avec go version

Ouvrez un terminal et tapez :

```bash
go version
```

Vous devriez voir :

```
go version go1.22.0 linux/amd64
```

⚠️ Si la commande n’est pas reconnue, vérifiez votre variable d’environnement PATH.

---

3. Structure d’un programme Go

Un fichier Go (extension .go) suit toujours cette base :

```go
package main

import "fmt"

func main() {
    // Le code s’écrit ici
}
```

3.1. package main

· Un programme Go est organisé en paquets (packages)
· package main est spécial : il indique que ce fichier est le point d’entrée du programme exécutable
· Si vous écrivez une bibliothèque (réutilisable), vous utiliserez un autre nom comme package maths

3.2. import "fmt"

· import sert à inclure d’autres paquets
· fmt est le paquet standard pour les entrées/sorties formatées (afficher, lire)
· On peut importer plusieurs paquets :

```go
import (
    "fmt"
    "math"
)
```

3.3. func main()

· C’est la fonction principale : tout programme commence ici
· Elle ne prend aucun paramètre et ne retourne rien
· Si vous oubliez func main(), Go compilera mais ne produira rien (bibliothèque)

---

4. Premier programme : Hello, World!

Créez un fichier main.go et écrivez :

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Explication ligne par ligne :

1. package main → ce programme peut être exécuté
2. import "fmt" → on utilise le paquet pour afficher du texte
3. func main() → point d’entrée
4. fmt.Println() → affiche une ligne et passe à la ligne suivante

🧪 Testez : Vous pouvez aussi afficher fmt.Print("Hello") (sans retour à la ligne).

---

5. Compilation et exécution

5.1. go run (exécution directe)

```bash
go run main.go
```

Ceci :

· compile le code en mémoire
· l’exécute immédiatement
· ne laisse aucun fichier exécutable sur le disque

✅ Idéal pour tester rapidement pendant le développement.

5.2. go build (génération d’exécutable)

```bash
go build main.go
```

Cela produit :

· Linux/Mac : un fichier main (sans extension)
· Windows : main.exe

Ensuite, lancez-le :

```bash
./main        # Linux/Mac
main.exe      # Windows
```

✅ Idéal pour distribuer ou utiliser le programme sans avoir Go installé.

💡 Différence :
go run = compile + exécute (fichier temporaire)
go build = crée un fichier exécutable permanent

---

TP final – Module 1

Énoncé

Écrivez un programme Go qui affiche votre prénom et votre âge (valeurs écrites en dur dans le code).

Exemple de sortie :

```
Prénom : Aline
Âge : 28 ans
```

Étapes à suivre

1. Créez un dossier monPremierProjet
2. Dans ce dossier, créez un fichier presentation.go
3. Écrivez le code avec :
   · package main
   · import "fmt"
   · une fonction main()
   · deux appels à fmt.Println() ou un seul avec fmt.Printf()

Correction possible

```go
package main

import "fmt"

func main() {
    prenom := "Aline"
    age := 28
    fmt.Println("Prénom :", prenom)
    fmt.Println("Âge :", age, "ans")
}
```

Ou plus concis :

```go
fmt.Printf("Prénom : %s\nÂge : %d ans\n", prenom, age)
```

Exécution et validation

1. Exécutez directement :
   ```bash
   go run presentation.go
   ```
2. Générez un exécutable :
   ```bash
   go build presentation.go
   ```
3. Lancez l’exécutable généré

Suggestions pour tester que tout fonctionne

· Modifiez les valeurs de prenom et age → recompilez → vérifiez que l’affichage change.
· Essayez d’afficher une phrase entière (ex: Je m'appelle Aline et j'ai 28 ans.)
· Testez avec go run puis go build sur un autre ordinateur sans Go installé (l’exécutable doit fonctionner).

---

Félicitations ! Vous avez écrit, compilé et exécuté votre premier programme Go.

➡️ Dans le Module 2, nous apprendrons à manipuler les types et les variables plus en détail.