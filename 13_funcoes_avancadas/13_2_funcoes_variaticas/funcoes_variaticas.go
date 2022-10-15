package main

import "fmt"

func total_soma(valores ...int) int {
	total := 0

	for _, numero := range valores {
		total += numero
	}

	return total
}

func main() {
	fmt.Println(total_soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
