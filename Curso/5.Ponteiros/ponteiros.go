package main

// Ponteiros são uma REFERÊNCIA DE MEMÓRIA. PONTEIROS NÃO GUARDAM VALOR, VARS FAZEM ISSO
//O valor de um ponteiro não é zero, é nil caso não tenha sido atribuído nada a ele
// Ponteiros apenas guardam / apontam para um ENDEREÇO DE MEMÓRIA a partir do uso de *int ou qualquer outro tipo
// O uso do & serve para o compilador apontar e nos dizer o endereço da parte da memória onde o valor referenciado está alocado
// Para receber o valor savo dentro do ponteiro, usamos *
import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var v1 int = 10
	var v2 int = v1
	fmt.Println(v2)

	var v3 int
	var ponteiro *int // ==> apontou para o endereço de memória

	v3 = 100
	ponteiro = &v3

	fmt.Println(v3, ponteiro) // apontou para o endereço onde está a v3 por causa do sinal "&""
	//recebido da v3
	fmt.Println(v3, *ponteiro) // desreferenciação ==> apontou para o valor da v3

	v3 = 150
	ponteiro = &v3
	fmt.Println(v3, ponteiro)
}
