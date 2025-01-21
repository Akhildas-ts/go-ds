package main

import (
	"fmt"
	"sync"
)

// creating producer consumer program  with bufferd channel

// producer send data to buffered channel consumer recive

// challenge when the bufferd is full producer need to stop
// and consumer also get data full without deadlock

func main() {

	ch := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go producer(ch, &wg)

	wg.Add(1)
	go consumer(ch, &wg)

	wg.Wait()

}

func producer(ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 5; i++ {

		ch <- i
	}

	close(ch)

}

func consumer(ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for val := range ch {

		fmt.Println("value", val)

	}

}
