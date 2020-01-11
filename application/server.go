package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(id string) (int, bool)
	RecordPlayerScore(id string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.recordPlayerScore(w, id)
	case http.MethodGet:
		p.getPlayerScore(w, id)
	}
}

func (p *PlayerServer) getPlayerScore(w http.ResponseWriter, id string) {
	score, ok := p.store.GetPlayerScore(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Fprint(w, score)
	}

}

func (p *PlayerServer) recordPlayerScore(w http.ResponseWriter, id string) {
	score := p.store.RecordPlayerScore(id)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, score)
}
