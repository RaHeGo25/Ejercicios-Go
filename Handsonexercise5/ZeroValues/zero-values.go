package main

import "fmt"

func main() {
	var i int                               //variable de tipo entero
	var f float64                           //variable tipo float de 64 bits
	var b bool                              //variable tipo booleana
	var s string                            //variable de tipo string
	fmt.Printf("%v %v %v %q\n", i, f, b, s) //%v imprime los valores por default %q sintaxis de Go
}
