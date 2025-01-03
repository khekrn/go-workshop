package main

import (
	"context"
	"fmt"
)

func worker(ctx context.Context) {
	userID := ctx.Value("userID")
	fmt.Printf("Processing userID: %v\n", userID)
}

func main() {
	ctx := context.WithValue(context.Background(), "userID", 42)

	worker(ctx)
}
