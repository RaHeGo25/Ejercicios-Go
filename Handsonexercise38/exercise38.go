package main

import (
	"fmt"
	"math/rand"
)

func main() {

	e := 1
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("iteración %v \t cuenta total %v \t x es %v\n", i, e, x)
			e++
		}
	}

}
