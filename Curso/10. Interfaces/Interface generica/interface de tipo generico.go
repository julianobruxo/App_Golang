package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func soma() int {
	a := 10
	b := 10
	return a + b
}

func main() {
	generica("String")
	generica(1)
	generica(true)
	generica(soma())

}
