package main

import (
	"fmt"
)

func main() {
	numerito := 42
	fmt.Printf("42 es en binario, %b \n", numerito)
	fmt.Printf("42 es en hexadecimal, %x \n", numerito)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	fmt.Printf("0 es en binario, %b y en hexadecimal, %x \n", a, a)
	fmt.Printf("1 es en binario, %b y en hexadecimal, %x \n", b, b)
	fmt.Printf("2 es en binario, %b y en hexadecimal, %x \n", c, c)
	fmt.Printf("3 es en binario, %b y en hexadecimal, %x \n", d, d)
	fmt.Printf("4 es en binario, %b y en hexadecimal, %x \n", e, e)
	fmt.Printf("5 es en binario, %b y en hexadecimal, %x \n", f, f)

}
