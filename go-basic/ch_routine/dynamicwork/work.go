package main

import (
	"fmt"
)

// dynamic work pool
// high priority , low priority task distrubute by channel

func workPool(lowPriority chan int, highPriority chan int, close chan struct{}) {

	for {
		select {

		case val := <-highPriority:
			fmt.Println(val)

		case low := <-lowPriority:

			fmt.Println(low)

		case <-close:
			fmt.Println("closed ")
			return
		}

	}

}

func main() {

	lowPriority := make(chan int)
	highPriority := make(chan int)
	close := make(chan struct{})

	go workPool(lowPriority, highPriority, close)

	for i := 1; i <= 20; i++ {

		if i%2 == 0 {
			highPriority <- i
		} else {

			lowPriority <- i
		}
	}

	close <- struct{}{}

	// time.Sleep(5 * time.Second)

}
