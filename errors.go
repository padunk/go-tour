package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, error
	}
	var z float64 = 1

	for z < x {
		z -= (z*z - x) / (2 * z)
		if z*z-x < 0.1 {
			break
		}
	}
	return z, error
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
