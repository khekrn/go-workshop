package main

import "fmt"

func main() {
	fmt.Println("Demystifying Byte")
	byteSnippet()
	fmt.Println("Demystifying Rune")
	runeSnippet()
}

func byteSnippet() {
	var b byte = 'A' // ASCII value of 'A' is 65
	fmt.Printf("Value: %c, ASCII: %d, Type: %T\n", b, b, b)
}

func runeSnippet() {
	var r rune = '😊' // Unicode code point for '😊'
	fmt.Printf("Value: %c, Unicode: %U, Type: %T\n", r, r, r)
}

func stringRep() {
	str := "Go😊"

	// Iterate as bytes
	fmt.Println("Bytes:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d: %x\n", i, str[i])
	}

	// Iterate as runes
	fmt.Println("\nRunes:")
	for i, r := range str {
		fmt.Printf("%d: %c (%U)\n", i, r, r)
	}
}
