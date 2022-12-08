package main

import "fmt"

// Há duas maneiras de se declarar variáveis em GO:
// uma de forma explícita:

// É necessário sempre que ela for explícita declarar o tipo da variavel sendo usada
// (string, float, int, bool...etc)

// ou seja:

// var + nome + tipo + sinal de atribuição(=)
// var gopher string = "alguma coisa"

// e outra implícita:

// gopher := "alguma coisa"

// nas implícitas, o  GO automaticamente o tipo da variavel
// basta inicializar ela com a marmota :=

func main() {

	var texto string
	number := 55

	if number > 5 {
		texto = "é maior que 5"
	} else {
		texto = "é menor que 5"
	}

	fmt.Println(number, texto)
}
