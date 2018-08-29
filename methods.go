package main

import (
	"fmt"
	"math"
	// "strings"
)

type Vertex struct {
	X, Y float64
}

// normal writing
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods with receiver argments
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointers receiver
// use so it can modify the original value
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) // modify X and Y
	fmt.Println(Abs(v))
	fmt.Println(v.Abs())
}
