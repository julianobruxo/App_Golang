package main

import (
	"fmt"
	auxiliar "modulo/Pacotes/Auxiliar"

	checkmail "github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("julianobruxo666@gmail.co")
	fmt.Println(erro)
}
