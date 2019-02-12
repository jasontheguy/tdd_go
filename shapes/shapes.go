package main

import "math"

//Implementing a struct != class
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

//Go interface for Shapes
type Shape interface {
	Area() float64
}

//Struct methods begin with first letter of type then Type
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)

}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
