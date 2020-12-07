# Defer

A defer statement defers the execution of a function until the surrounding function returns. 

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

``` go
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
```