package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var user usuario
	user.nome = "Aldo"
	user.idade = 3

	var endereco endereco
	endereco.logradouro = "Rua Tal"
	endereco.numero = 42

	fmt.Println(user)

	user2 := usuario{"Antonio", 15, endereco}

	fmt.Println((user2))

	user3 := usuario{nome: "Dalva"}

	fmt.Println(user3)

	user3_idade := usuario{idade: 41}

	fmt.Println(user3_idade)
}
