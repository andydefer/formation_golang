# Module 13 – Web et API

Bienvenue dans ce treizième module ! À la fin de ce module, vous saurez créer des serveurs HTTP, des API REST, manipuler du JSON, utiliser les middleware, et comprendre les bases du développement web avec Go. Vous pourrez ainsi construire des services backend complets.

---

## 0. Pourquoi le web avec Go ?

### 0.1. Les forces de Go pour le web

| Caractéristique | Avantage |
|-----------------|----------|
| **Simplicité** | Pas de framework complexe, la stdlib suffit pour 80% des cas |
| **Performance** | Gère des milliers de requêtes par seconde |
| **Concurrence** | Chaque requête est une goroutine légère |
| **Déploiement** | Un seul binaire, pas de dépendances externes |

### 0.2. Ce qu'on va construire

```
Serveur HTTP → Routes → Handlers → Middleware → API REST
     ↓           ↓          ↓            ↓           ↓
  net/http    /users   fonction   logging/auth   JSON
```

---

## 1. Le package `net/http` – Les bases

### 1.1. Votre premier serveur HTTP

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // Associer une route à un handler
    http.HandleFunc("/", helloHandler)

    // Démarrer le serveur
    fmt.Println("Serveur démarré sur http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

**Testez dans votre navigateur :** `http://localhost:8080`

### 1.2. Comprendre les paramètres du handler

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // w http.ResponseWriter : interface pour écrire la réponse
    // r *http.Request      : structure contenant la requête

    // Lire des données de la requête
    method := r.Method
    url := r.URL.Path
    query := r.URL.Query()

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintf(w, "Méthode: %s, URL: %s, Query: %v", method, url, query)
}
```

### 1.3. Différentes façons d'écrire une réponse

```go
func reponseHandler(w http.ResponseWriter, r *http.Request) {
    // Méthode 1: fmt.Fprintf
    fmt.Fprintf(w, "Message formaté: %d\n", 42)

    // Méthode 2: w.Write avec []byte
    w.Write([]byte("Message en bytes\n"))

    // Méthode 3: définir le status code
    w.WriteHeader(http.StatusCreated) // 201

    // Méthode 4: définir des headers
    w.Header().Set("X-Custom-Header", "ma-valeur")
}
```

---

## 2. Routes et méthodes HTTP

### 2.1. Détection des méthodes HTTP

```go
func userHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        fmt.Fprintln(w, "GET: Récupérer l'utilisateur")
    case http.MethodPost:
        fmt.Fprintln(w, "POST: Créer un utilisateur")
    case http.MethodPut:
        fmt.Fprintln(w, "PUT: Modifier l'utilisateur")
    case http.MethodDelete:
        fmt.Fprintln(w, "DELETE: Supprimer l'utilisateur")
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "Méthode %s non supportée", r.Method)
    }
}

func main() {
    http.HandleFunc("/user", userHandler)
    http.ListenAndServe(":8080", nil)
}
```

### 2.2. Routes paramétrées (sans framework)

```go
func dynamicHandler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    parts := r.URL.Path[1:] // enlève le premier '/'

    switch {
    case path == "/":
        fmt.Fprintln(w, "Accueil")
    case path == "/hello":
        fmt.Fprintln(w, "Hello!")
    default:
        // Paramètre dynamique : /users/123
        if len(parts) > 6 && parts[:6] == "users/" {
            userID := parts[6:]
            fmt.Fprintf(w, "Utilisateur ID: %s\n", userID)
            return
        }
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintln(w, "404 - Page non trouvée")
    }
}
```

### 2.3. Structure plus propre avec ServeMux

```go
func main() {
    mux := http.NewServeMux()

    // Routes exactes
    mux.HandleFunc("/", homeHandler)
    mux.HandleFunc("/about", aboutHandler)
    mux.HandleFunc("/contact", contactHandler)

    // Pattern avec slash final (sous-arbre)
    mux.HandleFunc("/static/", staticHandler)

    // Pattern exact (plus spécifique)
    mux.HandleFunc("/users/", usersHandler)

    http.ListenAndServe(":8080", mux)
}
```

---

## 3. JSON – La base des API modernes

### 3.1. Encoder (Marshal) – Go → JSON

```go
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    Password  string    `json:"-"` // Ignorer ce champ
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    user := User{
        ID:        1,
        Name:      "Alice",
        Email:     "alice@example.com",
        CreatedAt: time.Now(),
        Password:  "secret",
    }

    // Convertir en JSON
    jsonData, err := json.Marshal(user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}
