package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Receiver by Pointer
func (p *Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Receiver by Pointer - Modifications will persist
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Println("Happy Birthday! (Inside Method) Age is now:", p.Age)
}

func main() {
	p := Person{Name: "Bob", Age: 30}
	p.Greet()
	p.HaveBirthday()
	fmt.Println("Outside Method Age is now:", p.Age)
}
