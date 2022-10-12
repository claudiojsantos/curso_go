package main

import (
	"errors"
	"fmt"
)

func main() {
	var num1 int = 1000
	fmt.Println(num1)

	var num2 int8 = 100
	fmt.Println(num2)

	var num3 int16 = 10000
	fmt.Println((num3))

	var num4 int32 = 1000000000
	fmt.Println(num4)

	var flt1 float32 = 1234.32
	fmt.Println(flt1)

	var isTrue bool = true
	fmt.Println((isTrue))

	var erro error = errors.New("erro generico")
	fmt.Println(erro)
}
