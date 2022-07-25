package main

import (
	"errors"
	"github.com/anugrahsinghal/http_server_golang/internal/database"
	"net/http"
	"time"
)

const address = "localhost:8080"

func main() {
	serverMultiplexer := http.NewServeMux()

	serverMultiplexer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, 200, database.User{
			Email: "test@example.com",
		})
	})
	serverMultiplexer.HandleFunc("/err", func(writer http.ResponseWriter, request *http.Request) {
		respondWithError(writer, 400, errors.New("error handler"))
	})

	config := apiConfig{dbClient: database.NewClient("db.json")}
	config.dbClient.EnsureDB()

	serverMultiplexer.HandleFunc("/users", config.endpointUsersHandler)
	serverMultiplexer.HandleFunc("/users/", config.endpointUsersHandler)

	serverMultiplexer.HandleFunc("/posts", config.endpointPostHandler)
	serverMultiplexer.HandleFunc("/posts/", config.endpointPostHandler)

	server := http.Server{
		Handler:      serverMultiplexer,
		Addr:         address,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	server.ListenAndServe()
}
