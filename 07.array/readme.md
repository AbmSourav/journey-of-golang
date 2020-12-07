# Array

The type [n]T is an array of n values of type T. Array has fixed size.
``` go
var a [2]string
a[0] = "Hello"
a[1] = "World"
```
<br>

if you asign bigger size and provide less values in the array, then it add 0 for `int` type and `""` for `string` type and will achieve that initial size
``` go
var greetings [3]string
a[0] = "Hello"
a[1] = "World"

primes := [6]int{2, 3, 5, 7, 11}
fmt.Println(greetings, primes)
// [Hello World ]
// [2 3 5 7 11, 0]
```
<br>

## Slices
A slice is formed by specifying two indices, a low and high bound, separated by a colon: 
``` go
var slices []int = primes[1:4]
```

A slice does not store any data, it just describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array. 
``` go
slices[0] = 100
fmt.Println(primes, slices) // [2 100 5 7 11] [100 5 7]
```
<br>

## Slice Literals
A slice literal is like an array literal without the length. Dynamically sized array
``` go
arr := []int{1, 2, 3}

strch := []struct {
	i int
	b bool
}{
	{2, true},
	{3, false},
	{5, true},
	{7, true},
	{11, false},
	{13, true},
}
fmt.Println(strch) // [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
```
<br>

## Slice length and capacity
The *length* of a slice is the number of elements it contains. <br>
The *capacity* of a slice is the number of elements in the underlying array, counting from the first element in the slice.
``` go
slice := []int{2, 3, 5, 7, 11, 13}
fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice) // len=6 cap=6 [2 3 5 7 11 13]

// Slice the slice to give it zero length.
slice = slice[:0]
fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice) // len=0 cap=6 []

// Extend its length.
slice = slice[:4]
fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice) // len=4 cap=6 [2 3 5 7]

slice = s[2:]
fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice) // len=2 cap=4 [5 7]
```
<br>

Creating a slice with `make` function. Dynamically-sized arrays.
``` go
a := make([]int, 5)
b := make([]int, 0, 4)
```
