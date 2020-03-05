package main

import (
	"sync"
	"testing"
)

func TestGetPlayerScore(t *testing.T) {
	repo := &InMemoryPlayerStore{
		map[string]int{
			"1": 20,
			"2": 10,
		},
		sync.RWMutex{},
	}
	player1Score, _ := repo.GetPlayerScore("1")
	assertPlayerScore(t, &PlayerScore{"1", 20}, player1Score)

	player2Score, _ := repo.GetPlayerScore("2")
	assertPlayerScore(t, &PlayerScore{"2", 10}, player2Score)

	player3Score, _ := repo.GetPlayerScore("3")
	assertNilPlayerScore(t, player3Score)
}

func assertPlayerScore(t *testing.T, want, got *PlayerScore) {
	t.Helper()
	if *got != *want {
		t.Errorf("\nincorrect player score\nwant: %q, got %q", want, got)
	}
}

func assertNilPlayerScore(t *testing.T, got *PlayerScore) {
	t.Helper()
	if got != nil {
		t.Errorf("\nunexpected player score\nwant: nil, got: %q", got)
	}
}
