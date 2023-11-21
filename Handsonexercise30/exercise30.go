package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i <= 42; i++ {

		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Printf("Iteración %v \t x es igual a %v \n", i, x)
		case 1:
			fmt.Printf("Iteración %v \t x es igual a %v \n", i, x)
		case 2:
			fmt.Printf("Iteración %v \t x es igual a %v \n", i, x)
		case 3:
			fmt.Printf("Iteración %v \t x es igual a %v \n", i, x)
		case 4:
			fmt.Printf("Iteración %v \t x es igual a %v \n", i, x)
		}

	}
}
