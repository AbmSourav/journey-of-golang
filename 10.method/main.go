package main

import (
	"fmt"
)

type Recipe struct {
	X, Y int
}

func (v Recipe) Sum() int {
	return v.X + v.Y
}

func main() {
	result := Recipe{4, 6}
	fmt.Println(result.Sum())
}
