# Module 10 – Entrées/Sorties et fichiers

Bienvenue dans ce dixième module ! À la fin de ce module, vous saurez lire et écrire des fichiers, utiliser les arguments de ligne de commande, et manipuler les flux d'entrée/sortie standards. Vous pourrez ainsi créer des outils en ligne de commande (CLI) complets.

---

## 0. Pourquoi les entrées/sorties ?

Jusqu'à présent, vos programmes lisaient les données depuis le code (variables) ou depuis l'utilisateur via `fmt.Scanln()`. Mais dans la réalité, les données viennent de :

- **Fichiers** : logs, configurations, données CSV/JSON
- **Entrée standard** : redirection depuis d'autres programmes (`cat fichier | mon_prog`)
- **Arguments en ligne de commande** : `mon_prog --input data.txt --output result.txt`

**Problème :** Sans gestion des fichiers et des E/S, vos programmes sont limités à des données en mémoire volatile.

**Solution :** Go fournit des packages puissants pour manipuler les fichiers et les flux : `os`, `io`, `bufio`, `ioutil` (obsolète), `path/filepath`.

---

## 1. Les flux standards (stdin, stdout, stderr)

En Go, trois flux sont toujours ouverts :

| Flux | Variable | Utilisation |
|------|----------|-------------|
| Entrée standard | `os.Stdin` | `fmt.Scan()`, `bufio.Scanner` |
| Sortie standard | `os.Stdout` | `fmt.Print()`, `fmt.Println()` |
| Erreur standard | `os.Stderr` | `fmt.Fprint(os.Stderr, ...)` |

### 1.1. Pourquoi différencier stdout et stderr ?

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Sortie normale
    fmt.Println("Ceci va vers stdout")

    // Sortie d'erreur
    fmt.Fprintln(os.Stderr, "Ceci va vers stderr")
}
```

**En ligne de commande :**
```bash
go run main.go > output.txt      # stdout → output.txt (stderr s'affiche)
go run main.go 2> errors.txt     # stderr → errors.txt
go run main.go > all.txt 2>&1    # stdout ET stderr → all.txt
```

💡 **Bonnes pratiques :**
- Les **résultats normaux** (data, succès) → `stdout`
- Les **messages d'erreur** → `stderr`
- Les **logs** → `stderr` ou fichier dédié

---

## 2. Lire depuis l'entrée standard

### 2.1. Avec `fmt.Scan()` (basique mais limité)

```go
var nom string
var age int

fmt.Print("Entrez votre nom et âge: ")
fmt.Scan(&nom, &age)  // Lit jusqu'au premier espace/blanc
fmt.Printf("Bonjour %s, %d ans\n", nom, age)
```

**Limites :** `fmt.Scan()` ne lit pas les lignes complètes avec espaces.

### 2.2. Avec `bufio.Scanner` (recommandé)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Entrez une phrase: ")
    scanner.Scan()                    // Lit jusqu'au \n
    ligne := scanner.Text()           // Récupère le texte
    fmt.Printf("Vous avez dit: %s\n", ligne)

    // Gestion d'erreur
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Erreur de lecture:", err)
    }
}
```

### 2.3. Lire plusieurs lignes

```go
scanner := bufio.NewScanner(os.Stdin)

fmt.Println("Entrez plusieurs lignes (Ctrl+D pour terminer):")
for scanner.Scan() {
    ligne := scanner.Text()
    if ligne == "exit" {
        break
    }
    fmt.Printf("> %s\n", ligne)
}
```

### 2.4. Lire depuis une redirection

```go
// Compte les lignes de l'entrée standard
scanner := bufio.NewScanner(os.Stdin)
compteur := 0

for scanner.Scan() {
    compteur++
}

fmt.Printf("%d lignes lues\n", compteur)
```

**Utilisation :**
```bash
cat fichier.txt | go run compteur.go
go run compteur.go < fichier.txt
```

---

## 3. Arguments de ligne de commande (`os.Args`)

