package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a, a[0])

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)

	// slices
	var slices []int = primes[1:4]
	slices[0] = 100
	fmt.Println(primes, slices)

	// slice literals. dynamically sized
	arr := []int{1, 2, 3}
	fmt.Println(arr)
}
