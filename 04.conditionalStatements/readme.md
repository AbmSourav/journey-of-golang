# Conditionals

## if
``` go
v := 5
if v < 6 {
	fmt.Println(v)
}
```
<br>

## if else
we also can declared a variable in if statement just before the condition. Scope of the variable is in `if` block.
``` go
if i := 14; i < 12 {
	fmt.Println(i)
} else {
	fmt.Println(i, "is greater-than 12")
}
```
<br>

## switch
``` go
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
```
<br>

declared a variable before switch condition
``` go
switch os := runtime.GOOS; os {
case "darwin":
	fmt.Println("OS X.")
case "linux":
	fmt.Println("Linux.")
default:
	fmt.Printf("%s.\n", os)
}
```
<br>

switch without condition
``` go
moment := time.Now()
switch {
case moment.Hour() < 12:
	fmt.Println("Good morning...")
case moment.Hour() < 17:
	fmt.Println("Good afternoon.")
default:
	fmt.Println("Good evening.")
}
```
