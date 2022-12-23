package main

import "fmt"

type carro struct {
	marca        string
	modelo       string
	ano          uint
	valorDeVenda float64
	tabelaFipe   float64
}

func (c *carro) receberMarca() {
	fmt.Scan(&c.marca)
	fmt.Printf("Veículo da marca %s recebido.", c.marca)
}

func (c *carro) receberModelo() {
	fmt.Printf("Veículo modelo %s recebido.", c.modelo)
}

func (c *carro) receberAnoFabricacao() {
	fmt.Printf("Veículo modelo %d recebido.", c.ano)
}

func (c *carro) receberValorDeVenda() {
	fmt.Printf("Veículo modelo %v recebido.", c.valorDeVenda)
}

func (c *carro) recebervalorTabelaFipe() {
	fmt.Printf("Veículo modelo %v recebido.", c.tabelaFipe)
}

func (c *carro) verificarOportunidade() string {
	if c.valorDeVenda < c.tabelaFipe {
		return "Vale a pena comprar"
	} else {
		return "Não vale a pena comprar"
	}
}
func main() {
	veiculo := carro{"Peugeot", "2008", 2020, 95000, 91000}
	message := "Seu veículo da marca %v , modelo %v , do ano %v ,"
	message += "com valor de venda de %v , cuja tabela FIPE é %v, "
	message += "recebeu a seguinte avaliação:\n"
	fmt.Println(veiculo)
	fmt.Printf(message, veiculo.marca, veiculo.modelo, veiculo.ano, veiculo.valorDeVenda, veiculo.tabelaFipe)
	veiculo.verificarOportunidade()
	fmt.Println(veiculo.verificarOportunidade())
}
