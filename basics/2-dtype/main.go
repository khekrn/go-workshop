package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	var a int = 10            // Signed integer
	var b uint = 20           // Unsigned integer
	var c float64 = 3.14      // Floating-point
	var d complex128 = 1 + 4i // Complex number

	fmt.Println(a, b, c, d)
	fmt.Println("- GOARCH : ", runtime.GOARCH)
	fmt.Println("- int size : ", unsafe.Sizeof(int(10)))

	bin := 0b_1010_1001
	hex := 0x_1_2_3_4_5_6_7_8
	oct := 0o_12

	fmt.Println("Binary = ", bin)
	fmt.Println("Hex = ", hex)
	fmt.Println("Oct = ", oct)

	div := 0
	fmt.Println(10 / div)
}
