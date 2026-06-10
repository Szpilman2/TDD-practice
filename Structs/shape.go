package main

import (
	"fmt"
	"math"
)

type ShapeArea interface {
	Area() float64
}

type ShapePerimeter interface {
	Perimeter() float64
}

type Rectangle struct{
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}

func (r Rectangle) Perimeter() float64{
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func main() {
	fmt.Printf("this is the value: %d", 0)
}