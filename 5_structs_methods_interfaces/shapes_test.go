package oop

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	want := 40.0
	got := Perimeter(rectangle)

	if got != want {
		t.Errorf("\nwant: %.2f, got: %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{5.0, 8.0}
	want := 40.0
	got := Area(rectangle)

	if got != want {
		t.Errorf("\nwant: %.2f, got: %.2f", want, got)
	}
}
