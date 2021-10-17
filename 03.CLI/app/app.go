package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação linha de comando"
	app.Usage = "Busca IPs e Nomes de servidores na internet"

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags(),
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca servidores na web",
			Flags:  flags(),
			Action: buscarServidores,
		},
	}

	return app
}

func flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}
}

func buscarIps(c *cli.Context) {

	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal("Erro:", erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)

		//Retorna a lista de servidores mapeados para um IP
		addresses, erro2 := net.LookupAddr(ip.String())
		if erro2 != nil {
			log.Fatal("Erro 2:", erro2)
		}
		for _, address := range addresses {
			fmt.Println("Address Name: ", address)
		}

	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //NS - name server

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println("Servidor: ", servidor.Host)
	}

}
