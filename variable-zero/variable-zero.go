package main

import "fmt"

func main() {
	a := 42 //Manera corta de declarar una variable
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "Felicidad" //declaracion multiple de variables
	fmt.Println(b, c, d, f)

	// otra forma de declaracion de una variable
	var g int
	fmt.Println(g)
	g = 43
	fmt.Println(g)

	// Declaracion de variable y su tipo de variable
	var h int = 44
	fmt.Println(h)
}
