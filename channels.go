package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	c := make(chan string)
// 	go boring("boring!", c)

// 	for i := 0; i < 5; i++ {
// 		// receive expression is just a value
// 		fmt.Printf("you say: %q\n", <-c)
// 	}
// }

// func boring(msg string, c chan string) {
// 	for i := 0; ; i++ {
// 		// expression to be sent can be any suitable value
// 		c <- fmt.Sprintf("%s %d", msg, i)
// 		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// 	}
// }

// Refactored
func main() {
	c := boring("boring!")
	joe := boring("joe")
	ann := boring("ann")

	for i := 0; i < 5; i++ {
		fmt.Printf("you say %q\n", <-c)
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("i'm leaving")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
