package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar - Return Console Application
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Console Application"
	app.Usage = "Get IP and Name"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Get IP",
			Flags:  flags,
			Action: getIP,
		},
		{
			Name:   "server",
			Usage:  "Get Server Name",
			Flags:  flags,
			Action: getServer,
		},
	}

	return app
}

func getIP(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServer(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
