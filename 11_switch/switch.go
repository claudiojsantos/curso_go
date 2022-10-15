package main

import "fmt"

func dia_semana(numero int) string {
	switch numero {
	case 1:
		return "Segunda"
	case 2:
		return "Segunda"
	default:
		return "Outros dias"
	}
}

func main() {
	resultado := dia_semana(2)

	fmt.Println(resultado)
}
