package players

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type PlayerScore struct {
	PlayerID string
	Score    int
}

func (ps PlayerScore) String() string {
	return fmt.Sprintf("PlayerID: %s, Score: %d", ps.PlayerID, ps.Score)
}

func (ps PlayerScore) IsEmpty() bool {
	return ps.PlayerID == ""
}

type PlayerRepo interface {
	GetPlayerScore(id string) (PlayerScore, bool)
	ListPlayerScores() []PlayerScore
	RecordPlayerScore(id string) PlayerScore
}

func NewPlayerHandler(repo PlayerRepo) *PlayerHandler {
	return &PlayerHandler{repo}
}

type PlayerHandler struct {
	repo PlayerRepo
}

func (p *PlayerHandler) GetPlayerScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	playerScore, ok := p.repo.GetPlayerScore(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(playerScore)
	}
}

func (p *PlayerHandler) ListPlayerScores(w http.ResponseWriter, r *http.Request) {
	playerScores := p.repo.ListPlayerScores()
	json.NewEncoder(w).Encode(playerScores)
}

func (p *PlayerHandler) RecordPlayerScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	playerScore := p.repo.RecordPlayerScore(id)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(playerScore)
}
