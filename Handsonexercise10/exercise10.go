package main

import "fmt"

var y int

func main() {
	fmt.Println(y)

	z := 25
	fmt.Println(z)

	a, b := 28, "Prueba"
	fmt.Println(a, b)

	var c float32 = 54.48
	fmt.Printf("%v es tipo %T\n", c, c)

	e, f, _, g, _, h := 182, 225, 500, 414, 488, 896
	fmt.Println(e, f, g, h)

}
