package structsmethodsandinterfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Area() float64
}

// In Go interface resolution is implicit. If the type you pass in
// matches what the interface is asking for, it will compile

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
