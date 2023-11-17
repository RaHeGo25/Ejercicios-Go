package main

import "fmt"

func operacion(sum int) (x, y int) {
	x = sum * 4 / 10 //Retorna varios resultados
	y = sum - x
	return
}

func main() {
	fmt.Println(operacion(100))
}
