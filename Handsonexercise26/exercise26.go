package main

import (
	"fmt"
	"math/rand"
)

func init() {

	fmt.Println("Esta es la función Init")

}

func main() {
	z := rand.Intn(250)
	fmt.Printf("El valor del numero es %#v \n", z)

	switch {
	case z <= 100:
		fmt.Println("El número es menor a 100")
	case z >= 101 && z < 200:
		fmt.Println("El número esta entre 101 y 200")
	case z >= 201 && z < 200:
		fmt.Println("EL número esta entre 201 y 250")
	}

	fmt.Println(rand.Intn(3))
}
