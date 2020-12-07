# Loop

Go has only one looping construct, the `for` loop. 
The basic for loop has three components separated by semicolons: 
* Init statement: executed before the first iteration
* Condition expression: evaluated before every iteration
* Post statement: executed at the end of every iteration

``` go 
sum := 0
for i := 0; i < 10; i++ {
	sum += i
}
fmt.Println(sum)
```
<br>

The init and post statements are optional
``` go
sum := 1
for ; sum < 10; {
	sum += sum
	fmt.Println(sum)
}
```
<br>

`for` loop of Go also can be written like `while` loop of other languages.
``` go
b := 1
for b < 100 {
	b += b
}
fmt.Println(b)
```
<br>

instead of checking array length, we can use `range`
``` go
primes := [5]int{2, 3, 5, 7, 11}
for i, v := range primes {
	fmt.Println(i, " -- ", v)
}
```
