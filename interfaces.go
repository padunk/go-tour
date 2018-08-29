package main

import (
	"fmt"
	"math"
	// "strings"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt(2))
	// v := Vertex{3, 4}

	a = f //myfloat implements abser
	// a = &v //*Vertex implements abser

	// a = v
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

package main

import (
	"fmt"
	// "math"
	// "strings"
)

type I interface {
	M()
}

type T struct {
	S string
}

// this methods means type I implements the interface I
// but we don't need to explicitly declare that it does so
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}