package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(id string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/player/")
	score := p.store.GetPlayerScore(id)
	fmt.Fprint(w, score)
}

func GetPlayerScore(id string) string {
	if id == "1" {
		return "20"
	}
	if id == "2" {
		return "10"
	}
	return ""
}
