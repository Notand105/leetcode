package structs

import "math"

// rectangles and circles have a method area that returns float64, ergo, they implement Shape
type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

type Triangle struct {
	width  float64
	height float64
}

func (t Triangle) Area() float64 {
	return (t.height * t.width) / 2
}
