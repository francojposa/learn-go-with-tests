package main

import (
	"reflect"
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

func TestListPlayerScores(t *testing.T) {

	t.Run("list player scores for non-empty repo", func(t *testing.T) {
		repo := &InMemoryPlayerStore{
			map[string]int{
				"1": 20,
				"2": 10,
			},
			sync.RWMutex{},
		}
		expectedScores := []*PlayerScore{
			&PlayerScore{"1", 20},
			&PlayerScore{"2", 10},
		}
		playerScores := repo.ListPlayerScores()
		assertPlayerScoreSlice(t, expectedScores, playerScores)
	})

}

func assertPlayerScore(t *testing.T, want, got *PlayerScore) {
	t.Helper()
	if *got != *want {
		t.Errorf("\nincorrect player score\nwant: %q, got: %q", want, got)
	}
}

func assertNilPlayerScore(t *testing.T, got *PlayerScore) {
	t.Helper()
	if got != nil {
		t.Errorf("\nunexpected player score\nwant: nil, got: %q", got)
	}
}

func assertPlayerScoreSlice(t *testing.T, want, got []*PlayerScore) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf(
			"\nincorrect slice of player scores\nwant: %q, got: %q",
			want,
			got,
		)
	}
}
