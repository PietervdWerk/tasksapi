package main

import (
	"crypto/rand"
	"crypto/rsa"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pietervdwerk/tasksapi/internal/api"
	"github.com/pietervdwerk/tasksapi/internal/clients"
	"github.com/pietervdwerk/tasksapi/internal/tasks"
)

func main() {
	// Initialize the database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Migrate the database
	err = migrate(db)
	if err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	// Initialize the router
	r := chi.NewRouter()

	// Default middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Initialize the API
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("error creating private key: %v", err)
	}
	tasksRepo := tasks.New(db)
	clientsRepo := clients.New(db)
	app, err := api.NewAPI(tasksRepo, clientsRepo, api.WithPrivateKey(privateKey))
	if err != nil {
		log.Fatalf("error creating API: %v", err)
	}

	handlers := api.ServerInterfaceWrapper{
		Handler: api.NewStrictHandler(app, nil),
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
	}

	// Create a JWT token middleware
	tokenMiddleware := jwtauth.New("RS256", nil, privateKey.PublicKey)

	r.Post("/token", handlers.PostToken)
	r.Route("/tasks", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenMiddleware))
		r.Use(jwtauth.Authenticator(tokenMiddleware))

		r.Get("/", handlers.GetTasks)
		r.Post("/", handlers.PostTasks)
		r.Delete("/{taskId}", handlers.DeleteTasksTaskId)
		r.Get("/{taskId}", handlers.GetTasksTaskId)
		r.Put("/{taskId}", handlers.PutTasksTaskId)

	})

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(clients.SchemaSQL)
	if err != nil {
		return fmt.Errorf("error creating clients table: %w", err)
	}

	_, err = db.Exec(tasks.SchemaSQL)
	if err != nil {
		return fmt.Errorf("error creating tasks table: %w", err)
	}

	_, err = db.Exec(`
INSERT INTO clients (id, secret) 
VALUES ('b218e1c1-1eb7-484c-89c2-a5d6ce731553', 'dummy-secret');`)
	if err != nil {
		return fmt.Errorf("error seeding database: %w", err)
	}

	return nil
}
