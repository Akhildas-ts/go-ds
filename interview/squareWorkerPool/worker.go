package main

import (
	"fmt"
	"sync"
)

// create a worker pool pattern that take a n integers then sqaure the intger return with channel,

func worker(task chan int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range task {

		result <- val * val
	}

}

func main() {

	task := make(chan int)
	result := make(chan int)
	wg := sync.WaitGroup{}

	workers := 3

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(task, result, &wg)

	}

	go func() {
		for i := 1; i <= 10; i++ {
			task <- i
		}
		close(task)
	}()
	go func() {
		wg.Wait()
		close(result)
	}()

	for val := range result {
		fmt.Println("square :", val)
	}

}
