package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for a, b := range xi {
		fmt.Printf("vuelta %v \t el valor es %v\n", a, b)
	}

	fmt.Println("-----------")

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for c, d := range m {
		fmt.Printf("key %v \t el valor es %v\n", c, d)
	}
	fmt.Println("-----------")

	age1 := m["James"]
	fmt.Println("Edad de Bond", age1)
	if v, ok := m["James"]; ok {
		fmt.Println("Busqueda de Bond la edad es:", v)
	}

	age2 := m["Q"]
	fmt.Println("Edad de Q:", age2)

	if v, ok := m["Q"]; !ok {
		fmt.Println("No hay Q,es un valor cero de un Int", v)
	}

}
