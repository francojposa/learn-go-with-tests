package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning URL of fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("\ndid not expect an error but received one: %v", err)
		}

		if got != want {
			t.Errorf("\nwant: %q, got: %q", want, got)
		}
	})

	t.Run("returns an error if a server doesn't response within timeout", func(t *testing.T) {
		server := makeDelayedServer(10 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 5*time.Millisecond)

		if err == nil {
			t.Error("\nexpected an error, but didn't receive one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}
	return httptest.NewServer(http.HandlerFunc(handler))
}
