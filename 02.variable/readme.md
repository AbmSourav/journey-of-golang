# Variable

**Structure:** `var | variableName | type | = | value`

<br>

## Variable Scopes

``` go
// global scope variable
var greetings string = "Hello"
// variables with initializations
var i, j int = 1, 2

func main() {
	// block scope variables

	// variable can be re declair on defferent scope, but not in same scope.
	// with var we can reassign the value
	var greetings string = "Hello World"
	greetings = "Hello Universe"
	fmt.Println(greetings)

	// global varialbles also can be call in block scope
	fmt.Println(i, j)
}
```

<br>

## Short variable declarations & block
``` go
// short hand, go will find out data type
message := 2020
review := "What a Year !!!"
fmt.Println(message, review)

// variable block
var (
	name string = "John Doe."
	designation string = "Go Developer"
	age int = 100
	salary float32 = 80000.00
)
fmt.Printf("%s %s, age: %d salary: %f \n", name, designation, age, salary)
```

Go is a **strongly typed** language.
On above varible block, all varibles still have their data type after `Printf`.
Check data type: `fmt.Printf("%T %T %T", name, age, salary)`

<br>

## Zero values
Variables declared without initial values are given their *zero* values.
``` go
var i int
var f float64
var b bool
var s string
fmt.Printf("%v %v %v %q\n", i, f, b, s) // result: 0 0 false ""
```

<br>

## Constants
**Structure:** `const | name | type | = | value`

Constants are declared like variables, but with the `const` keyword. Constants cannot be declared using the `:=` syntax. Constants also can be re redeclared in different scope but can not redeclared in same scope. Consttants are also can't be reassigned.
``` go
const Pi = 3.14

func main() {
	const Greetings string = "Hello Universe"
	const Greetings string = "Hello World"
	fmt.Println(Greetings) // Error: Greetings redeclared in this block
	
	const Pi = 3.141
	fmt.Println(Pi)
}
```
