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

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("\nwant: %.2f, got: %.2f", want, got)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{5.0, 8.0}
		want := 40.0
		checkArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})

}

func TestAreaWithTables(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("\nwant: %g, got: %g", tt.want, got)
		}
	}
}
