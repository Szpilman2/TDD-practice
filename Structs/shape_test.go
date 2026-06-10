package main

import "testing"

func TestArea(t *testing.T)  {
	areaTests := []struct {
		shape ShapeArea
		want float64
	}{
		{Circle{Radius: 3}, 28.274333882308138},
		{Rectangle{Height: 12.0, Width: 6.0}, 72.0},
		{Triangle{Base: 6.0, Height: 12.0}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		want := tt.want

		assertion(t, got, want)
	}
}


func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape ShapePerimeter
		want float64
	}{
		{Circle{Radius: 3}, 18.84955592153876},
		{Rectangle{Height: 10.0, Width: 10.0}, 40.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		want := tt.want

		assertion(t, got, want)
	}
}