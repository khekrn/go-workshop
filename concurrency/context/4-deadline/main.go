package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped: ", ctx.Err())
			return
		default:
			userID := ctx.Value("userID")
			fmt.Printf("Processing userID: %v\n", userID)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "userID", 42)
	go worker(ctx)

	time.Sleep(4 * time.Second)
	fmt.Println("Main finished")
}
