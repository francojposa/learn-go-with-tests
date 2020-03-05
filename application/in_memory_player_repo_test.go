package main

import (
	"math/rand"
	"reflect"
	"sync"
	"testing"
)

func TestGetPlayerScore(t *testing.T) {
	repo := &InMemoryPlayerStore{
		map[string]PlayerScore{
			"1": PlayerScore{"1", 20},
			"2": PlayerScore{"2", 10},
		},
		sync.RWMutex{},
	}
	player1Score, _ := repo.GetPlayerScore("1")
	assertPlayerScore(t, PlayerScore{"1", 20}, player1Score)

	player2Score, _ := repo.GetPlayerScore("2")
	assertPlayerScore(t, PlayerScore{"2", 10}, player2Score)

	player3Score, _ := repo.GetPlayerScore("3")
	assertEmptyPlayerScore(t, player3Score)
}

func TestListPlayerScores(t *testing.T) {

	t.Run("list player scores for non-empty repo", func(t *testing.T) {
		repo := &InMemoryPlayerStore{
			map[string]PlayerScore{
				"1": PlayerScore{"1", 20},
				"2": PlayerScore{"2", 10},
			},
			sync.RWMutex{},
		}
		expectedScores := []PlayerScore{
			PlayerScore{"1", 20},
			PlayerScore{"2", 10},
		}
		playerScores := repo.ListPlayerScores()
		assertPlayerScoreSlice(t, expectedScores, playerScores)
	})

	t.Run("list player scores for empty repo", func(t *testing.T) {
		repo := NewInMemoryPlayerRepo()
		expectedScores := []PlayerScore{}
		playerScores := repo.ListPlayerScores()
		assertPlayerScoreSlice(t, expectedScores, playerScores)
	})

}

func TestRecordPlayerScore(t *testing.T) {
	t.Run("create first player score for nonexistent player", func(t *testing.T) {
		repo := NewInMemoryPlayerRepo()
		expectedPlayer1Score := PlayerScore{"1", 1}
		player1Score := repo.RecordPlayerScore("1")
		assertPlayerScore(t, expectedPlayer1Score, player1Score)
	})

	t.Run("increment player score for existing player", func(t *testing.T) {
		repo := &InMemoryPlayerStore{
			map[string]PlayerScore{
				"1": PlayerScore{"1", 20},
				"2": PlayerScore{"2", 10},
			},
			sync.RWMutex{},
		}

		expectedPlayer1Score := PlayerScore{"1", 22}
		repo.RecordPlayerScore("1")
		player1Score := repo.RecordPlayerScore("1")

		assertPlayerScore(t, expectedPlayer1Score, player1Score)

		expectedPlayer2Score := PlayerScore{"2", 11}
		player2Score := repo.RecordPlayerScore("2")

		assertPlayerScore(t, expectedPlayer2Score, player2Score)
	})
}

func BenchmarkRecordPlayerScore(b *testing.B) {
	repo := NewInMemoryPlayerRepo()
	for i := 0; i < b.N; i++ {
		id := string(rand.Intn(50))
		repo.RecordPlayerScore(id)
	}
}

func assertPlayerScore(t *testing.T, want, got PlayerScore) {
	t.Helper()
	if got != want {
		t.Errorf("\nincorrect player score\nwant: %q, got: %q", want, got)
	}
}

func assertEmptyPlayerScore(t *testing.T, got PlayerScore) {
	t.Helper()
	if !got.IsEmpty() {
		t.Errorf(
			"\nunexpected non-empty player score\nwant: %q, got: %q",
			PlayerScore{},
			got,
		)
	}
}

func assertPlayerScoreSlice(t *testing.T, want, got []PlayerScore) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf(
			"\nincorrect slice of player scores\nwant: %q, got: %q",
			want,
			got,
		)
	}
}
