package main

import (
	"fmt"
	"runtime"
	"time"
)

func ifElse() {
	v := 5
	if v < 6 {
		fmt.Println(v)
	}

	if i := 14; i < 12 {
		fmt.Println(i)
	} else {
		fmt.Println(i, "is greater-than 12")
	}
}

func switchStatement() {
	today := time.Now().Weekday()
	switch time.Friday {
	case today:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
}

func main() {
	ifElse()
	switchStatement()
}
