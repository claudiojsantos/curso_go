package main

import "fmt"

func inverterSinal(numero int) int {
	novoNumero := numero * -1

	return novoNumero
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 9

	novoNumero := inverterSinal(numero)

	fmt.Println(numero)
	fmt.Println(novoNumero)

	numeroComPonteiro := -10

	fmt.Println(numeroComPonteiro)

	inverterSinalComPonteiro(&numeroComPonteiro)

	fmt.Println(numeroComPonteiro)
}
