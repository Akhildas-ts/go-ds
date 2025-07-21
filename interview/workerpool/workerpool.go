package main

import (
	"fmt"
	"sync"
)

// worker pool pattern implemention..

func workerpool(task <-chan int, id int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range task {

		result <- value * id
	}
}
func main() {

	worker := 3
	numberOfTask := 10
	var wg sync.WaitGroup

	task := make(chan int, numberOfTask)
	result := make(chan int, numberOfTask)

	for i := 1; i <= worker; i++ {
		//here we started the worker's
		wg.Add(1)
		go workerpool(task, i, result, &wg)

	}

	//sending task
	for j := 1; j <= numberOfTask; j++ {
		task <- j

	}

	// then close  (BUFFERD CHANNEL!)
	close(task)

	go func() {
		wg.Wait()
		close(result)

	}()
	for val := range result {
		fmt.Println("result from worker:", val)
	}

}
