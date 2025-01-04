package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	basic()
	deferLifo()
	practicalUseCase()
}

func basic() {
	defer fmt.Println("This will print last")
	fmt.Println("This will print first")
}

func deferLifo() {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")
	fmt.Println("Normal Execution")
}

func practicalUseCase() {
	fmt.Println("reading file")
	file, err := os.Open("file.txt")
	if err != nil {
		panic("error while reading the file")
	}
	defer file.Close()
	data := make([]byte, 4096)
	for {
		n, err := file.Read(data)
		if errors.Is(err, io.EOF) {
			fmt.Println("Reached EOF")
			break
		}
		if n > 0 {
			fmt.Println(string(data[:n]))
		}
	}
}
