package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Franco")

	want := "Hello, Franco"
	got := buffer.String()

	if got != want {
		t.Errorf("\nwant: %q, got: %q", want, got)
	}
}
