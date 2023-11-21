package main

import (
	"fmt"
	"math/rand"
)

func main() {
	z := rand.Intn(250)
	fmt.Printf("El valor del numero es %#v \n", z)
	if z <= 100 {
		fmt.Println("El numero esta entre 0 y 100 ")

	}
	if z >= 101 && z <= 200 {

		fmt.Println("El numero esta entre 101 y 200")

	}
	if z >= 201 && z <= 250 {

		fmt.Println("El numero esta entre 201 y 250")

	}
	fmt.Println(rand.Intn(3))
}