### 3.1. Syntaxe de base

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // os.Args[0] est le nom du programme
    // os.Args[1:] sont les arguments
    fmt.Println("Arguments:", os.Args)

    if len(os.Args) < 2 {
        fmt.Println("Usage: programme <nom>")
        return
    }

    nom := os.Args[1]
    fmt.Printf("Bonjour, %s!\n", nom)
}
```

**Exécution :**
```bash
go run main.go Alice
# Output: Arguments: [/tmp/go-build.../main Alice]
#         Bonjour, Alice!
```

### 3.2. Parser des flags simples (sans package flag)

```go
func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: ./prog [--verbose] [--output fichier]")
        return
    }

    verbose := false
    outputFile := "resultat.txt"

    for i := 1; i < len(os.Args); i++ {
        switch os.Args[i] {
        case "--verbose", "-v":
            verbose = true
        case "--output", "-o":
            if i+1 < len(os.Args) {
                outputFile = os.Args[i+1]
                i++
            }
        }
    }

    if verbose {
        fmt.Println("Mode verbose activé")
    }
    fmt.Printf("Fichier de sortie: %s\n", outputFile)
}
```

### 3.3. Package `flag` (solution professionnelle)

```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    // Définition des flags
    nom := flag.String("nom", "monde", "Nom à saluer")
    age := flag.Int("age", 0, "Âge de la personne")
    verbose := flag.Bool("verbose", false, "Activer les logs détaillés")

    // Parser les arguments
    flag.Parse()

    if *verbose {
        fmt.Println("Mode verbose activé")
    }

    fmt.Printf("Bonjour, %s!\n", *nom)
    if *age > 0 {
        fmt.Printf("Vous avez %d ans\n", *age)
    }
}
```

**Utilisation :**
```bash
go run main.go -nom Alice -age 30 -verbose
go run main.go -h              # Affiche l'aide
```

| Fonction | Rôle |
|----------|------|
| `flag.String()` | Flag de type string (retourne `*string`) |
| `flag.Int()` | Flag de type int |
| `flag.Bool()` | Flag de type bool |
| `flag.Parse()` | Analyse les arguments |
| `flag.Args()` | Arguments non-flag |

---

## 4. Lire un fichier

### 4.1. Lire tout le fichier (`os.ReadFile`)

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Lecture complète du fichier
    data, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
        return
    }

    fmt.Println("Contenu:")
    fmt.Println(string(data))
}
```

⚠️ **Attention :** `os.ReadFile()` charge TOUT le fichier en mémoire. Pour les gros fichiers, utilisez une lecture ligne par ligne.

### 4.2. Lire ligne par ligne (`bufio.Scanner`)

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Ouvrir le fichier
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
        return
    }
    defer file.Close()  // Toujours fermer le fichier !

    // Scanner ligne par ligne
    scanner := bufio.NewScanner(file)
    numeroLigne := 1

    for scanner.Scan() {
        ligne := scanner.Text()
        fmt.Printf("%d: %s\n", numeroLigne, ligne)
        numeroLigne++
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Erreur de lecture: %v\n", err)
    }
}
```

### 4.3. Lire par blocs (buffers)

```go
file, err := os.Open("gros_fichier.bin")
if err != nil {
    panic(err)
}
defer file.Close()

buffer := make([]byte, 4096)  // Buffer de 4 Ko
for {
    n, err := file.Read(buffer)
    if n > 0 {
        // Traiter les n bytes lus
        fmt.Printf("Lu %d bytes: %x\n", n, buffer[:n])
    }
    if err == io.EOF {
        break
    }
    if err != nil {
        fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
        break
    }
}
```

---

## 5. Écrire dans un fichier

### 5.1. Écrire tout d'un coup (`os.WriteFile`)

```go
contenu := []byte("Hello, fichier!\nLigne 2\n")
err := os.WriteFile("output.txt", contenu, 0644)
// 0644 = permissions : rw-r--r-- (lecture/écriture pour user, lecture pour groupe/autres)

if err != nil {
    fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
}
```

### 5.2. Écrire avec `bufio.Writer`

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("output.txt")  // Crée ou tronque
    if err != nil {
        panic(err)
    }
    defer file.Close()

    writer := bufio.NewWriter(file)

    writer.WriteString("Première ligne\n")
    writer.WriteString("Deuxième ligne\n")
    fmt.Fprintf(writer, "Valeur: %d\n", 42)

    writer.Flush()  // IMPORTANT : vide le buffer vers le fichier
}
```

### 5.3. Ajouter à la fin d'un fichier

```go
file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    panic(err)
}
defer file.Close()

_, err = file.WriteString(time.Now().Format(time.RFC3339) + " - Événement\n")
```

| Flag | Signification |
|------|---------------|
| `os.O_RDONLY` | Lecture seule |
| `os.O_WRONLY` | Écriture seule |
| `os.O_RDWR` | Lecture et écriture |
| `os.O_APPEND` | Ajouter à la fin |
| `os.O_CREATE` | Créer si n'existe pas |
| `os.O_TRUNC` | Tronquer (vider) |

---

