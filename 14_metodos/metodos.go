package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) mostrarDadosUser() {
	fmt.Printf("O nome do usuário é %s e tem %d anos\n", u.nome, u.idade)
}

func (u *usuario) aniversario() {
	u.idade++
}

func main() {
	user1 := usuario{nome: "Cláudio", idade: 48}

	user1.mostrarDadosUser()

	user1.aniversario()

	user1.mostrarDadosUser()
}
