package main

import (
	"fmt"
	"log"
	"net/http"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
	}
}

type InMemoryPlayerStore struct {
	scores map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(id string) (int, bool) {
	score, ok := i.scores[id]
	return score, ok
}

func (i *InMemoryPlayerStore) RecordPlayerScore(id string) int {
	i.scores[id]++
	return i.scores[id]
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	fmt.Println("starting http server on port 5000")
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
