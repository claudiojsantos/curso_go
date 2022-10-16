package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação linha de Comando"
	app.Usage = "Pega IP e Nome de Servidores na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "sistnet.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Bus ca IPs de endereços na Internet",
			Flags:  flags,
			Action: pegarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Pega o nome do servidor na Internet",
			Flags:  flags,
			Action: pegarServidores,
		},
	}

	return app
}

func pegarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func pegarServidores(c *cli.Context) {
	host := c.String("host")

	server, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range server {
		fmt.Println(servidor.Host)
	}
}
