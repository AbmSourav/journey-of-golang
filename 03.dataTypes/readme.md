# Data Types

## Integers
#### Unsigned Integers (no negative values)
* uint8 (0 - 255)
* uint16 (0 - 65535)
* uinit32 (0 - 4294967295)
* uint64 (0 - 18446744073709551615)

#### Signed Integers
* int8 (-128 to 127)
* int16 (-32768 to 32767)
* int32 / rune (-2147483648 to 2147483647)
* int64 (-9223372036854775808 to 9223372036854775807)

#### 3 machine dependent types
* unit (32 or 64 bits)
* int (same as unit)
* uinitptr

#### Floats
* float32
* float64

#### Complex
* complex64
* complex128

<br>

## String
* "Hello"

<br>

## Boleans
* true
* false

<br><br>

## Example
``` go
var x int = 160
var y int = 50
fmt.Println(x - y)

var outOfRange uint8 = 260
fmt.Println(outOfRange) // Error: constant 260 overflows uint8

var greetings string = "Hello Universe"
fmt.Println(greetings)

var boo bool = true
fmt.Printf("%v, Data Type: %T \n", boo, boo)
```