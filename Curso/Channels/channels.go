package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Eu sou a Goroutine1 ", canal)
	mensagem := <-canal // essa sintaxe mostra que canal vai receber um valor
	fmt.Println(mensagem)
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // canal + <- + valor = envia dados para dentro do canal
		time.Sleep(time.Second)
	}
}
