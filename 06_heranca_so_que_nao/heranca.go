package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso string
}

func main() {
	p1 := pessoa{
		nome:      "Antonio",
		sobrenome: "Claudio",
		idade:     15,
		altura:    170,
	}

	e1 := estudante{p1, "Fundamental"}

	fmt.Println(e1)
}
