package main

import "fmt"

func main() {
	dados := map[string]string{
		"nome":      "Claudio",
		"sobrenome": "Santos",
	}

	fmt.Println(dados)

	dados2 := map[string]map[string]string{
		"nome_completo": {
			"nome":      "Cláudio",
			"sobrenome": "Santos",
		},
	}

	fmt.Println(dados2)

	fmt.Println(dados["nome_completo"])
}
