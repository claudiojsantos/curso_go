package main

import (
	"fmt"
	"time"
)

func main() {
	var i uint = 0

	for i < 10 {
		fmt.Println("Loop simples", i)
		time.Sleep(time.Second / 3)
		i++
	}

	for j := 1; j < 10; j++ {
		fmt.Println("Loop resumido", j)
		time.Sleep(time.Second / 3)
	}

	nomes := [3]string{"Aldo", "Antonio", "Dalva"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
		time.Sleep(time.Second / 2)
	}

	for _, nome := range nomes {
		fmt.Println("Nome:", nome)
		time.Sleep(time.Second / 2)
	}

	texto := "Minha família é tudo"

	for _, letra := range texto {
		fmt.Println(letra)
		time.Sleep(time.Second / 3)
	}

	for _, letra := range texto {
		fmt.Println(string(letra))
		time.Sleep(time.Second / 3)
	}

	nome_completo := map[string]string{
		"nome":      "Aldo",
		"sobrenome": "Campos",
	}

	for key, value := range nome_completo {
		fmt.Println(key, value)
		time.Sleep(time.Second / 3)
	}

	for {
		fmt.Println("loop infinito. Para parar apenas no teclado")
		time.Sleep(time.Second / 4)
	}
}
