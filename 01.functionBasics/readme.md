## Function Basics

Functions call in a `main` function and `main` function calls automatically by Go compiler.
Before declaration of `main` you need to bring `package main`. function declarations starts with `func` keyword.

``` go
package main

import (
	"fmt"
)

func main() {
	message := "Hello Universe"
	fmt.Println(message)
}
```

Here I bring `main` package, so that I can use it. Imported `fmt` package, because I want to use `Println` method and it's belongs to `fmt`.

Functions can take zero or more arguments. In this example, `add` takes two parameters of `int` type. It'll also provide `int` type output.
``` go
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(100, 20))
}
```

#### Named return values
Go's return values may be named. If so, they are treated as variables defined at the top of the function. 
These names should be used to document the meaning of the return values. A `return` statement without arguments returns the named return values. This is known as a *naked* return.
``` go
func split(sum int) (x, y int) {
	x = sum / 3
	y = sum - x
	return
}

func main() {
	x, y := split(15) // x = 15 / 3 and y = 15 - x
	fmt.Printf("x is: %d\ny is: %d\n", x, y)
}
```
