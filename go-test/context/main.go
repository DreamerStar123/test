package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Work cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go doWork(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("Main finished")
}
