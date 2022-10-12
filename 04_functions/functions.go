package main

import "fmt"

func somar(v1 int8, v2 int8) int8 {
	return v1 + v2
}

func multiplos_retornos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	multiplicacao := n1 * n2

	return soma, multiplicacao
}

func main() {
	var soma = somar(2, 5)

	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("teste")

	soma, multiplicacao := multiplos_retornos(3, 2)

	fmt.Println(soma, multiplicacao)

	nova_soma, _ := multiplos_retornos(10, 15)

	fmt.Println(nova_soma)

	_, nova_multiplicacao := multiplos_retornos(10, 15)

	fmt.Println(nova_multiplicacao)
}
