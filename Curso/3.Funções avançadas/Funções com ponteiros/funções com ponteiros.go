package main

import (
	"fmt"
)

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	// agora com isso (*numero) eu irei desreferenciar o valor da variável
	//e dizer ao compilador que é um endereço de memoria, não um int
	*numero = *numero * -1
}

func main() {
	numero := 20                             // numero original na memória end. 0xc000018098
	numeroInvertido := inverterSinal(numero) // resultará em uma cópia do original (20), sem alterar o
	//original no endereço que estava, pois aponta para um novo endereço (numeroInvertido)
	fmt.Println(numeroInvertido) // end. 0xc0000180b0
	fmt.Println(numero)

}
