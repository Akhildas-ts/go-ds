package main

import (
	"fmt"
	"sync"
	"time"
)

type bank struct {
	balance int
	mu      *sync.Mutex
}

func (b *bank) deposit(money int, ch chan struct{}) {
	ch <- struct{}{}
	b.mu.Lock()
	b.balance += money
	b.mu.Unlock()
}
func (b *bank) withdraw(money int, ch chan struct{}) {

	b.mu.Lock()
	b.balance -= money
	b.mu.Unlock()
	<-ch
	close(ch)

}
func (b *bank) getbalance() int {

	b.mu.Lock()
	b.mu.Unlock()
	return b.balance
}

func main() {

	ch := make(chan struct{})
	b := bank{
		mu: &sync.Mutex{},
	}

	var wg sync.WaitGroup

	wg.Add(4)
	go func() {
		defer wg.Done()
		b.deposit(100, ch)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(50 * time.Millisecond)
		fmt.Println(b.getbalance())
	}()
	go func() {

		defer wg.Done()
		b.withdraw(50, ch)
	}()

	go func() {

		time.Sleep(500 * time.Microsecond)
		defer wg.Done()
		fmt.Println(b.getbalance())
	}()

	wg.Wait()

}
