package main

import (
	"fmt"
)

// worker pools are for when we have a queue of work to be done and miltiple concurrent workers pulling items of the queue
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	//using one worker will hog CPU usage and will get progressively slower as the batch continues due to the algorithmn
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	//usage four worker will use 4x the CPU usage since its uing multiple cores, it will be faster
	//using mulitple does not garuantee the correct order of the fib(n)

	//workers filling up the jobs queue
	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)

	//workers filling up the results queue
	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}

}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
