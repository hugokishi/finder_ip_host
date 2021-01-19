package main

import (
	"linecommand/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()

	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
