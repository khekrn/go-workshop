package main

import (
	"fmt"
)

func main() {
	snippet1()
	snippet2()
}

func snippet1() {
	m := map[string]int{"Alice": 25, "Bob": 30}
	fmt.Println(m)
}

func snippet2() {
	m := make(map[string]int)
	m["Alice"] = 25
	m["Bob"] = 30
	m["Vince"] = 65

	fmt.Println("Snippet 2, \nMap = ", m)

	fmt.Println("Taking Alice = ", m["Alice"])
	if value, exist := m["Bob"]; exist {
		fmt.Println("Taking Bob = ", value)
	}

	fmt.Println("Iterating Map")
	for key, value := range m {
		fmt.Println(key, " = ", value)
	}

	delete(m, "Alice")
	fmt.Println("Deleting Alice = ", m)
}
