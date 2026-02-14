package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	ctxs, cance := context.WithCancel(context.Background())

	go worker(ctx, "worker 1")
	go worker(ctxs, "worker 2")
	time.Sleep(2 * time.Millisecond)
	cancel()

	time.Sleep(2 * time.Millisecond)
	cance()
}

func worker(ctx context.Context, v string) {

	for {

		select {

		case <-ctx.Done():
			fmt.Println("worker:", v, "is terminaated")
			return
		default:
			fmt.Println("worker", v, "is working")
		}
	}
}
