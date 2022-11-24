package main

import (
	"fundamentos-golang/17-Aplicacao_CLI/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()

	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
