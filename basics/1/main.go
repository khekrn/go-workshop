package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello Go !!")
	go sayHello()
	time.Sleep(1 * time.Second)
}

func sayHello() {
	fmt.Println("Hello from Goroutine!")
}
