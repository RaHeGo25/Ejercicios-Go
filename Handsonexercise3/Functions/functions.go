package main

import "fmt" //Libreria fmt

func add(x int, y int) int { //funcion add, suma los valores retornados de la funcion main
	return x + y
}

func main() { // funcion main
	fmt.Println(add(42, 13)) //Regresa los valores dados a x y, e imprime el valor de la suma
}
