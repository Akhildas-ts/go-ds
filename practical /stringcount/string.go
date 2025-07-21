package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	arr1 := []int{4, 7, 12, 15, 23, 2, 17}
	ch := make(chan int)
	sum := 0
	res := 0
	wg.Add(1)
	go func() {
		for i := 0; i < len(arr1); i++ {
			sum += arr1[i]
		}
		// time.Sleep(5 * time.Second)
		ch <- sum
		wg.Done()
	}()
	wg.Wait()
	res = <-ch

	fmt.Println(res)
}
