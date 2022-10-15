package main

import "fmt"

func fct1() {
	fmt.Println("Função 01")
}

func fct2() {
	fmt.Println("Função 02")
}

func resultado_aluno(n1, n2 float32) bool {
	defer fmt.Println("Resultado avaliação do aluno")

	resultado := (n1 + n2) / 2

	if resultado >= 6 {
		return true
	}

	return false
}

func main() {
	defer fct1()

	fct2()

	fmt.Println(resultado_aluno(4, 8))
}
