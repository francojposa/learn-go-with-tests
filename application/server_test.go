package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(id string) int {
	score := s.scores[id]
	return score

}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"1": 20,
			"2": 10,
		},
	}

	server := &PlayerServer{&store}

	t.Run("returns player 1's score", func(t *testing.T) {
		request := newGetScoreRequest("1")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		want := "20"
		got := response.Body.String()

		assertResponseBody(t, want, got)
	})

	t.Run("return's player 2's score", func(t *testing.T) {
		request := newGetScoreRequest("2")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		want := "10"
		got := response.Body.String()
		assertResponseBody(t, want, got)

	})
}

func newGetScoreRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", id), nil)
	return req
}

func assertResponseBody(t *testing.T, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("\nwant: %q, got %q", want, got)
	}
}
