package main

import (
	"fmt"
)

func main() {
	y := 42
	z := 42.0
	fmt.Printf("%v  es tipo %T \n", y, y)
	fmt.Printf("%v  es tipo %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v  es tipo  %T \n", m, m)

	// otra forma de declarar una variable
	z = float64(m)
	fmt.Printf("%v es tipo %T \n", z, z)

}
