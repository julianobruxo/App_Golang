package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de Parida")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