```

**Sortie :**
```json
{
    "id": 1,
    "name": "Alice",
    "email": "alice@example.com",
    "created_at": "2024-01-15T10:30:00Z"
}
```

### 3.2. Décoder (Unmarshal) – JSON → Go

```go
func createUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    // Lire le body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Décoder JSON
    var user User
    err = json.Unmarshal(body, &user)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "JSON invalide: %v", err)
        return
    }

    // Traiter l'utilisateur
    fmt.Printf("Reçu: %+v\n", user)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"status": "created"})
}
```

### 3.3. Utiliser `json.NewEncoder` (plus efficace)

```go
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
    users := []User{
        {ID: 1, Name: "Alice", Email: "alice@example.com"},
        {ID: 2, Name: "Bob", Email: "bob@example.com"},
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // Encodage direct dans la réponse
    err := json.NewEncoder(w).Encode(users)
    if err != nil {
        log.Printf("Erreur d'encodage: %v", err)
    }
}
```

### 3.4. Tags JSON avancés

```go
type Product struct {
    ID          int     `json:"id"`
    Name        string  `json:"name,omitempty"`        // Omettre si vide
    Price       float64 `json:"price"`
    Category    string  `json:"category,omitempty"`
    InStock     bool    `json:"in_stock"`
    SecretCode  string  `json:"-"`                     // Toujours omis
    Discount    float64 `json:"discount,string"`       // Convertir en string
    // Renommer et omitempty
    Description string  `json:"desc,omitempty"`
}
```

---

## 4. API REST complète

### 4.1. Structure d'une API REST

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
    "sync"
)

// Modèle
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Stockage en mémoire (simule une base de données)
var (
    users   = make(map[int]User)
    nextID  = 1
    usersMu sync.RWMutex
)

// Handlers
func usersHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getAllUsers(w, r)
    case http.MethodPost:
        createUser(w, r)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(map[string]string{"error": "Méthode non supportée"})
    }
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    // Extraire l'ID de l'URL : /users/123
    parts := strings.Split(r.URL.Path, "/")
    if len(parts) != 3 {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(parts[2])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "ID invalide"})
        return
    }

    switch r.Method {
    case http.MethodGet:
        getUserByID(w, r, id)
    case http.MethodPut:
        updateUser(w, r, id)
    case http.MethodDelete:
        deleteUser(w, r, id)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
    usersMu.RLock()
    defer usersMu.RUnlock()

    userList := make([]User, 0, len(users))
    for _, u := range users {
        userList = append(userList, u)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(userList)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "JSON invalide"})
        return
    }

    usersMu.Lock()
    defer usersMu.Unlock()

    user.ID = nextID
    nextID++
    users[user.ID] = user

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func getUserByID(w http.ResponseWriter, r *http.Request, id int) {
    usersMu.RLock()
    defer usersMu.RUnlock()

    user, exists := users[id]
    if !exists {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Utilisateur non trouvé"})
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request, id int) {
    var updated User
    if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "JSON invalide"})
        return
    }

    usersMu.Lock()
    defer usersMu.Unlock()

    if _, exists := users[id]; !exists {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Utilisateur non trouvé"})
        return
    }

    updated.ID = id
    users[id] = updated

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updated)
}

func deleteUser(w http.ResponseWriter, r *http.Request, id int) {
    usersMu.Lock()
    defer usersMu.Unlock()

    if _, exists := users[id]; !exists {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Utilisateur non trouvé"})
        return
    }

    delete(users, id)
    w.WriteHeader(http.StatusNoContent) // 204
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/users", usersHandler)
    mux.HandleFunc("/users/", userHandler)

    log.Println("Serveur démarré sur :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
```

### 4.2. Tester l'API avec curl

```bash
# GET - Récupérer tous les utilisateurs
curl http://localhost:8080/users

# POST - Créer un utilisateur
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'

# GET - Récupérer un utilisateur
curl http://localhost:8080/users/1

# PUT - Modifier un utilisateur
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated","email":"alice@new.com"}'

# DELETE - Supprimer un utilisateur
curl -X DELETE http://localhost:8080/users/1
```

---

## 5. Middleware

### 5.1. Qu'est-ce qu'un middleware ?

Un **middleware** est une fonction qui enveloppe un handler pour ajouter des fonctionnalités transversales :

- Logging
- Authentification
- CORS
- Rate limiting
- Récupération après panic

