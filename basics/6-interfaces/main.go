package main

import "fmt"

// Interface
type Walker interface {
	Walk() string
}

// Structs implementing the interface
type Human struct{}

func (h Human) Walk() string {
	return "Human is walking"
}

type Robot struct{}

func (r Robot) Walk() string {
	return "Robot is walking"
}

// Function accepting Walker interface
func StartWalking(w Walker) {
	fmt.Println(w.Walk())
}

func main() {
	h := Human{}
	r := Robot{}

	StartWalking(h) // Output: Human is walking
	StartWalking(r) // Output: Robot is walking
}
