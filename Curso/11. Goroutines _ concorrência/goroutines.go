package main

import (
	"fmt"
	"time"
)

// CONCORRENCIA != PARALELISMO
// PARALELEISMO TEMOS 2+ TAREFAS SENDO EXECUTADAS EXATAMENTE AO MESMO TEMPO
// goroutines servem para fazer com que uma função se inicie antes de outra terminar
func main() {
	go escrever("Olá Mundo")
	escrever("Programando em Go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
