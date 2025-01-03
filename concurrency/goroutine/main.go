package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNumbers() // Launching Goroutine
	fmt.Println("Main function continues execution")
	time.Sleep(3 * time.Second) // Give Goroutine time to finish
}
