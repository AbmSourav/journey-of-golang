package main

import (
	"fmt"
	"time"
)

func loop(name string) {
    for i := 0; i < 3; i++ {
        fmt.Println(name, ":", i)
		time.Sleep(time.Second * 2)
    }
}

func main() {
	go loop("Abm Sourav")
	loop("Sourav")
	
	fmt.Println("Done")
}
