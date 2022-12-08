package main

import "fmt"

// A cláusula defer (adiar)  serve para interromper / adiar a execução de uma função
//até o último momento possível

func função1() {
	fmt.Println("Executando a função 1")
}

func função2() {
	fmt.Println("Executando a função 2")
}
func main() {
	defer função1()
	função2()

}
