package Interface

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
	Draw()
}
type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pow(circle.Radius, 2) * 3.14159
}
func (circle Circle) Draw() {
	fmt.Printf("circle area : %v", circle.Area())
}

type Triangle struct {
	Base, Height float64
}

func (triangle Triangle) Area() float64 {
	return triangle.Base * triangle.Height
}
func (triangle Triangle) Draw() {
	fmt.Printf("triangle area : %v", triangle.Area())
}

type Square struct {
	Length float64
}

func (square Square) Area() float64 {
	return square.Length * square.Length
}
func (square Square) Draw() {
	fmt.Printf("square area : %v", square.Area())
}
