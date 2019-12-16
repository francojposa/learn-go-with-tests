package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	wantOutput := "3\n2\n1\nGo!\n"
	gotOutput := buffer.String()

	if gotOutput != wantOutput {
		t.Errorf("\nwant: %q, got: %q", wantOutput, gotOutput)
	}

	wantCalls := 4

	if spySleeper.Calls != wantCalls {
		t.Errorf("\nNot enough calls to Sleep, want: %d, got: %d", wantCalls, spySleeper.Calls)
	}

}
