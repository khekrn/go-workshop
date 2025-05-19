package main

import "fmt"

func main() {
	//unbufferedChannel()
	//bufferedChannel()

	ch := make(chan string)
	defer close(ch)
	go sendChannel(ch)
	receiveChannel(ch)
}

func unbufferedChannel() {
	ch := make(chan string)
	defer close(ch)
	go func() {
		ch <- "Hello from Goroutine - unbuffered channel"
	}()
	//ch <- "Hello from Goroutine - unbuffered channel"

	msg := <-ch // Receive data from channel
	fmt.Println(msg)
}

func bufferedChannel() {
	ch := make(chan string, 2)
	defer close(ch)
	go func() {
		ch <- "Hello from Goroutine - buffered channel"
		ch <- "Hello from Goroutine - buffered channel-2"
	}()

	msg := <-ch // Receive data from channel
	fmt.Println(msg)
	// msg = <-ch // Receive data from channel
	// fmt.Println(msg)
}

func sendChannel(ch chan<- string) {
	ch <- "Sending msg from sendChannel"
}

func receiveChannel(ch <-chan string) {
	fmt.Println(<-ch)
}
