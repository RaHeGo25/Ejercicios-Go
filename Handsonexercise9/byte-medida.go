package main

import "fmt"

type Bytesize int

const (
	_           = iota // iota un identificador especial incorporado predeclarado que simplifica la definición de constantes incrementales.
	kb Bytesize = 1 << (10 * iota)
	mb
	gb
	tb
	pb
	eb
)

func main() {
	fmt.Printf("%d \t %b \n", kb, kb)
	fmt.Printf("%d \t %b \n", mb, mb)
	fmt.Printf("%d \t %b \n", gb, gb)
	fmt.Printf("%d \t %b \n", tb, tb)
	fmt.Printf("%d \t %b \n", pb, pb)
	fmt.Printf("%d \t %b \n", eb, eb)
}
