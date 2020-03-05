package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type PlayerScore struct {
	PlayerId string
	Score    int
}

func (ps *PlayerScore) String() string {
	return fmt.Sprintf("PlayerId: %s, Score: %d", ps.PlayerId, ps.Score)
}

type PlayerRepo interface {
	GetPlayerScore(id string) (*PlayerScore, bool)
	//ListPlayerScores() []PlayerScore
	RecordPlayerScore(id string) *PlayerScore
}

type PlayerHandler struct {
	store PlayerRepo
}

func (p *PlayerHandler) getPlayerScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	playerScore, ok := p.store.GetPlayerScore(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Fprint(w, playerScore.Score)
	}

}

func (p *PlayerHandler) recordPlayerScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	playerScore := p.store.RecordPlayerScore(id)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, playerScore.Score)
}
