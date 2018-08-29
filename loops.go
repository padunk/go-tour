package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z float64 = 1

	for z < x {
		z -= (z*z - x) / (2 * z)
		if z*z-x < 0.1 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
