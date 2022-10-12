package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array [5]string

	fmt.Println(array)

	array[0] = "Posicao01"
	array[1] = "Posicao02"
	array[2] = "Posicao03"
	array[3] = "Posicao04"
	array[4] = "Posicao05"

	fmt.Println(array)

	slice := []int{}

	fmt.Println(slice)

	slice = append(slice, 10)
	slice = append(slice, 20)
	slice = append(slice, 30)
	slice = append(slice, 40)

	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))

	slice2 := array[1:3]

	fmt.Println(slice2)
}
