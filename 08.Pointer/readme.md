# Pointers

Go has pointers. A pointer holds the memory address of a value. <br>
The & operator generates a pointer to its operand. The * operator denotes the pointer's underlying value. 

``` go
a := 10
fmt.Println(a)

// point to a
c := &a
fmt.Println(c) // 0xc0000b8010. memory address of a
fmt.Println(*c) // value of a

*c = 15 // it changed the value of a
fmt.Printf("a: %d, *c: %d", a, *c)
```
