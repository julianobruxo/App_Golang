package main

import (
	"fmt"
	"sync"
	"time"
)

//WaitGroup é um recurso bem útil para sincronizar todas
//as execuções concorrentes. O exemplo acima utlizando WaitGroup
//é bem simples de entender e sempre que possível, por sua simplicidade,
//é o ideal de ser usado para sincronização de goroutines.

//Para usá-lo, usamos o pacote sync.Waitgroup do go, e dentro dos
//parenteses do Waitgroup.Add() colocamos o numero de goroutines que queremos
//que rodem ao mesmo tempo.

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 3; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
