package main

import (
	"aplicacao/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Inicio")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
