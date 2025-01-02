package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Receiver by Value
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Receiver by Value - Modifications won't persist
func (p Person) HaveBirthday() {
	p.Age++
	fmt.Println("Happy Birthday! (Inside Method) Age is now:", p.Age)
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	p.Greet()
	p.HaveBirthday()
	fmt.Println("Outside Method Age is still:", p.Age)
}
