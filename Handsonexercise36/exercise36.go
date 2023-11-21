package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for a, b := range xi {
		fmt.Printf("vuelta %v \t el valor es %v\n", a, b)
	}

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for c, d := range m {
		fmt.Printf("key %v \t el valor es %v\n", c, d)
	}
}
