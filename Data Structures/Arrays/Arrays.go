package main

import "fmt"

func main() {
	var nums [4]int32 = [4]int32{1, 2, 3, 4}
	fmt.Println(nums)

	strings := [5]string{"Ishaan", "Joshi", "Is", "Learning", "Go"}
	fmt.Println(strings)

	// booleans := [...]bool{false, true, false, true}

	var name []string = strings[0:2]
	// fmt.Println(booleans)
	fmt.Println(name)

	looping := []string{
		"This",
		"Is",
		"An",
		"Iteration",
	}

	// fmt.Print(looping)
	for i := 0; i < len(looping); i++ {
		fmt.Print(looping[i] + " ")
	}

	for index, str := range looping {
		fmt.Print(index)
		fmt.Print(". " + str + "\n")
	}
}
