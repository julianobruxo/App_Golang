package main

import "fmt"

type user struct {
	nome          string
	sobrenome     string
	idade         uint8
	profissão     string
	nacionalidade string
	endereço      endereço
}
type endereço struct {
	rua    string
	número uint8
	bairro string
}

func main() {
	fmt.Println("Arquivo structs")
	var u user
	u.nome = "Juliano"
	u.sobrenome = "Tatsch da Silva"
	u.idade = 39
	u.profissão = "Professor"
	u.nacionalidade = "Brasileiro"
	fmt.Println(u)

	endereçoExemplo := endereço{"Rua 2", 45, "Bairro Glória"}

	user2 := user{"JOÃO", "TATSCH DA SILVA", 87, "APOSENTADO", "BRASILEIRO", endereçoExemplo}
	fmt.Println(user2)
}
