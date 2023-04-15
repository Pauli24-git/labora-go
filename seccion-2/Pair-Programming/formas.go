package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Trapezoid struct {
	Base1  float64
	Base2  float64
	Height float64
}

func (t Trapezoid) Area() float64 {
	return 0.5 * (t.Base1 + t.Base2) * t.Height
}

type Polygon struct {
	Sides  int
	Length float64
}

func (p Polygon) Area() float64 {
	return (p.Length * p.Length * float64(p.Sides)) / (4 * math.Tan(math.Pi/float64(p.Sides)))
}

func main() {
	shapes := map[string]Shape{
		"triangle":  Triangle{Base: 4, Height: 5},
		"rectangle": Rectangle{Width: 10, Height: 2},
		"circle":    Circle{Radius: 3},
		"trapezoid": Trapezoid{Base1: 5, Base2: 3, Height: 4},
		"pentagon":  Polygon{Sides: 5, Length: 3},
	}

	for name, shape := range shapes {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()

		area := shape.Area()
		if area > 100 {
			panic("El Area es demasiado grande")
		}
		fmt.Printf("Shape: %s, Area: %f\n", name, area)
	}
}
