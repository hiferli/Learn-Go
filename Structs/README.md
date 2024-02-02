# Structs

## A `struct` (short for `struct`ure) is used to create a collection of members of different data types, into a single variable.

## While arrays are used to store multiple values of the same data type into a single variable, `struct`s are used to store multiple values of different data types into a single variable.

## A `struct` can be useful for grouping data together to create records. As an example, an employee details like Name (string), ID (number), etc can be stored in a `struct`

## Syntax
```go
type <StructName> struct {
    <FieldName1> <DataType1>
    <FieldName2> <DataType2>
    <FieldName3> <DataType3>
    .
    .
    .
}
```

## Example:
```go
type Student struct {
	Name       string
	ID         uint32
	Department string
	Present    bool
}
```

## Note
- Fields in the struct adopts a default value, which is the default value of the datatype of the field
- In case fields are declared without the field names (B), we must assign all the field names
- You can also assign methods to structs, however, they are declared explicitly out of the struct body, as shown below

```go
type Student struct {
	Name       string
	ID         uint32
	Department string
	Present    bool
}

func (student Student) markPresent() {
	student.Present = true
}
```

## Interfaces
### Inshort, interfaces are structs that stores only methods.
### Syntax:
```go
type <InterfaceName> interface {
    <MethodName1> <ReturnDataType1>
    <MethodName2> <ReturnDataType2>
    <MethodName3> <ReturnDataType3>
    .
    .
    .
}
```

# Read more about structs and interfaces [Here](https://www.golang-book.com/books/intro/9)