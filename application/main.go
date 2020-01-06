package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	fmt.Println("running http server on port 5000")
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