## 6. Manipulation des chemins (`path/filepath`)

### 6.1. Fonctions essentielles

```go
package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    // Joindre des chemins (cross-platform)
    chemin := filepath.Join("dossier", "sous-dossier", "fichier.txt")
    fmt.Println(chemin)  // dossier/sous-dossier/fichier.txt (ou \ sous Windows)

    // Extraire les composants
    dir := filepath.Dir(chemin)      // dossier/sous-dossier
    base := filepath.Base(chemin)    // fichier.txt
    ext := filepath.Ext(chemin)      // .txt

    fmt.Printf("Dir: %s, Base: %s, Ext: %s\n", dir, base, ext)

    // Chemin absolu
    abs, _ := filepath.Abs(".")
    fmt.Println("Absolu:", abs)

    // Nettoyer un chemin
    sale := "dossier//sous-dossier/../fichier.txt"
    propre := filepath.Clean(sale)
    fmt.Println(propre)  // dossier/fichier.txt
}
```

### 6.2. Parcourir un répertoire

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    root := "."

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.IsDir() {
            fmt.Printf("[DIR]  %s\n", path)
        } else {
            taille := float64(info.Size()) / 1024
            fmt.Printf("[FILE] %s (%.1f Ko)\n", path, taille)
        }
        return nil
    })

    if err != nil {
        fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
    }
}
```

---

## 7. Vérifications et informations sur les fichiers

```go
// Vérifier si un fichier existe
func fichierExiste(nom string) bool {
    info, err := os.Stat(nom)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

// Obtenir les informations
info, err := os.Stat("fichier.txt")
if err == nil {
    fmt.Println("Nom:", info.Name())
    fmt.Println("Taille:", info.Size(), "bytes")
    fmt.Println("Modification:", info.ModTime())
    fmt.Println("Est un dossier:", info.IsDir())
    fmt.Println("Permissions:", info.Mode())
}

// Créer un dossier
err := os.Mkdir("mon_dossier", 0755)
err = os.MkdirAll("chemin/complet/avec/sous-dossiers", 0755)

// Supprimer
err := os.Remove("fichier.txt")
err = os.RemoveAll("dossier")
```

---

## 8. Redirection et pipes

Go permet de rediriger les flux d'un programme vers un autre.

```go
package main

import (
    "bytes"
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("grep", "error")
    cmd.Stdin = bytes.NewBufferString("line1\nline with error\nline3")

    output, _ := cmd.Output()
    fmt.Println(string(output))  // line with error
}
```

---

## TP final – Module 10 – Analyseur de logs

### Énoncé

Créez un outil CLI qui analyse un fichier de log et produit un rapport.

**Fonctionnalités :**

1. **Arguments en ligne de commande** (package `flag`) :
   - `-file` : chemin du fichier de log (obligatoire)
   - `-level` : niveau minimum (INFO, WARN, ERROR) – défaut: INFO
   - `-output` : fichier de sortie – défaut: stdout
   - `-stats` : afficher les statistiques

2. **Format de log attendu** :
   ```
   2024-01-15 10:30:45 INFO User login success
   2024-01-15 10:31:20 WARN High memory usage
   2024-01-15 10:32:01 ERROR Database connection failed
   ```

3. **Fonctions à implémenter** :
   - Lire le fichier ligne par ligne
   - Filtrer par niveau
   - Compter les occurrences par niveau
   - Écrire le résultat filtré dans le fichier de sortie ou stdout

4. **Sortie attendue** (mode normal) :
   ```
   [INFO] 2024-01-15 10:30:45 User login success
   [WARN] 2024-01-15 10:31:20 High memory usage
   ```

5. **Sortie avec `-stats`** :
   ```
   === STATISTIQUES ===
   INFO: 45 lignes
   WARN: 12 lignes
   ERROR: 3 lignes
   TOTAL: 60 lignes
   === FIN STATS ===
   ```

### Structure attendue

```
module10/
├── go.mod
├── main.go
└── test.log
```

---

### Étapes à suivre

**Étape 1 – Créer la structure de base**

```bash
mkdir module10
cd module10
go mod init module10
touch main.go
```

**Étape 2 – Définir les flags**

```go
package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // Définition des flags
    filePath := flag.String("file", "", "Chemin du fichier de log (obligatoire)")
    level := flag.String("level", "INFO", "Niveau minimum (INFO, WARN, ERROR)")
    outputPath := flag.String("output", "", "Fichier de sortie (défaut: stdout)")
    showStats := flag.Bool("stats", false, "Afficher les statistiques")

    flag.Parse()

    // Vérification
    if *filePath == "" {
        fmt.Fprintf(os.Stderr, "Erreur: le flag -file est obligatoire\n")
        flag.Usage()
        os.Exit(1)
    }
}
```

**Étape 3 – Fonction de parsing d'une ligne**

```go
type LogEntry struct {
    Date    string
    Time    string
    Level   string
    Message string
}

