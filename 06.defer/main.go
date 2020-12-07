package main

import "fmt"

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Before loop actually starts")

	defer loop()

	fmt.Println("Because of defer this line executed befoe the loop.")
}
