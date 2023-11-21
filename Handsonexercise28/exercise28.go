package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("El valor de x es %v el valor de y es %v \n", x, y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("Ambos son menores a 4")
	case x > 6 && y < 6:
		fmt.Println("Ambos son mayorez que 6")
	case x >= 4 && x <= 6:
		fmt.Println("x es un numero entre 4 y 6 ")
	case y != 5:
		fmt.Println("y no es 5")
	default:
		fmt.Println("No se cumplio ninguna condicion")
	}
}
