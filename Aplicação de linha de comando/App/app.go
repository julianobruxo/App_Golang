package app

import "github.com/urfave/cli"

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e Nomes de servidor na Internet"
	return app

}
