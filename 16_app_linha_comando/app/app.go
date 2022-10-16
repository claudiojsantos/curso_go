package app

import "github.com/urfave/cli"

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação linha de Comando"
	app.Usage = "Pega IP e Nome de Servidores na Internet"

	return app
}
