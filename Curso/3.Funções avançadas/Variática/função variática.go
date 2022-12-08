package main

import "fmt"

// na função variática, você pode colocar N valores do mesmo tipo usando () com o nome da variavel e ...
//mais o tipo do valor seguido da descrição do valor antes das chaves
//usamos isso para não passar valor por valor dentro do parametro
func soma(numeros ...int) int { // esse int entre o paramentro e as chaves diz o que
	//será retornado da função
	total := 0
	for _, numero := range numeros { // range sempre deve ser inicializado
		total += numero // numero é o resultado da varredura nos paramentros da var números
	}
	return total
}

// ao imprimir no console, essa função retornará um slice , onde poderá haver iteração (repetição)

func main() {
	totalDaSoma := soma(1, 2, 8, 9, 4, 63, 65)
	// totalDaSoma receberá o total mais o range da var numero
	fmt.Println(totalDaSoma)
}
