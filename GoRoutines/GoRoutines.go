package main

import (
	"fmt"
	"strconv"
	"sync"
)

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
