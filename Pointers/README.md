# Pointers
## A pointer holds the memory address of a value.

## Syntax
```go
var <PointerName> * <DataType> = & <VariableName>
```

## Usage
- The `*` symbol is used to signify that this variable is a pointer
- The `&` symbol represents that the pointer is refering to the memory address of the variable mentioned.
- `*` also acts as a dereferencing operator. When attached in front of a pointer, it returns the value stored at the memory address of the pointer it is pointing.
- Pointers are useful to share data types between functions, which reduces making copies of datatype and hence saves storage.

## Example:
```go
var a int32 = 10
var b *int32 = &a
fmt.Printf("The memory address a is %v\n", b)
fmt.Printf("The value stored at the memory address %v is %v", b, *b)
// The memory address a is 0xc00000a0b8 (Can vary system to system)
// The value stored at the memory address 0xc00000a0b8 is 10
```
## Note:
- Pointers store the address of the memory where data is stored.
- When declared a pointer, the address it pointing to is assigned the default value, according to the data type the pointer is pointing. 
- A pointer must be referenced while initiating, else, would return a <b>Null Pointer Exception</b>.
- If a value stored in the memory location of a pointer is changed, the variable value would also be changed. This is because <b>Pointers point to the value at a memory address, hence changing the value pointed by a pointer would change the value stored by the variable.</b>

