package main

import "fmt"

func add(x, y int) int { //Cuando dos o más parámetros de función con nombre consecutivos comparten un tipo, se puede omitir el tipo de todos menos el último.
	return x * y
}

func main() {
	fmt.Println(add(10, 135))
}
