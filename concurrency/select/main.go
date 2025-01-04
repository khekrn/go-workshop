package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine for Channel 1
	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch1 <- "Message from Channel 1"
		}
	}()

	// Goroutine for Channel 2
	go func() {
		for {
			time.Sleep(2 * time.Second)
			ch2 <- "Message from Channel 2"
		}
	}()

	// Handle Ctrl+C for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Continuous select without explicit loop
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-stop:
			fmt.Println("Received interrupt signal. Exiting program.")
			return
		}
	}
}
