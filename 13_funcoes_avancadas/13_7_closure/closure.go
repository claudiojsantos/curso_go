package main

import "fmt"

func closure() func() {
	texto := "Texto da Funçao Clojure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Texto da Função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
