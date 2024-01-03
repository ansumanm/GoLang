package main

import (
	"fmt"
	"sync"
)

// One subroutine prints odd numbers, another prints even numbers. Synchronize them
// So that both print natural numbers.

func printEven(evenChannel chan bool, oddChannel chan bool, outputChannel chan<- int, wg *sync.WaitGroup, max int) {
	// Wait for the signal before proceeding to print.
	defer func() {
		close(oddChannel)
		wg.Done()
	}()

	for i := 2; i < max; i += 2 {
		<-evenChannel
		outputChannel <- i
		oddChannel <- true
	}
}

func printOdd(evenChannel chan bool, oddChannel chan bool, outputChannel chan<- int, wg *sync.WaitGroup, max int) {
	defer func() {
		close(evenChannel)
		wg.Done()
	}()

	// Print and send signal
	for i := 1; i < max; i += 2 {
		<-oddChannel
		outputChannel <- i
		evenChannel <- true
	}
}

func main() {
	max := 10
	evenChannel := make(chan bool, 1)
	oddChannel := make(chan bool, 1)
	outputChannel := make(chan int, max)

	var wg sync.WaitGroup

	oddChannel <- true // Start the odd channel first

	go printEven(evenChannel, oddChannel, outputChannel, &wg, max)
	go printOdd(evenChannel, oddChannel, outputChannel, &wg, max)
	wg.Add(2)

	wg.Wait()
	close(outputChannel)

	// Print the output now.
	for x := range outputChannel {
		fmt.Println(x)
	}
	fmt.Println()
}
