# Functions
### Functions are lines of code which are executed whenever invoked. They may or may not take some input parameters and may or may not return a value.

### Note that the `main()` function is also a function. It is a special type of function which is invoked when the program is executed. Hence, we write all the code in the `main()` function.

## Function Structure in Go
```go
func <functionName> (<FunctionParameter(s)> <ParameterType(s)>) <ReturnType(s)> { // Note that this opening strong braces should be within the same line. IDK Why so but yeah note that!
    // Function body
    return <ReturnValue(s)> // If any
}
```
Example:
```go
func greet(name string) string {
    return ("Hello " + name);
}

fmt.Println(greet("Ishaan")) // "Hello Ishaan"
```

## Note:
- Functions are type specific. It would return an error if the types of parameters mentioned and the parameters provided doesn't match.
- Functions can return multiple values at the same time an example is shown below:
```go
package main
import "fmt"

func division(numerator int, denominator int) (int, int) {
	var quotient int = numerator / denominator
	var remainder int = numerator % denominator
	return quotient, remainder
}

func main() {
	var num1 int = 10
	var num2 int = 5
	var quotient, remainder = division(num1, num2)
	fmt.Printf("The division of %v and %v gives quotient %v and remainder %v", num1, num2, quotient, remainder)
    //The division of 10 and 5 gives quotient 2 and remainder 0
}
```

## Error Handling
### A function might return an error. An example can be a Divide by Zero error if we pass 0 as denominator in the above `division()` function. 

### Handling errors in Go can be done by a standard practice of using error values. Go code uses error values to indicate an abnormal state. It is imported using the `errors` package

### The default error value is `nil`, which shows that no error occured during the code execution. 

### As an example, the `division()` function can be modified as shown below:
```go
package main

import (
	"errors"
	"fmt"
)

func division(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		var error = errors.New("Ishaan this side! You cannot divide a number by 0. Learn some maths!")
		return -1, -1, error
	}

	var quotient int = numerator / denominator
	var remainder int = numerator % denominator

	return quotient, remainder, nil
}

func main() {
	var num1 int = 20
	var num2 int = 0
	var quotient, remainder, error = division(num1, num2)

	if error != nil {
		fmt.Println(error.Error())
	} else {
		// No division error
		fmt.Printf("The division of %v and %v gives quotient %v and remainder %v", num1, num2, quotient, remainder)

        // Ishaan this side! You cannot divide a number by 0. Learn some maths!
	}
}
```

## Switch Statements
### Similar to functions, we have switch statements. These are a set of statements defined for different conditions, which are refered as `cases`

### Syntax
```go
switch (<VariableName> (or nothing)) {
    case <Expression1>:
        <Code1>
    case <Expression2>:
        <Code2>
    .
    .
    default:
        <Code>
}
```

An example is shown below:

```go
package main

import (
	"errors"
	"fmt"
)


func month(monthNumber int) (string, error) {
	var error error

	switch monthNumber {
	case 1:
		return "January", nil
	case 2:
		return "February", nil
	case 3:
		return "March", nil
	case 4:
		return "April", nil
	case 5:
		return "May", nil
	case 6:
		return "June", nil
	case 7:
		return "July", nil
	case 8:
		return "August", nil
	case 9:
		return "September", nil
	case 10:
		return "October", nil
	case 11:
		return "November", nil
	case 12:
		return "December", nil
	default:
		error = errors.New("Ishaan this side again! Please enter a valid month number ranging from 1 to 12")
	}

	return "", error
}

func main() {
	var num int = 14
	var monthName, error = month(num)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Printf("%v is the %v month", monthName, num)
	}
}
```

## Again note that:
- There is no need to explicitly use the `break` statement in Go for switch statements. The switch breaks automatically
- We can have both conditional and unconditional switch statements