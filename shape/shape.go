package shape

import (
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	sideLenght float32
}

type Circle struct {
	radius float32
}

func (s Square) Area() float32 {
	return s.sideLenght * s.sideLenght
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func NewSquare(lenght float32) Square {
	return Square{
		sideLenght: lenght,
	}
}
func NewCircle(lenght float32) Circle {
	return Circle{
		radius: lenght,
	}
}
