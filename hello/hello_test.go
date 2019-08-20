package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Franco")
	want := "Hello, Franco"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}