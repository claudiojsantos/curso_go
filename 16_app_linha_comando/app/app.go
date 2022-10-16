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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Bus ca IPs de endereços na Internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "sistnet.com.br",
				},
			},
			Action: pegarIps,
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
