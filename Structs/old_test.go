package main

import (
	"testing"

)

func TestSuite(t *testing.T)  {
	t.Run("rectangle perimeter testcase", func(t *testing.T) {
		rectangle := Rectangle{Width: 10.0, Height: 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		assertion(t, got, want)

	})

	t.Run("rectangle area testcase", func(t *testing.T) {
		rectangle := Rectangle{Width: 12.0, Height: 6.0}
		got := rectangle.Area()
		want := 72.0

		assertion(t, got, want)
	})

	t.Run("circle perimeter", func(t *testing.T) {
		circle := Circle{Radius: 3}
		got := circle.Perimeter()
		want := 18.84955592153876

		assertion(t, got, want)
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{Radius: 3}
		got := circle.Area()
		want := 28.274333882308138

		assertion(t, got, want)
	})	
}

func assertion(t testing.TB, got, want float64)  {
	t.Helper()

	if got != want {
		t.Errorf("want: %v but got: %v", want, got)
	}
}