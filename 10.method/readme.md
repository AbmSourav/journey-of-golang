# Methods

Go does not have classes. However, you can define methods on types. <br>
A *method* is a function with a special receiver argument. The receiver appears in its own argument list between the `func` keyword and the method name. 

``` go
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
```
<br>

A method is just a function with a receiver argument.
``` go
func Sum(v Recipe) int {
	return v.X + v.Y
}
func main() {
	result := Recipe{4, 6}
	fmt.Println(Sum(result))
}
```
