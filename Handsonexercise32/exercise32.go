package main

import "fmt"

func main() {

	r := 0
	for {
		if r >= 182 {
			break
		}
		fmt.Println(r)
		r++
	}

}
