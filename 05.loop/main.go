package main

import (
	"fmt"
)

func main() {
	a := 0
	for i := 0; i < 10; i++ {
		a += i
	}
	fmt.Println(a)

	// sum := 1
	// for ; sum < 10; {
	// 	sum += sum
	// 	fmt.Println(sum)
	// }

	// b := 1
	// for b < 100 {
	// 	b += b
	// 	fmt.Println(b)
	// }

	// primes := [5]int{2, 3, 5, 7, 11}
	// for i, v := range primes {
	// 	fmt.Printf("%d -- %d\n", i, v)
	// }
}
