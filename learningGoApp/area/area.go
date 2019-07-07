package area

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.length * s.length
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	Area() float64
}

func info(s Shape) {
	fmt.Println(s.Area())
}

// func moveToMain(Shape) float64 {
// 	if area {
// 		fmt.Print("Area can help with Calculating the area of the following shapes.")
// 		fmt.Println(area.Square)
// 	}
// }
