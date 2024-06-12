
package httpserver

import (
    "context"
    "encoding/json"
    "user-management-service/database"
    "user-management-service/models"
    "net/http"
    "log"
    "user-management-service/config"
    "github.com/gorilla/mux"
    "os"
    "os/signal"
    "time"
)

type Server struct {
    httpServer *http.Server
}

func NewServer(cfg *config.Config) *Server {
    router := mux.NewRouter()

    // Define routes
    router.HandleFunc("/users", createUserHandler).Methods("POST")
    router.HandleFunc("/users/login", loginUserHandler).Methods("POST")
    router.HandleFunc("/users/{userId}", getUserHandler).Methods("GET")

    httpServer := &http.Server{
        Addr:    ":" + cfg.HTTPPort,
        Handler: router,
    }

    return &Server{httpServer: httpServer}
}

func (s *Server) Start() error {
    // Start the server in a new goroutine
    go func() {
        if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("could not listen on %s: %v
", s.httpServer.Addr, err)
        }
    }()
    log.Printf("Server is ready to handle requests at %s", s.httpServer.Addr)

    // Wait for interrupt signal to gracefully shutdown the server
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)

    <-stop

    log.Println("Server is shutting down...")

    ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
    defer cancel()

    return s.httpServer.Shutdown(ctx)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := database.CreateUser(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func loginUserHandler(w http.ResponseWriter, r *http.Request) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    token, err := database.AuthenticateUser(credentials.Email, credentials.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["userId"]

    user, err := database.GetUserByID(userId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
