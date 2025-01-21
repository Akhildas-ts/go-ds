package main

import (
	"fmt"
	"time"
)

// we need to terminate the workinglong function at a specific time with go routine

func workingLong(ch chan bool) {

	fmt.Println("before 5")

	time.Sleep(10 * time.Second)

	fmt.Println("after 5")
	ch <- true

}

func main() {

	ch := make(chan bool)

	times := 3 * time.Second

	timechan := time.After(times)

	go workingLong(ch)

	select {

	case <-ch:
		fmt.Println("go routine recived with expect time")

	case <-timechan:
		fmt.Println("go routine is closed ")
	}

}
