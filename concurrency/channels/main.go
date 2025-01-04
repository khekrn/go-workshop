package main

import "fmt"

func main() {
	unbufferedChannel()
	bufferedChannel()
}

func unbufferedChannel() {
	ch := make(chan string)
	defer close(ch)
	go func() {
		ch <- "Hello from Goroutine - unbuffered channel"
	}()

	msg := <-ch // Receive data from channel
	fmt.Println(msg)
}

func bufferedChannel() {
	ch := make(chan string, 1)
	defer close(ch)
	go func() {
		ch <- "Hello from Goroutine - buffered channel"
	}()

	msg := <-ch // Receive data from channel
	fmt.Println(msg)
}
