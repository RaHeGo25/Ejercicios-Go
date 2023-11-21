package main

import "fmt"

func main() {

	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			fmt.Printf(" Ciclo primario %v \t Ciclo secundario %v\n", a, b)
		}
		fmt.Println("")
	}
}