### 5.2. Structure d'un middleware

```go
type Middleware func(http.HandlerFunc) http.HandlerFunc

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Début: %s %s", r.Method, r.URL.Path)

        next(w, r) // Appelle le handler original

        log.Printf("Fin: %s %s (%v)", r.Method, r.URL.Path, time.Since(start))
    }
}
```

### 5.3. Middleware de logging complet

```go
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Capturer le status code
        rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

        log.Printf("→ %s %s", r.Method, r.URL.Path)

        next(rw, r)

        duration := time.Since(start)
        log.Printf("← %s %s - %d (%v)",
            r.Method, r.URL.Path, rw.statusCode, duration)
    }
}

// Wrapper pour capturer le status code
type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}
```

### 5.4. Middleware d'authentification

```go
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Vérifier le header Authorization
        token := r.Header.Get("Authorization")
        if token == "" {
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(map[string]string{
                "error": "Token manquant",
            })
            return
        }

        // Vérifier le token (simplifié)
        if token != "Bearer secret-token-123" {
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(map[string]string{
                "error": "Token invalide",
            })
            return
        }

        // Ajouter l'utilisateur au contexte
        ctx := context.WithValue(r.Context(), "user", "admin")
        next(w, r.WithContext(ctx))
    }
}
```

### 5.5. Middleware CORS

```go
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        next(w, r)
    }
}
```

### 5.6. Middleware de récupération (recovery)

```go
func recoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("PANIC: %v", err)
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(map[string]string{
                    "error": "Erreur interne du serveur",
                })
            }
        }()
        next(w, r)
    }
}
```

### 5.7. Chaînage de middleware

```go
func chain(middlewares ...Middleware) Middleware {
    return func(next http.HandlerFunc) http.HandlerFunc {
        for i := len(middlewares) - 1; i >= 0; i-- {
            next = middlewares[i](next)
        }
        return next
    }
}

func main() {
    // Appliquer plusieurs middleware
    handler := chain(
        recoveryMiddleware,
        loggingMiddleware,
        corsMiddleware,
        authMiddleware,
    )(protectedHandler)

    http.HandleFunc("/protected", handler)
    http.ListenAndServe(":8080", nil)
}
```

---

## 6. Servir des fichiers statiques

### 6.1. Serveur de fichiers simple

```go
func main() {
    // Servir tout le dossier "static"
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Fichier spécifique
    http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/favicon.ico")
    })

    http.ListenAndServe(":8080", nil)
}
```

### 6.2. Single Page Application (SPA)

```go
func spaHandler(w http.ResponseWriter, r *http.Request) {
    // Si le chemin demande un fichier qui existe, le servir
    path := "./static" + r.URL.Path
    if _, err := os.Stat(path); err == nil {
        http.ServeFile(w, r, path)
        return
    }

    // Sinon, servir index.html (pour le routage frontend)
    http.ServeFile(w, r, "./static/index.html")
}

func main() {
    http.HandleFunc("/", spaHandler)
    http.ListenAndServe(":8080", nil)
}
```

---

## 7. Serveur HTTP avancé

### 7.1. Configuration personnalisée

```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", helloHandler)

    server := &http.Server{
        Addr:         ":8080",
        Handler:      mux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
        MaxHeaderBytes: 1 << 20, // 1 MB
    }

    log.Println("Serveur démarré sur :8080")
    log.Fatal(server.ListenAndServe())
}
```

### 7.2. Arrêt gracieux (graceful shutdown)

```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", helloHandler)

    server := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }

    // Goroutine pour le serveur
    go func() {
        log.Println("Serveur démarré sur :8080")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Erreur serveur: %v", err)
        }
    }()

    // Attendre un signal d'arrêt
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
    <-stop

    log.Println("Arrêt du serveur...")

    // Contexte avec timeout pour l'arrêt
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Erreur lors de l'arrêt: %v", err)
    }

    log.Println("Serveur arrêté proprement")
}
```

### 7.3. HTTPS / TLS

```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", helloHandler)

    // Générer des certificats auto-signés (pour le développement)
    // go run /usr/local/go/src/crypto/tls/generate_cert.go --host=localhost

    log.Println("Serveur HTTPS démarré sur :8443")
    log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", mux))
}
```

---

## 8. Client HTTP

### 8.1. Requêtes simples

