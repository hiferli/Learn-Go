# Routines (or, GoRoutines)

## A goroutine is an independent function that executes simultaneously in some separate lightweight threads managed by Go.
## In programming, a thread is the smallest series of related instructions involved in a process, which can involve many threads. 
## GoLang provides it to support concurrency in Go.

## To make a routine, we add the keyword `go` in front of the function call

## Example:
```go
func print(name string, count int) {
	fmt.Printf("My name is %v %v\n", name, count)
}

func main() {
	for i := 0; i < 10; i++ {
		go print("Ishaan", (i + 1))
	}
} // Output: None
```
## The above code produces no output. This is because as soon as the first `print()` function is called, the second `print()` function is called and so on until the `main()` function is completely executed, which in the end exits and ends the whole program.

## To overcome the above problem, we introduce something called `WaitGroups`

## To wait for multiple goroutines to finish, we can use a wait group.

## Example:
```go
var waitGroup = sync.WaitGroup{}

func print(name string, count int) {
	fmt.Printf("My name is %v %v\n", name, count)
	waitGroup.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go print("Ishaan", (i + 1))
	}

	waitGroup.Wait()
}

/*
    My name is Ishaan 10
    My name is Ishaan 3
    My name is Ishaan 7
    My name is Ishaan 8
    My name is Ishaan 9
    My name is Ishaan 5
    My name is Ishaan 6
    My name is Ishaan 1
    My name is Ishaan 2
    My name is Ishaan 4
*/
```

## The `waitGroup.Add(1)` line adds a counter to the wait-group and the `waitGroup.Done()` decrements the counter. Finally, `waitGroup.Wait()` checks whether counter is `0` or not and waits until it becomes `0`. 

# Mutex
## In computer programming, a mutual exclusion (mutex) is a program object that prevents multiple threads from accessing the same shared resource simultaneously. 

## An example is shown below. In this example, we can see that the order contains missing values. This is because of inappropriate use of shared resources.

## Example:
```go
var waitGroup = sync.WaitGroup{}
var order = []int{}

func print(name string, count int) {
	var result string = "My name is " + name + " " + strconv.Itoa(count)

	order = append(order, count)
	fmt.Println(result)

	waitGroup.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go print("Ishaan", (i + 1))
	}

	waitGroup.Wait()
	fmt.Printf("The final order is %v", order)
}

/*
    My name is Ishaan 10
    My name is Ishaan 3
    My name is Ishaan 1
    My name is Ishaan 6
    My name is Ishaan 5
    My name is Ishaan 8
    My name is Ishaan 4
    My name is Ishaan 7
    My name is Ishaan 9
    My name is Ishaan 2
    The final order is [10 1 4 8 9] // Missing values
*/
```

## This problem can be solved by adding a simple `mutex Lock()` on the shared data structures and then using `mutex Unlock()` to free the shared memory, as shown below:

```go
var waitGroup = sync.WaitGroup{}
var mutex = sync.Mutex{}
var order = []int{}

func print(name string, count int) {
	var result string = "My name is " + name + " " + strconv.Itoa(count)

	mutex.Lock()
	order = append(order, count)
	mutex.Unlock()

	fmt.Println(result)

	waitGroup.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go print("Ishaan", (i + 1))
	}

	waitGroup.Wait()
	fmt.Printf("The final order is %v", order)
}

/*
    My name is Ishaan 5
    My name is Ishaan 4
    My name is Ishaan 7
    My name is Ishaan 8
    My name is Ishaan 1
    My name is Ishaan 9
    My name is Ishaan 2
    My name is Ishaan 6
    My name is Ishaan 10
    My name is Ishaan 3
    The final order is [1 2 10 6 4 5 7 8 9 3] // This works now!
*/
```