# Maps
## Maps is a data structure which stores data in a key-value pair, where every value is associated with a key and can be accessed using the same.

## Syntax:
```go
var <MapName> map [<KeyType>] ValueType
```

## Example
```go
package main

import "fmt"

func main() {
	var cost map[string]int = make(map[string]int)

	cost["Apple"] = 100
	cost["Banana"] = 50
	cost["Pineapple"] = 60

	fmt.Printf("The cost of an apple is %v, whereas the cost of banana is %v and finally the pineapple would cost %v", cost["Apple"], cost["Banana"], cost["Pineapple"])

    // The cost of an apple is 100, whereas the cost of banana is 50 and finally the pineapple would cost 60
}
```

## Node
- If the key is not present in a map, it would return the default value for the value datatype. As an example, if a map has stirng keys and int value type, for an absent key, we would obtain `0` as a return value
- Hence, maps always returns something, even if you don't want to! To avoid this error, we can obtain a boolean value, which represent the presence of the key in the map 

```go
package main

import "fmt"

func main() {
	var name map[string]string = make(map[string]string)
	name["firstName"] = "Ishaan"
	// name["middleName"] = "" // Don't have any
	name["lastName"] = "Joshi"

	firstName, isFirstNamePresent := name["firstName"]
	middleName, isMiddleNamePresent := name["middleName"]
	lastName, isLastNamePresent := name["lastName"]

	if isFirstNamePresent && isMiddleNamePresent && isLastNamePresent {
		fmt.Printf("The name is %v %v, %v %v %v", firstName, middleName, firstName, middleName, lastName)
	} else if isFirstNamePresent && isLastNamePresent {
		fmt.Printf("The Name is %v, %v %v", lastName, firstName, lastName)
	}
}
```
- Deleting a key in map is shown below. Note that no error occurs when a key not present in the map is deleted:
```go
func main() {
	var cost map[string]int = make(map[string]int)

	cost["Apple"] = 100
	cost["Banana"] = 50
	cost["Pineapple"] = 60

	delete(cost, "Pineapple") // Deletes the key "Pineapple"
	delete(cost, "Watermellon") // No error
}
```

## Iteration in Maps
### Iteration of maps can be done using the  `range` keyword in the `for` loop. An example is shown below. Note that the order of results can be random, and DOESN'T follow the order in which the data is stored
```go
abcd := map[byte]string{
    'a': "Apple", 
    'b': "Ball", 
    'c': "Cat", 
    'd': "Dog"
}

for character, word := range abcd {
    fmt.Println(string(character) + " -> " + word)
}

// b -> Ball
// c -> Cat
// d -> Dog
// a -> Apple
```