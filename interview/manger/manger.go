package main

import (
	"fmt"
	"sync"
)

func main() {

	worker := 10
	wg := sync.WaitGroup{}

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go workers(i, &wg)
	}

	wg.Wait()
	fmt.Println("all worker done")

}

func workers(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker Id: ", id)

}
