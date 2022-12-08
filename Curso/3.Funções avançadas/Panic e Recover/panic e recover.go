package main

import "fmt"

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recuperada com sucesso")
	}
}

func alunoAprovadoOuNao(nota1, nota2 float64) bool {
	defer recovery()
	media := (nota1 + nota2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Panic executado")

}

func main() {
	fmt.Println(alunoAprovadoOuNao(8, 6))
	fmt.Println("Pós execução")
}
