# Arrays
## Arrays are contiguous blocks of memory which are used to store a collection of data.

## Syntax

```go
var <ArrayName> [<Size>] <Datatype>
```

- Arrays are Indexed, starting from 0.
- Arrays are of fixed length.
- Arrays must contain the same datatype throughout.
- The default value in the array would be the default value of the datatype stored in the array, which is 0 for integer array, "" for string array, `false` for boolean array and so on

## Example
```go
var numbers [3] int32 = [3] int32 {1 , 2 , 3} 
// [1 2 3]

strings := [5]string{"Ishaan", "Joshi", "Is", "Learning", "Go"} 
// [Ishaan Joshi Is Learning Go]

booleans := [...]bool{false, true, false, true}
// [false true false true]
```

## Slices
### Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.

### An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

### Example:
```go
strings := [5]string{"Ishaan", "Joshi", "Is", "Learning", "Go"}
fmt.Println(strings)
// [Ishaan Joshi Is Learning Go]

var name []string = strings[0:2]
fmt.Println(name)
// [Ishaan Joshi]
```

You can read more about slices [Here](https://gobyexample.com/slices)

## Iteration
### An array can be iterated in the following way(s)
### 1. Using `for` loop and array lenth:

```go
looping := []string{
    "This",
    "Is",
    "An",
    "Iteration",
}

for i := 0; i < len(looping); i++ {
    fmt.Print(looping[i] + " ")
} // This Is An Iteration
```

### 2. Using `for` loop and `range`:
```go
looping := []string{
    "This",
    "Is",
    "An",
    "Iteration",
}

for index, str := range looping {
    fmt.Print(index)
    fmt.Print(". " + str + "\n")
} 

// 0. This
// 1. Is
// 2. An
// 3. Iteration
```

