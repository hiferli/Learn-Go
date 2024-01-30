## Data Types
### Datatypes are the ways data can be stored in a variable. Some of the datatypes in GoLang are mentioned below:

<table border="1">
    <thead>
        <tr>
            <th>Data Type</th>
            <th>Description</th>
            <th>Zero Value</th>
            <th>Example</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>int</td>
            <td>Signed integer</td>
            <td>0</td>
            <td><code>var i int</code></td>
        </tr>
        <tr>
            <td>int8</td>
            <td>8-bit signed integer</td>
            <td>0</td>
            <td><code>var i int8</code></td>
        </tr>
        <tr>
            <td>int16</td>
            <td>16-bit signed integer</td>
            <td>0</td>
            <td><code>var i int16</code></td>
        </tr>
        <tr>
            <td>int32</td>
            <td>32-bit signed integer</td>
            <td>0</td>
            <td><code>var i int32</code></td>
        </tr>
        <tr>
            <td>int64</td>
            <td>64-bit signed integer</td>
            <td>0</td>
            <td><code>var i int64</code></td>
        </tr>
        <tr>
            <td>uint</td>
            <td>Unsigned integer</td>
            <td>0</td>
            <td><code>var i uint</code></td>
        </tr>
        <tr>
            <td>uint8</td>
            <td>8-bit unsigned integer</td>
            <td>0</td>
            <td><code>var i uint8</code></td>
        </tr>
        <tr>
            <td>uint16</td>
            <td>16-bit unsigned integer</td>
            <td>0</td>
            <td><code>var i uint16</code></td>
        </tr>
        <tr>
            <td>uint32</td>
            <td>32-bit unsigned integer</td>
            <td>0</td>
            <td><code>var i uint32</code></td>
        </tr>
        <tr>
            <td>uint64</td>
            <td>64-bit unsigned integer</td>
            <td>0</td>
            <td><code>var i uint64</code></td>
        </tr>
        <tr>
            <td>float32</td>
            <td>32-bit floating-point number</td>
            <td>0.0</td>
            <td><code>var f float32</code></td>
        </tr>
        <tr>
            <td>float64</td>
            <td>64-bit floating-point number</td>
            <td>0.0</td>
            <td><code>var f float64</code></td>
        </tr>
        <tr>
            <td>bool</td>
            <td>Boolean</td>
            <td>false</td>
            <td><code>var b bool</code></td>
        </tr>
        <tr>
            <td>string</td>
            <td>String</td>
            <td>"" (empty)</td>
            <td><code>var s string</code></td>
        </tr>
        <tr>
            <td>byte</td>
            <td>Alias for <code>uint8</code></td>
            <td>0</td>
            <td><code>var b byte</code></td>
        </tr>
        <tr>
            <td>rune</td>
            <td>Alias for <code>int32</code></td>
            <td>0</td>
            <td><code>var r rune</code></td>
        </tr>
        <tr>
            <td>complex64</td>
            <td>Complex number with float32 real and imaginary parts</td>
            <td>0 + 0i</td>
            <td><code>var c complex64</code></td>
        </tr>
        <tr>
            <td>complex128</td>
            <td>Complex number with float64 real and imaginary parts</td>
            <td>0 + 0i</td>
            <td><code>var c complex128</code></td>
        </tr>
        <tr>
            <td>uintptr</td>
            <td>Unsigned integer to store the address of a pointer</td>
            <td>0</td>
            <td><code>var u uintptr</code></td>
        </tr>
    </tbody>
</table>

## Constants
### To make a variable as constant in Go, add the keyword `const` in front of the datatype. Example: `const A int = 1`

## Exported Names
### A variable name with the first letter as capital is exportedable, whereas the variable names with first letter as lowercase character is not exported.

## Example:
### A common example can be the constant Pi, which is exported from the `math` package.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.pi) // undefined: math.pi
}
```

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi) // 3.141592653589793
}
```

## Note:
- Just like packages, every variable which is declared must be used, which prevents using extra memory as well as avoid making unused variables
```go
package main
import "fmt"

func main() {
	var a int = 1; // a declared and not used
}
```
- The `int` variable by default is either 32 or 64 bits depending on your system. However, `float` variable needs a description on the number of bits used, which is it must be either `float32` or `float64` and never `float`

- Just like Python, backquote string literals using (```) is printed the way it is written. However, the string literals using (") is printed in one line.

## Type Casting
### Type casting can be done in the data types like Integers. To do so, you can follow:
`var A <DataType1> = <DataType1>(B)`
### Type casting is important for arithmetic operations, in which you need to have the same datatype as operands to perform the operation
```go
package main
import "fmt"

func main() {
	var a int32 = 1;
	var b int64 = 2;
	var x = a + b;
	fmt.Println(x); // invalid operation: a + b (mismatched types int32 and int64)
}
```

```go
package main
import "fmt"

func main() {
	var a int32 = 1;
	var b int64 = 2;
	var x = a + int32(b);
	fmt.Println(x); // 3
}
```

## Rune: An introduction
### Consider the codes below:

```go
package main
import "fmt"

func main() {
	var name string = "Ishaan";
	fmt.Println(len(name)); // 6
}
```

```go
package main
import "fmt"

func main() {
	var name string = "Ишаан";
	fmt.Println(len(name)); // 10
}
```

### Why? 
### This is because the `len()` function in string returns the number of bits used in the string, and not the number of characters in the string.
### In this case, the string `Ишаан` contains non-ASCII Characters, which takes up more number of bits, hence the answer is 10 instead of 5.

### Here is where Rune comes handy
### To correct the code above, we follow the procedure as mentioned below:
```go
package main
import "fmt"
import "unicode/utf8"

func main() {
	var name string = "Ишаан";
	fmt.Println(utf8.RuneCountInString(name));
}
```

Read more about this [Here](https://go.dev/blog/strings)