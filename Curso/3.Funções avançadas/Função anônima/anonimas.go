package main

import "fmt"

/* func main() {
	func() { // a função anônima não possui nome, apenas uma função autoexecutável
		fmt.Println("Olá mundo")
	}() // usa-se esse () para mandar o compilador executar qq coisa dentro dela ao final da leitura
}*/

/*
// anônima com parâmetros:
func main() {
	func(texto string) {
		fmt.Println(texto)
	}("Passando o parâmetro chamado texto") // parâmetro é atribuído no final

}*/

// anônima com RETORNO:
func main() {
	varQueIráRetornarOParametro := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto) // Sprintf é usado para concatenar informações
		// %s diz ao compilador que o parametro é uma string
		// sempre que quiser contatenar com um parâmetro, é necessário dizer ao compilador que
		// parâmetro é esse (%s string; %d numero)
	}("Passando o parâmetro chamado texto") // parâmetro é atribuído no final
	fmt.Println(varQueIráRetornarOParametro)

}
