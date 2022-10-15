package main

import "fmt"

func recuperacao_panic() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func avaliacao_final(n1, n2 float32) bool {
	defer recuperacao_panic()

	resultado := (n1 + n2) / 2

	if resultado < 6 {
		return false
	} else if resultado > 6 {
		return true
	} else {
		panic("A MÉDIA FOI EXATAMENTE 6")
	}
}

func main() {
	fmt.Println(avaliacao_final(5, 7))
}
