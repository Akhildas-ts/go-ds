package main

import (
	"fmt"
	"sync"
)

func printA(wg *sync.WaitGroup, ch chan struct{}) {
	//A print cheyhd
	fmt.Println("A")

	//Adone data passing
	ch <- struct{}{}
	wg.Done()
}
func printb(wg *sync.WaitGroup, ch chan struct{}, Bdone chan struct{}) {
	// Adone channel nnu waiting
	//Adone go shedular switch go routine

	<-ch

	fmt.Println("B")

	//Bdone <- data passcheyyum
	Bdone <- struct{}{}
	wg.Done()

}
func printc(wg *sync.WaitGroup, ch chan struct{}) {
	//bdone waiting
	//waiting- go shedular switch

	<-ch

	fmt.Println("c")
	wg.Done()

}

func main() {

	wg := sync.WaitGroup{}
	Adone := make(chan struct{})
	Bdone := make(chan struct{})

	wg.Add(3)
	go printA(&wg, Adone)
	go printb(&wg, Adone, Bdone)
	go printc(&wg, Bdone)

	wg.Wait()

}
