package main

import "fmt"

type usuario struct {
	nome          string
	idade         uint8
	nacionalidade string
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados de %s no banco de dados\n", u.nome)
}
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u usuario) checarNacionalidade() bool {
	n1 := "Brasil"
	if u.nacionalidade == n1 {
		return true
	} else {
		return false
	}
}

func main() {
	user := usuario{"Juliano Silva", 17, "Brasil"}
	fmt.Println(user)
	user.salvar()

	//verificar idade
	maiorDeIdade := user.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	//verificar nacionalidade
	nacionalidadeUsuario := user.checarNacionalidade()
	fmt.Println(nacionalidadeUsuario)

}
