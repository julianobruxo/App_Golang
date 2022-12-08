package main

import (
	"fmt"
)

func main() {
	// i := 10
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j em ", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Suzy", "Juliano", "Isabelle"}
	// sempre inicia pelo indice seguido do valor
	for índiceDoArray, valorDentroDoArray := range nomes { // range percorre o array, string, int ou o que for e retorna
		// o indice e os valores dentro do tipo. Para ignorar o indice, usa-se um _
		fmt.Println(índiceDoArray, valorDentroDoArray)
	}

	for indice, letra := range "Pepeka" { // ==> range feito em uma string apenas
		// poderia ser feito com quaisquer tipos
		fmt.Println(indice, string(letra))
	}

	for _, letra := range "Pepeka" { // ==> range feito em uma string apenas
		// poderia ser feito com quaisquer tipos
		fmt.Println(string(letra))
	}
	//MAPS ==> var usuario recebe um map com chaves do tipo string e valor do tipo string
	usuario := map[string]string{ // é possível iterar sobre um map usando o for para assim
		// valor
		// percorrer as chaves e os valores dese map
		"nome": "Juliano",
		// chave	 // valor
		"sobrenome": "Silva",
	}

	for chaveDoMap, valorDoMAp := range usuario {
		fmt.Println(chaveDoMap, valorDoMAp)
	}
}
