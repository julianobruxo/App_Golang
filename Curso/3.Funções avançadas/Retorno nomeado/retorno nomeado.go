package main

import "fmt"

func calculosMatematicos(n1 int, n2 int) (soma int, sub int) {
	// insiro entre o 1º () os valores q quero somar e no 2º() o nome das variáveis que receberão os valores

	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	//dentro das chaves, lanço a varíavel de acordo com o numero de itens da função nomeada e
	//inicializo com a função e os valores que quero atribuir
	soma, sub := calculosMatematicos(20, 10)
	fmt.Println(soma, sub)

}
