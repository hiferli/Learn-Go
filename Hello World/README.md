## Now that the installation is done, it is the time for you to make your first GoLang Program.
### Here is the basic `Hello World` Progam in Go

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

## Let's understand the components in the program:
- `package main`: In Go language, the main package is a special package which is used with the programs that are executable and this package contains `main()` function.
- `import "fmt"`: `fmt` is a package in GoLang, which stands for Formating Text. (Read more: [Here](https://pkg.go.dev/fmt)) 
- `func main() {}`: `func` is a reserved keyword which is used to create a function. The `main()`` function is a special type of function and it is the entry point of the executable programs.
- `fmt.Println()`: This is the function in the `fmt` package that takes a string as input and prints it in the console. In this case, the string is "Hello, World!"

# Execution
### To execute GoLang program, you can follow 2 procedures:
## 1. go run
- Open the terminal in the root directory (Where the file `<FileName>.go` is located)
- In the terminal, run the command `go build <FileName>.go`
- A `.exe` file is created with the name `<FileName>.exe`.
- Further, in the terminal, run the command `./<FileName>`
- Your GoLang Program is executed and the results are printed in the terminal, which in this case is 'Hello, Word!'

## 2. go build and execute
- Simply, open the terminal in the root directory and run the command `go run <FileName>.go`
- Your GoLang Program is executed and the results are printed in the terminal, which in this case is 'Hello, Word!'

## Note: The difference in the 2 methods is that with `go build` builds an <i>.exe</i> which can further be executed, whereas `go run` directly runs the program and shows you the result WITHOUT actually creating the <i>.exe</i> file.