package main

import (
	"fmt"
	// "math"
	// "strings"
)

type I interface {
	getFirstName() string
	getLastName() string
	getName() string
}

type T struct {
	firstName string
	lastName  string
	age       int
	id        int
}

func (t T) getFirstName() string {
	return t.firstName
}

func (t T) getLastName() string {
	return t.lastName
}

func (t T) getName() string {
	return t.firstName + " " + t.lastName
}

func main() {
	var i I = T{firstName: "John", lastName: "Doe", age: 42, id: 894839}
	var j I = T{firstName: "Jane", lastName: "Smith", age: 24, id: 987908}

	fmt.Printf("Hello %s\n", i.getName())
	fmt.Printf("Hello %s\n", j.getName())
}
