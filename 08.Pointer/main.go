package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)

	// point to a
	c := &a
	fmt.Println(*c)
	// fmt.Println(c) // 0xc0000b8010. memory address

	*c = 15
	fmt.Printf("a: %d, *c: %d", a, *c)
}