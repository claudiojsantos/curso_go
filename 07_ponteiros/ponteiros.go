package main

import "fmt"

func main() {
	var v1 uint
	var ponteiro *uint

	fmt.Println(v1, ponteiro)

	v1 = 100
	ponteiro = &v1

	fmt.Println(v1, ponteiro)
	fmt.Println(v1, *ponteiro)

	v1 = 150

	fmt.Println(v1, ponteiro)
	fmt.Println(v1, *ponteiro)
}
