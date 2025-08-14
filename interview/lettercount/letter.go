package main

import (
	"fmt"
	"sync"
)

func main() {
	s := "hello world akhil"

	worker := 10
	task := make(chan string)
	result := make(chan map[string]int)

	wg := sync.WaitGroup{}
	for i := 1; i <= worker; i++ {

		wg.Add(1)
		//worker start
		go countletter(task, &wg, result)
	}

	// task sending to workers
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			task <- string(s[i])
		}

	}

	// close task - looping ending
	close(task)

	go func() {
		wg.Wait()
		close(result)

	}()

	output := make(map[string]int)

	// 2 result -

	for data := range result {
		fmt.Println("data", data)
		for key, value := range data {
			output[key] += value
		}

	}

	fmt.Println(output)
}

func countletter(task chan string, wg *sync.WaitGroup, result chan map[string]int) {

	defer wg.Done()
	seen := make(map[string]int)

	for val := range task {
		seen[val]++
	}

	result <- seen

}
