package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 25, 28                         //Variables de tipo entero
	var f float64 = math.Sqrt(float64(x*x + y*y)) //se pasan a la variable f que ya es de tipo flotante
	var z uint = uint(f)                          //f se pasa a la variable z que ya  de tipo uint
	var r uint8 = uint8(z)

	fmt.Println(x, y, z, f, r)
}
