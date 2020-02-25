package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type PlayerStore interface {
	GetPlayerScore(id string) (int, bool)
	RecordPlayerScore(id string) int
}

type PlayerHandler struct {
	store PlayerStore
}

func (p *PlayerHandler) getPlayerScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	score, ok := p.store.GetPlayerScore(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Fprint(w, score)
	}

}

func (p *PlayerHandler) recordPlayerScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	score := p.store.RecordPlayerScore(id)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, score)
}
