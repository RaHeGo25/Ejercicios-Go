package main

import "fmt"

func main() {
	a, b, c := 100, 1502, 14

	fmt.Printf("%d \t %b \t \t %#X \n", a, a, a) //%X	base 16, letra mayuscula para A-
	fmt.Printf("%d \t %b \t \t %#X \n", b, b, b)
	fmt.Printf("%d \t %b \t \t \t %#X \n", c, c, c)

}
