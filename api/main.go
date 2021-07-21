package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/Dragonqos/go_boilerplate/api/core/middleware"
	"github.com/Dragonqos/go_boilerplate/api/core/db"
	"github.com/Dragonqos/go_boilerplate/api/repository"
	"github.com/Dragonqos/go_boilerplate/api/routes"
	"log"
	"net/http"
)

func init() {
	// Laod Dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load .env: %v", err)
	}

	// Load Database
	dbs, ctx, err := db.GetDatabase()
	if err != nil {
		log.Fatalf("Database configuration failed: %v", err)
	}

	_ = repository.CreateChannelRepository(dbs, ctx)
}

func main() {

	// run MQ handlers

	// define routers
	handler := mux.NewRouter()
	handler.HandleFunc("/", routes.Index).Methods("GET")

	// Registry
	handler.HandleFunc("/channels", routes.GetCollectionChannels).Methods("GET")
	handler.HandleFunc("/channels", routes.PostChannel).Methods("POST")
	handler.HandleFunc("/channels/{id}", routes.GetChannel).Methods("GET")
	//handler.HandleFunc("/channels/{id}", routes.PutChannel).Methods("PUT")
	//handler.HandleFunc("/channels/{id}", routes.DeleteChannel).Methods("DELETE")

	origins := handlers.AllowedOrigins([]string{"http://stage1-local.com", "https://stage2.test.com", "https://rc.test.com"})

	handler.Use(middleware.AuthMiddleware)

	// Listen requests
	http.ListenAndServe(":8080", handlers.CORS(origins)(handler))
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
