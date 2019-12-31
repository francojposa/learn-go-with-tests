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

func (s *StubPlayerStore) GetPlayerScore(id string) (int, bool) {
	score, ok := s.scores[id]
	return score, ok
}

func (s *StubPlayerStore) RecordPlayerScore(id string) int {
	s.scores[id]++
	return s.scores[id]
}

func TestGETPlayerScore(t *testing.T) {
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

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("return's player 2's score", func(t *testing.T) {
		request := newGetScoreRequest("2")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("return 404 for player id not found", func(t *testing.T) {
		request := newGetScoreRequest("3")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertResponseStatus(t, http.StatusNotFound, response.Code)
	})
}

func TestPOSTPlayerScore(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
	}
	server := &PlayerServer{&store}

	t.Run("returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/3", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertResponseStatus(t, http.StatusCreated, response.Code)
	})
}

func newGetScoreRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", id), nil)
	return req
}

func assertResponseStatus(t *testing.T, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("\nincorrect http status\nwant: %d, got %d", want, got)
	}
}

func assertResponseBody(t *testing.T, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("\nincorrect response body\nwant: %q, got %q", want, got)
	}
}
