package main

import "fmt"

func main() {
	var cost map[string]int = make(map[string]int)

	cost["Apple"] = 100
	cost["Banana"] = 50
	cost["Pineapple"] = 60

	fmt.Printf("The cost of an apple is %v, whereas the cost of banana is %v and finally the pineapple would cost %v", cost["Apple"], cost["Banana"], cost["Pineapple"])

	delete(cost, "Pineapple")
	delete(cost, "Watermellon")

	abcd := map[byte]string{'a': "Apple", 'b': "Ball", 'c': "Cat", 'd': "Dog"}

	for character, word := range abcd {
		fmt.Println(string(character) + " -> " + word)
	}
}
