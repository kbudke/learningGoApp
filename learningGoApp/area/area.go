package area

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

type Circle struct {
	Radius float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

func info(s Shape) {
	fmt.Println("area", s.Area())
}

// func moveToMain(Shape) float64 {
// 	if area {
// 		fmt.Print("Area can help with Calculating the area of the following shapes.")
// 		fmt.Println(area.Square)
// 	}
// }