```go
func main() {
    // GET simple
    resp, err := http.Get("https://api.github.com/users/golang")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))

    // GET avec paramètres
    resp, err = http.Get("https://api.example.com/search?q=golang&limit=10")
}

// POST avec JSON
func postJSON() {
    user := map[string]interface{}{
        "name":  "Alice",
        "email": "alice@example.com",
    }

    jsonData, _ := json.Marshal(user)

    resp, err := http.Post(
        "https://api.example.com/users",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
}
```

### 8.2. Client configurable

```go
func createHTTPClient() *http.Client {
    return &http.Client{
        Timeout: 10 * time.Second,
        Transport: &http.Transport{
            MaxIdleConns:        100,
            MaxIdleConnsPerHost: 10,
            IdleConnTimeout:     90 * time.Second,
        },
    }
}

func makeRequest(client *http.Client, url string) (*http.Response, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("User-Agent", "MonApp/1.0")
    req.Header.Set("Authorization", "Bearer token123")

    return client.Do(req)
}
```

### 8.3. Requête avec contexte et timeout

```go
func fetchWithTimeout(url string, timeout time.Duration) ([]byte, error) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return io.ReadAll(resp.Body)
}
```

---

## TP final – Module 13 – API de gestion de tâches (Todo API)

### Énoncé

Créez une API REST complète pour gérer une liste de tâches (Todo List).

**Fonctionnalités :**

1. **CRUD complet** sur les tâches
2. **Validation des données**
3. **Middleware** : logging + authentification basique
4. **Filtrage** : récupérer les tâches par statut (complétées/non complétées)
5. **Persistance** : sauvegarde dans un fichier JSON (optionnel)

**Structure d'une tâche :**
```go
type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
}
```

**Endpoints :**

| Méthode | Endpoint | Description |
|---------|----------|-------------|
| GET | `/todos` | Récupérer toutes les tâches |
| GET | `/todos?completed=true` | Filtrer par statut |
| GET | `/todos/{id}` | Récupérer une tâche |
| POST | `/todos` | Créer une tâche |
| PUT | `/todos/{id}` | Modifier une tâche |
| DELETE | `/todos/{id}` | Supprimer une tâche |
| PATCH | `/todos/{id}/toggle` | Basculer le statut |

### Solution

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
    "sync"
    "time"
)

// Modèle
type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
}

// Stockage
var (
    todos   = make(map[int]Todo)
    nextID  = 1
    todoMu  sync.RWMutex
)

// Middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("→ %s %s", r.Method, r.URL.Path)
        next(w, r)
        log.Printf("← %s %s (%v)", r.Method, r.URL.Path, time.Since(start))
    }
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        apiKey := r.Header.Get("X-API-Key")
        if apiKey != "secret-todo-key-2024" {
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(map[string]string{
                "error": "Clé API invalide",
            })
            return
        }
        next(w, r)
    }
}

func recoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("PANIC: %v", err)
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(map[string]string{
                    "error": "Erreur interne",
                })
            }
        }()
        next(w, r)
    }
}

func chain(middlewares ...Middleware) Middleware {
    return func(next http.HandlerFunc) http.HandlerFunc {
        for i := len(middlewares) - 1; i >= 0; i-- {
            next = middlewares[i](next)
        }
        return next
    }
}

// Handlers
func todosHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getAllTodos(w, r)
    case http.MethodPost:
        createTodo(w, r)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
    // Extraire l'ID
    parts := strings.Split(r.URL.Path, "/")
    if len(parts) != 3 {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(parts[2])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "ID invalide"})
        return
    }

    switch r.Method {
    case http.MethodGet:
        getTodoByID(w, r, id)
    case http.MethodPut:
        updateTodo(w, r, id)
    case http.MethodDelete:
        deleteTodo(w, r, id)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func toggleHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPatch {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    parts := strings.Split(r.URL.Path, "/")
    if len(parts) != 4 || parts[3] != "toggle" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(parts[2])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "ID invalide"})
        return
    }

    todoMu.Lock()
    defer todoMu.Unlock()

    todo, exists := todos[id]
    if !exists {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Tâche non trouvée"})
        return
    }

    todo.Completed = !todo.Completed
    todos[id] = todo

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
    todoMu.RLock()
    defer todoMu.RUnlock()

    // Filtrage par statut
    filterCompleted := r.URL.Query().Get("completed")
    var filtered []Todo

    for _, todo := range todos {
        if filterCompleted == "" {
            filtered = append(filtered, todo)
        } else if filterCompleted == "true" && todo.Completed {
            filtered = append(filtered, todo)
        } else if filterCompleted == "false" && !todo.Completed {
            filtered = append(filtered, todo)
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(filtered)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "JSON invalide"})
        return
    }

    // Validation
    if todo.Title == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Le titre est requis"})
        return
    }

    todoMu.Lock()
    defer todoMu.Unlock()

    todo.ID = nextID
    todo.CreatedAt = time.Now()
    todo.Completed = false
    nextID++
    todos[todo.ID] = todo

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(todo)
}

