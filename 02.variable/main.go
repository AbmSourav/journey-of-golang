package main

import (
	"fmt"
)

// global variables
var greetings string = "Hello"
// variables with initializations
var i, j int = 1, 2

func main() {
	// variable can be redeclair on defferent scope, but not in same scope.
	// with var we can reassign the value
	var greetings string = "Hello World"
	greetings = "Hello Universe"
	fmt.Println(greetings)

	// global varialbles also can be call in block scope
	fmt.Println(i, j)

	// short hand, go will find out data type
	message := 2020
	review := "What a Year !!!"
	fmt.Println(message, review)

	// variable block
	var (
		name string = "John Doe."
		designation string = "Go Developer"
		age int = 100
		salary float32 = 80000.00
	)
	fmt.Printf("%s %s, age: %d salary: %f \n", name, designation, age, salary)
}
