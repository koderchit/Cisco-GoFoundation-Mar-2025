package main

import (
	"fmt"
	"math"
)

// v1.0
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Invalid argument")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }: //any object with Area() method
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Invalid argument")
	}
}
*/

func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

// v2.0 (add perimeter calculation)
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func PrintPerimeter(x interface{ Perimeter() float64 }) {
	fmt.Println("Perimeter :", x.Perimeter())
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)
	PrintPerimeter(c)

	r := Rectangle{Height: 12, Width: 14}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)
	PrintPerimeter(r)

	s := Square{Side: 15}
	PrintArea(s)
	PrintPerimeter(s)

}
