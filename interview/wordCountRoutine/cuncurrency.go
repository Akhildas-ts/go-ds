package main

import (
	"fmt"
	"sync"
)

// here we want to find frequency of the word, (cuncurrency--)

// so here we want to implement the (any cuncurrency pattern to imporove the speed) like worker pool, fan in fan out,
// here im implementing worker pool
func main() {

	arr := []string{"apple", "banana", "apple", "orange", "banana", "apple", "pear", "orange", "banana", "apple"}
	task := make(chan string)
	result := make(chan map[string]int)
	wg := sync.WaitGroup{}
	worker := 3

	for i := 1; i <= worker; i++ {
		wg.Add(1)
		// the the worker's
		go coundOfWord(task, &wg, result)

	}

	go func() {
		for _, val := range arr {
			// sending task to workers
			task <- val

		}
		close(task)

	}()

	go func() {
		wg.Wait()
		close(result)

	}()

	output := make(map[string]int)

	for chanData := range result {

		for key, val := range chanData {

			output[key] += val
		}

	}

	fmt.Println(output)

}

func coundOfWord(task chan string, wg *sync.WaitGroup, result chan map[string]int) {
	defer wg.Done()

	seen := make(map[string]int)
	for val := range task {

		seen[val]++
	}
	result <- seen
}
