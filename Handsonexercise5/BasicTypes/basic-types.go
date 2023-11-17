package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = true //El ejemplo muestra variables de varios tipos
	Notbe  bool       = false
	MaxInt uint64     = 1<<64 - 1
	MinInt uint8      = 1<<8 - 1
	MidInt uint32     = 1<<32 - 1
	z      complex128 = cmplx.Sqrt(-25 + 28i)
)

func main() {
	fmt.Printf("Tipo de variable: %T Valor: %v\n", ToBe, ToBe)
	fmt.Printf("Tipo de variable: %T Valor: %v \n", Notbe, Notbe)
	fmt.Printf("Tipo de variable: %T Valor: %v\n", MaxInt, MaxInt)
	fmt.Printf("Tipo de variable: %T Valor: %v\n", MidInt, MidInt)
	fmt.Printf("Tipo de variable: %T Valor: %v\n", MinInt, MinInt)
	fmt.Printf("Tipo de variavle: %T Valor: %v\n", z, z)
}
