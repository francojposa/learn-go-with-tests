package main

import (
	"bytes"
	"testing"
)

func TestCountown(t *testing.T) {
	buffer := &bytes.Buffer{}
	Countdown(buffer)

	want := "3\n2\n1\nGo!"
	got := buffer.String()

	if got != want {
		t.Errorf("\nwant: %q, got: %q", want, got)
	}

}
