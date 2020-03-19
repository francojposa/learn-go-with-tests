package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/francojposa/learn-go-with-tests/app"
)

func main() {
	playerRepo := players.NewInMemoryPlayerRepo()
	playerHandler := players.NewPlayerHandler(playerRepo)

	router := mux.NewRouter()

	router.HandleFunc("/players/{id}", playerHandler.GetPlayerScore).Methods("GET")
	router.HandleFunc("/players/{id}", playerHandler.RecordPlayerScore).Methods("POST")
	router.HandleFunc("/players/", playerHandler.ListPlayerScores).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:5000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("running http server on port 5000")
	log.Fatal(srv.ListenAndServe())
}
