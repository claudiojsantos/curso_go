package main

import "fmt"

func soma_subtracao(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := soma_subtracao(10, 20)
	fmt.Println(soma, subtracao)
}