func getTodoByID(w http.ResponseWriter, r *http.Request, id int) {
    todoMu.RLock()
    defer todoMu.RUnlock()

    todo, exists := todos[id]
    if !exists {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Tâche non trouvée"})
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request, id int) {
    var updated Todo
    if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "JSON invalide"})
        return
    }

    todoMu.Lock()
    defer todoMu.Unlock()

    existing, exists := todos[id]
    if !exists {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Tâche non trouvée"})
        return
    }

    // Mise à jour partielle
    if updated.Title != "" {
        existing.Title = updated.Title
    }
    existing.Completed = updated.Completed
    todos[id] = existing

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(existing)
}

func deleteTodo(w http.ResponseWriter, r *http.Request, id int) {
    todoMu.Lock()
    defer todoMu.Unlock()

    if _, exists := todos[id]; !exists {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Tâche non trouvée"})
        return
    }

    delete(todos, id)
    w.WriteHeader(http.StatusNoContent)
}

func main() {
    // Ajouter quelques tâches de démonstration
    todos[1] = Todo{ID: 1, Title: "Apprendre Go", Completed: false, CreatedAt: time.Now()}
    todos[2] = Todo{ID: 2, Title: "Créer une API REST", Completed: false, CreatedAt: time.Now()}
    nextID = 3

    mux := http.NewServeMux()
    mux.HandleFunc("/todos", chain(
        recoveryMiddleware,
        loggingMiddleware,
        authMiddleware,
    )(todosHandler))
    mux.HandleFunc("/todos/", chain(
        recoveryMiddleware,
        loggingMiddleware,
        authMiddleware,
    )(func(w http.ResponseWriter, r *http.Request) {
        if strings.HasSuffix(r.URL.Path, "/toggle") {
            toggleHandler(w, r)
        } else {
            todoHandler(w, r)
        }
    }))

    server := &http.Server{
        Addr:         ":8080",
        Handler:      mux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    log.Println("📋 Todo API démarrée sur http://localhost:8080")
    log.Println("🔑 Clé API requise: X-API-Key: secret-todo-key-2024")
    log.Fatal(server.ListenAndServe())
}
```

### Tests avec curl

```bash
# Créer une tâche
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -H "X-API-Key: secret-todo-key-2024" \
  -d '{"title":"Faire les courses"}'

# Récupérer toutes les tâches
curl -H "X-API-Key: secret-todo-key-2024" http://localhost:8080/todos

# Filtrer les tâches complétées
curl -H "X-API-Key: secret-todo-key-2024" "http://localhost:8080/todos?completed=false"

# Basculer le statut
curl -X PATCH -H "X-API-Key: secret-todo-key-2024" http://localhost:8080/todos/1/toggle

# Supprimer une tâche
curl -X DELETE -H "X-API-Key: secret-todo-key-2024" http://localhost:8080/todos/1
```

---

## Récapitulatif des acquis

| Concept | Syntaxe | Utilisation |
|---------|---------|-------------|
| Serveur basique | `http.HandleFunc()` + `http.ListenAndServe()` | Points d'entrée |
| Routes | `http.NewServeMux()` | Organisation des routes |
| JSON | `json.Marshal()` / `json.Unmarshal()` | API REST |
| Middleware | Fonction qui enveloppe un handler | Logging, auth, CORS |
| Client HTTP | `http.Get()`, `http.Client` | Appels externes |
| Fichiers statiques | `http.FileServer()` | CSS, JS, images |
| Contexte | `r.Context()` | Annulation, timeout |

---

## Prochain module (aperçu)

Dans le Module 14, nous découvrirons :
- Les tests unitaires avec `testing`
- Les tests table-driven
- Les benchmarks
- La couverture de code

---

**Félicitations !** Vous savez maintenant créer des serveurs HTTP et des API REST complètes avec Go. Vous maîtrisez les middleware, la manipulation de JSON, et les bonnes pratiques du développement web. 🚀
