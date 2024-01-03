package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure that wg.Done() is called.
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
	// wg.Done()
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs) // Buffered channel.
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start 3 workers
	for w:= 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	// Wait for all jobs to be finished
	wg.Wait()

	// Close results channel and print results
	close(results)
	for result := range results {
		fmt.Println(result)
	}
}