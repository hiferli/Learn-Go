package main

import "fmt"

func main() {
	var a int32 = 10
	var b *int32 = &a
	fmt.Printf("The memory address a is %v\n", b)
	fmt.Printf("The value stored at the memory address %v is %v\n", b, *b)

	*b = 20
	fmt.Printf("The memory address a is still %v\n", b)
	fmt.Printf("After changing the value pointed by the pointer, the value stored at the memory address %v is %v\n", b, *b)
}
