package main

import (
	"fmt"
)

func integerType()  {
	var x int = 160
	var y int = 50
	fmt.Println( x - y )

	var a int = 10
	var b int8 = 5

	// int is a function that can change data type
	fmt.Println(a + int(b)) 
}

func stringType() {
	var greetings string = "Hello Universe"
	fmt.Println(greetings)
}

func boleanType() {
	var boo bool = true
	var result bool = 10 == 20

	fmt.Printf("%v, Data Type: %T \n", boo, boo)
	fmt.Printf("%v, Data Type: %T \n", result, result)
}

func main() {
	// integer
	integerType()

	// string
	stringType()

	// bolean
	boleanType()
}
