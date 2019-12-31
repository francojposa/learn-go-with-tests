package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(id string) (int, bool)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/player/")
	score, ok := p.store.GetPlayerScore(id)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func GetPlayerScore(id string) (string, bool) {
	if id == "1" {
		return "20", true
	}
	if id == "2" {
		return "10", true
	}
	return "", false
}
