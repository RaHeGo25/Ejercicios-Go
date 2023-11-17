package main

import "fmt"

func saludo(x, y string) (string, string) { //una funcion puede devolver varios resultados
	return y, x
}

func main() {
	a, b := saludo("sean bienvenidos", "hola") //la funcion saludo retorna 2 strings
	fmt.Println(a, b)
}
