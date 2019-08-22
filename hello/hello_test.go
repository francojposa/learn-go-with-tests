package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Franco", "")
		want := "Hello, Franco"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Franco", "Italian")
		want := "Ciao, Franco"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(T *testing.T) {
		got := Hello("Franco", "French")
		want := "Bonjour, Franco"
		assertCorrectMessage(t, got, want)
	})

}
