package main

import (
	"fmt"
	// "math"
	// "strings"
)

func fibonacci() func(int) int {
	x, y, temp := 0, 0, 0

	return func(z int) int {
		if z == 0 {
			return 0
		} else if z == 1 {
			y = 1
			return 1
		} else {
			temp = x + y
			x = y
			y = temp
			return temp
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
