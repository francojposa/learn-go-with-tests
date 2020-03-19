package players

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type StubPlayerRepo struct {
	scores map[string]int
}

func (r *StubPlayerRepo) GetPlayerScore(id string) (PlayerScore, bool) {
	score, ok := r.scores[id]
	return PlayerScore{id, score}, ok
}

func (r *StubPlayerRepo) ListPlayerScores() []PlayerScore {
	playerScores := []PlayerScore{}
	for k, v := range r.scores {
		playerScores = append(playerScores, PlayerScore{k, v})
	}
	return playerScores
}

func (r *StubPlayerRepo) RecordPlayerScore(id string) PlayerScore {
	r.scores[id]++
	return PlayerScore{id, r.scores[id]}
}

func SetupTestPlayerHandler(repo PlayerRepo) *mux.Router {
	playerHandler := &PlayerHandler{repo}

	router := mux.NewRouter()

	router.HandleFunc("/players/{id}", playerHandler.GetPlayerScore).Methods("GET")
	router.HandleFunc("/players/{id}", playerHandler.RecordPlayerScore).Methods("POST")
	router.HandleFunc("/players/", playerHandler.ListPlayerScores).Methods("GET")

	return router
}

func TestPOSTAndGETPlayerScore(t *testing.T) {
	playerHandler := SetupTestPlayerHandler(NewInMemoryPlayerRepo())
	playerID := "1"

	playerHandler.ServeHTTP(httptest.NewRecorder(), newPostPlayerScoreRequest(playerID))
	playerHandler.ServeHTTP(httptest.NewRecorder(), newPostPlayerScoreRequest(playerID))
	playerHandler.ServeHTTP(httptest.NewRecorder(), newPostPlayerScoreRequest(playerID))

	response := httptest.NewRecorder()
	playerHandler.ServeHTTP(response, newGetPlayerScoreRequest(playerID))

	assertResponseStatus(t, response.Code, http.StatusOK)
	assertResponsePlayerScore(t, response, PlayerScore{"1", 3})
}

func TestGETPlayerScore(t *testing.T) {
	repo := StubPlayerRepo{
		map[string]int{
			"1": 20,
			"2": 10,
		},
	}

	playerHandler := SetupTestPlayerHandler(&repo)

	t.Run("returns player 1's score", func(t *testing.T) {
		request := newGetPlayerScoreRequest("1")
		response := httptest.NewRecorder()
		playerHandler.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponsePlayerScore(t, response, PlayerScore{"1", 20})
	})

	t.Run("return's player 2's score", func(t *testing.T) {
		request := newGetPlayerScoreRequest("2")
		response := httptest.NewRecorder()
		playerHandler.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponsePlayerScore(t, response, PlayerScore{"2", 10})
	})

	t.Run("return 404 for player id not found", func(t *testing.T) {
		request := newGetPlayerScoreRequest("3")
		response := httptest.NewRecorder()
		playerHandler.ServeHTTP(response, request)
		assertResponseStatus(t, http.StatusNotFound, response.Code)
	})
}

func TestListPlayerScore(t *testing.T) {
	repo := StubPlayerRepo{
		map[string]int{
			"1": 20,
			"2": 10,
		},
	}

	playerHandler := SetupTestPlayerHandler(&repo)
	request := newListPlayerScoreRequest()
	response := httptest.NewRecorder()
	playerHandler.ServeHTTP(response, request)

	assertResponseStatus(t, response.Code, http.StatusOK)

	expectedPlayerScores := []PlayerScore{
		{"1", 20},
		{"2", 10},
	}

	assertResponsePlayerScores(t, response, expectedPlayerScores)
}

func TestPOSTPlayerScore(t *testing.T) {
	repo := StubPlayerRepo{
		map[string]int{},
	}
	playerHandler := SetupTestPlayerHandler(&repo)

	t.Run("returns accepted on POST", func(t *testing.T) {
		request := newPostPlayerScoreRequest("1")
		response := httptest.NewRecorder()
		playerHandler.ServeHTTP(response, request)
		assertResponseStatus(t, http.StatusCreated, response.Code)
	})
}

func newGetPlayerScoreRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", id), nil)
	return req
}

func newListPlayerScoreRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/", nil)
	return req
}

func newPostPlayerScoreRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", id), nil)
	return req
}

func assertResponseStatus(t *testing.T, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("\nincorrect http status\nwant: %d, got %d", want, got)
	}
}

func assertResponsePlayerScore(
	t *testing.T,
	response *httptest.ResponseRecorder,
	want PlayerScore,
) {
	t.Helper()
	gotPlayerScore := PlayerScore{}
	err := json.NewDecoder(response.Body).Decode(&gotPlayerScore)

	if err != nil {
		t.Errorf(
			"Unable to parse response from server %q into slice of PlayerScore, '%v'",
			response.Body,
			err,
		)
	}

	assertPlayerScore(t, want, gotPlayerScore)

}

func assertResponsePlayerScores(
	t *testing.T,
	response *httptest.ResponseRecorder,
	want []PlayerScore,
) {
	t.Helper()

	gotPlayerScores := []PlayerScore{}
	err := json.NewDecoder(response.Body).Decode(&gotPlayerScores)

	if err != nil {
		t.Errorf(
			"Unable to parse response from server %q into slice of PlayerScore, '%v'",
			response.Body,
			err,
		)
	}

	assertPlayerScoreSlice(t, want, gotPlayerScores)
}
