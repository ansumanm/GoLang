package main

import "fmt"

var increment func() int

func getIncrementFunction() func() int {
	count := 0
	increment := func() int {
		count++
		return count
	}
	return increment
}

func main() {
	increment = getIncrementFunction()
	// Calling increment from two different functions
	foo()
	bar()
}

func foo() {
	// increment := getIncrementFunction()
	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
}

func bar() {
	// increment := getIncrementFunction()
	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
	fmt.Println(increment()) // Output: 3
}