func parseLigne(ligne string) (LogEntry, bool) {
    parts := strings.SplitN(ligne, " ", 4)
    if len(parts) < 4 {
        return LogEntry{}, false
    }

    return LogEntry{
        Date:    parts[0],
        Time:    parts[1],
        Level:   parts[2],
        Message: parts[3],
    }, true
}
```

**Étape 4 – Fonction de filtrage**

```go
func niveauValide(level string, minLevel string) bool {
    niveaux := map[string]int{
        "INFO":  1,
        "WARN":  2,
        "ERROR": 3,
    }
    return niveaux[level] >= niveaux[minLevel]
}
```

**Étape 5 – Lecture et traitement**

```go
func traiterLogs(filePath string, minLevel string) ([]LogEntry, map[string]int, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, nil, err
    }
    defer file.Close()

    var logs []LogEntry
    stats := map[string]int{"INFO": 0, "WARN": 0, "ERROR": 0}
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        ligne := scanner.Text()
        entry, ok := parseLigne(ligne)
        if !ok {
            continue
        }

        stats[entry.Level]++
        if niveauValide(entry.Level, minLevel) {
            logs = append(logs, entry)
        }
    }

    return logs, stats, scanner.Err()
}
```

**Étape 6 – Affichage et écriture**

```go
func ecrireResultats(logs []LogEntry, stats map[string]int, outputPath string, showStats bool) error {
    var writer io.Writer

    if outputPath == "" {
        writer = os.Stdout
    } else {
        file, err := os.Create(outputPath)
        if err != nil {
            return err
        }
        defer file.Close()
        writer = file
    }

    // Écriture des logs filtrés
    for _, entry := range logs {
        fmt.Fprintf(writer, "[%s] %s %s %s\n",
            entry.Level, entry.Date, entry.Time, entry.Message)
    }

    // Écriture des stats si demandé
    if showStats {
        fmt.Fprintln(writer, "\n=== STATISTIQUES ===")
        total := 0
        for _, count := range stats {
            total += count
        }
        fmt.Fprintf(writer, "INFO: %d lignes\n", stats["INFO"])
        fmt.Fprintf(writer, "WARN: %d lignes\n", stats["WARN"])
        fmt.Fprintf(writer, "ERROR: %d lignes\n", stats["ERROR"])
        fmt.Fprintf(writer, "TOTAL: %d lignes\n", total)
        fmt.Fprintln(writer, "=== FIN STATS ===")
    }

    return nil
}
```

**Étape 7 – Assembler dans main()**

```go
func main() {
    // ... flags ...

    logs, stats, err := traiterLogs(*filePath, *level)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
        os.Exit(1)
    }

    if err := ecrireResultats(logs, stats, *outputPath, *showStats); err != nil {
        fmt.Fprintf(os.Stderr, "Erreur d'écriture: %v\n", err)
        os.Exit(1)
    }
}
```

**Étape 8 – Créer un fichier test.log**

```
2024-01-15 10:30:45 INFO User login success
2024-01-15 10:31:20 WARN High memory usage
2024-01-15 10:32:01 ERROR Database connection failed
2024-01-15 10:33:15 INFO File uploaded
2024-01-15 10:34:00 WARN Slow response time
2024-01-15 10:35:30 ERROR API timeout
```

### Tests à effectuer

```bash
# Afficher tous les logs
go run main.go -file test.log

# Filtrer WARN et plus
go run main.go -file test.log -level WARN

# Avec statistiques
go run main.go -file test.log -stats

# Écrire dans un fichier
go run main.go -file test.log -output result.txt

# Avec redirection shell
go run main.go -file test.log > output.txt 2> errors.log
```

---

## Solution complète

```go
package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
    "strings"
)

type LogEntry struct {
    Date    string
    Time    string
    Level   string
    Message string
}

func parseLigne(ligne string) (LogEntry, bool) {
    parts := strings.SplitN(ligne, " ", 4)
    if len(parts) < 4 {
        return LogEntry{}, false
    }
    return LogEntry{
        Date:    parts[0],
        Time:    parts[1],
        Level:   parts[2],
        Message: parts[3],
    }, true
}

