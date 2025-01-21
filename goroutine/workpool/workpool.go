package main

import (
	"fmt"
	"sync"
	"time"
)

// create a work pool,

// task 6 array sum dive into 3 workers

func main() {

	sendtask := [][]int{{1, 2, 3, 4}, {2, 5, 2, 3}, {8, 2, 5, 4}, {9, 3, 5, 2, 1}}

	task := make(chan []int)
	wg := sync.WaitGroup{}
	numWorker := 3

	go func() {

		for _, arr := range sendtask {

			task <- arr
		}

		close(task)

	}()

	for i := 1; i <= numWorker; i++ {
		wg.Add(1)
		go worker(&wg, task, i)

	}

	wg.Wait()

}

func worker(wg *sync.WaitGroup, ch chan []int, id int) {

	defer wg.Done()

	for arr := range ch {

		sum := 0

		for _, val := range arr {

			sum += val

		}

		fmt.Println("worker ", id, "sum ", sum)

		time.Sleep(500 * time.Millisecond)

	}

}
