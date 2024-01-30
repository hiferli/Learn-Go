package main

import (
	"errors"
	"fmt"
)

func greet(name string) string {
	return ("Hello " + name)
}

// func division(numerator int, denominator int) (int, int) {
// 	var quotient int = numerator / denominator
// 	var remainder int = numerator % denominator
// 	return quotient, remainder
// }

func division(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		var error = errors.New("Ishaan this side! You cannot divide a number by 0. Learn some maths!")
		return -1, -1, error
	}

	var quotient int = numerator / denominator
	var remainder int = numerator % denominator

	return quotient, remainder, nil
}

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
	var name string = "Ishaan"
	fmt.Println(greet(name))

	// var num1 int = 10
	// var num2 int = 5
	// var quotient, remainder = division(num1, num2)
	// fmt.Printf("The division of %v and %v gives quotient %v and remainder %v", num1, num2, quotient, remainder)

	var num1 int = 20
	var num2 int = 0
	var quotient, remainder, error = division(num1, num2)

	if error != nil {
		fmt.Println(error.Error())
	} else {
		// No division error
		fmt.Printf("The division of %v and %v gives quotient %v and remainder %v", num1, num2, quotient, remainder)
	}

	var num int = 14
	var monthName, monthError = month(num)
	if error != nil {
		fmt.Println(monthError.Error())
	} else {
		fmt.Printf("%v is the %v month", monthName, num)
	}
}
