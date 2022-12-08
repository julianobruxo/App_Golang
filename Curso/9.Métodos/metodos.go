package main

import "fmt"

type usuario struct {
	nome          string
	idade         uint8
	nacionalidade string
}

func (user usuario) salvar() {
	fmt.Printf("Salvando os dados de %s no banco de dados\n", user.nome)
}
func (user usuario) maiorDeIdade() bool {
	return user.idade >= 18
}

func (user usuario) checarNacionalidade() bool {
	n1 := "Brasil"
	if user.nacionalidade == n1 {
		return true
	} else {
		return false
	}
}

func main() {
	usuario1 := usuario{"Juliano Silva", 39, "Brasil"}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Isabelle", 9, "Argentuno"}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	nacionalidadeUsuario := usuario2.checarNacionalidade()
	fmt.Println(nacionalidadeUsuario)

}
