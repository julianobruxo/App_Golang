package main

import (
	"fmt"
	"math"
)

/*
	Então, foi criado uma interface para calcular a área do tipo retangulo e do tipo circulo para evitar ter que faze 2 funções com 2 códigos diferentes.

Para isso, foi usado oum método que calcula a area de ambos (interface) dentro de outras funções especificas deles, retornando o valor de imputado
*/
type retangulo struct {
	largura float64
	altura  float64
}

type circulo struct {
	raio float64
}
type forma interface {
	/* interfaces apenas possuem assinaturas de métodos, não é possível
	   passar campos ou criar typos com elas*/
	metodoQueCalculaAreaDasFormas() float64
}

func escreverAreaDasFormas(f forma) {
	fmt.Printf("A áerea da forma é %0.2f \n", f.metodoQueCalculaAreaDasFormas())
}

func (r retangulo) metodoQueCalculaAreaDasFormas() float64 {
	return r.largura * r.altura
}
func (c circulo) metodoQueCalculaAreaDasFormas() float64 {
	return math.Pi * (c.raio * c.raio)
}

func main() {
	r := retangulo{10, 15}
	c := circulo{20}
	escreverAreaDasFormas(r)
	escreverAreaDasFormas(c)

}
