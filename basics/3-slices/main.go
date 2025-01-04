package main

import (
	"fmt"
	"time"
)

func main() {
	array()
	slice()
	slice2()
	sliceRace()
	sliceRaceV2()
}

func array() {
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)
}

func slice() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice = append(slice, 100)
	slice = append(slice, 200)
	fmt.Println("After append = ", slice)

	fmt.Println("Removing value 1 = ", slice[1:])
	fmt.Println("Removing value 100 = ", slice[:len(slice)-1])
	fmt.Println("Removing value 1, 100 and 200 = ", slice[1:5])
}

func slice2() {
	slice := make([]int, 3, 6)
	var i = 1
	for i <= 10 {
		slice = append(slice, i+1000)
		i++
	}
	fmt.Println("Dynamic slice using make = ", slice)

	slice = make([]int, 0)
	i = 1
	for i <= 10 {
		slice = append(slice, i+1000)
		i++
	}
	fmt.Println("Dynamic slice using make with 0 size = ", slice)
}

func sliceRace() {
	s := make([]int, 1)
	go func() {
		t := append(s, 1)
		fmt.Println("T = ", t)
	}()

	go func() {
		u := append(s, 2)
		fmt.Println("U = ", u)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("S = ", s)
}

func sliceRaceV2() {
	fmt.Println("Data Race problem")
	s := make([]int, 0, 1)
	go func() {
		t := append(s, 1)
		fmt.Println("T = ", t)
	}()

	go func() {
		u := append(s, 2)
		fmt.Println("U = ", u)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("S = ", s)
}
