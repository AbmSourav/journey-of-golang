# Goroutine

A *goroutine* is a lightweight thread managed by the Go runtime. <br>
With Goroutine we can do different things *parallelly* or *asynchronously*. It's a way of doing *non-blocking* things.

<br>

Goroutine starts with `go` keyword. Go has a build-in goroutine, it exicutes `main` function. So if we create a *goroutine*, that'll be the second goroutine.
Here, in `main` function the evaluation of those three functions happen in the core goroutine. But the execution of `go loop("Abm Sourav")` happens in a new *goroutine*.
``` go
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
```
