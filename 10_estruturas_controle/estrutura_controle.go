package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 10 {
		fmt.Println("O número é maior que 10")
	} else {
		fmt.Println("O número é menor que 10")
	}

	if numero2 := numero; numero2 < -5 {
		fmt.Println("é menor que -5")
	} else {
		fmt.Println("é maior que -5")
	}

}
