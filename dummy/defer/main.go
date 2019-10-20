package main

import (
	"fmt"
	"time"
)

func main() {
	example()
	otherExample()
}

func example() {
	defer measure("example")()
	fmt.Println("some example code here")
}

func otherExample() {
	defer measure("other example")()
	fmt.Println("some other example code here")
}

func measure(name string) func(){
	start := time.Now()
	fmt.Printf("This function %s is starting now\n", name)
	return func() {
		fmt.Printf("The function %s is exiting at %s\n", name, time.Since(start))
	}
}