func niveauValide(level, minLevel string) bool {
    niveaux := map[string]int{"INFO": 1, "WARN": 2, "ERROR": 3}
    return niveaux[level] >= niveaux[minLevel]
}

func traiterLogs(filePath, minLevel string) ([]LogEntry, map[string]int, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, nil, err
    }
    defer file.Close()

    var logs []LogEntry
    stats := map[string]int{"INFO": 0, "WARN": 0, "ERROR": 0}
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        ligne := scanner.Text()
        entry, ok := parseLigne(ligne)
        if !ok {
            continue
        }
        stats[entry.Level]++
        if niveauValide(entry.Level, minLevel) {
            logs = append(logs, entry)
        }
    }
    return logs, stats, scanner.Err()
}

func ecrireResultats(logs []LogEntry, stats map[string]int, outputPath string, showStats bool) error {
    var writer io.Writer
    if outputPath == "" {
        writer = os.Stdout
    } else {
        file, err := os.Create(outputPath)
        if err != nil {
            return err
        }
        defer file.Close()
        writer = file
    }

    for _, entry := range logs {
        fmt.Fprintf(writer, "[%s] %s %s %s\n", entry.Level, entry.Date, entry.Time, entry.Message)
    }

    if showStats {
        fmt.Fprintln(writer, "\n=== STATISTIQUES ===")
        total := 0
        for _, count := range stats {
            total += count
        }
        fmt.Fprintf(writer, "INFO: %d lignes\n", stats["INFO"])
        fmt.Fprintf(writer, "WARN: %d lignes\n", stats["WARN"])
        fmt.Fprintf(writer, "ERROR: %d lignes\n", stats["ERROR"])
        fmt.Fprintf(writer, "TOTAL: %d lignes\n", total)
        fmt.Fprintln(writer, "=== FIN STATS ===")
    }
    return nil
}

func main() {
    filePath := flag.String("file", "", "Chemin du fichier de log (obligatoire)")
    level := flag.String("level", "INFO", "Niveau minimum (INFO, WARN, ERROR)")
    outputPath := flag.String("output", "", "Fichier de sortie (défaut: stdout)")
    showStats := flag.Bool("stats", false, "Afficher les statistiques")
    flag.Parse()

    if *filePath == "" {
        fmt.Fprintln(os.Stderr, "Erreur: le flag -file est obligatoire")
        flag.Usage()
        os.Exit(1)
    }

    logs, stats, err := traiterLogs(*filePath, *level)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Erreur de lecture: %v\n", err)
        os.Exit(1)
    }

    if err := ecrireResultats(logs, stats, *outputPath, *showStats); err != nil {
        fmt.Fprintf(os.Stderr, "Erreur d'écriture: %v\n", err)
        os.Exit(1)
    }
}
```

---

## Version améliorée (avec `defer` et gestion mémoire)

```go
// Ajoutez un flag --tail pour ne lire que les N dernières lignes
tail := flag.Int("tail", 0, "Afficher seulement les N dernières lignes")

// Dans traiterLogs, pour les gros fichiers
if *tail > 0 {
    // Lecture circulaire (plus efficace pour les gros fichiers)
    return lireDernieresLignes(filePath, *tail, minLevel)
}
```

---

## Récapitulatif des acquis

| Concept | Syntaxe | Utilisation |
|---------|---------|-------------|
| Flux standards | `os.Stdin`, `os.Stdout`, `os.Stderr` | Redirection, pipe |
| Arguments CLI | `os.Args`, `flag` | Paramètres du programme |
| Lire fichier | `os.ReadFile()`, `os.Open()` | Charger des données |
| Lire ligne par ligne | `bufio.NewScanner()` | Fichiers volumineux |
| Écrire fichier | `os.WriteFile()`, `os.Create()` | Sauvegarder des données |
| Ajouter à un fichier | `os.OpenFile()` avec `O_APPEND` | Logs, historique |
| Chemins | `path/filepath` | Cross-platform |
| Infos fichier | `os.Stat()` | Métadonnées |

---

## Prochain module (aperçu)

Dans le Module 11, nous découvrirons :
- La **concurrence avancée** avec `sync.Mutex`
- Les **channels bufferisés** en profondeur
- Les **select** pour orchestrer plusieurs channels
- Les patterns de synchronisation avancés

---

**Félicitations !** Vous savez maintenant lire et écrire des fichiers, parser des arguments en ligne de commande, et créer des outils CLI professionnels. C'est une compétence essentielle pour l'administration système et le développement d'outils ! 📁🚀
