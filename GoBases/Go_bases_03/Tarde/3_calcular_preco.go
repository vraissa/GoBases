package main

import (
	"fmt"
 	"sync"
 	"math"
)

//3 estruturas
type Produto struct {
	Nome    string
	Preco  float64
	Quantidade int
}

type Servico struct {
	Nome            string
	Preco          float64
	MinutosTrabalhados int
}

type Manutencao struct {
	Nome   string
	Preco float64
}

//3 funcoes
func SomarProdutos(produtos []Produto) float64 {
	total := 0.0
	for _, produto := range produtos {
		total += produto.Preco * float64(produto.Quantidade)
	}
	return total
}

func SomarServicos(servicos []Servico) float64 {
	total := 0.0
	for _ , servico := range servicos {
		minutosTrabalhados := float64(servico.MinutosTrabalhados)
		if minutosTrabalhados < 30 {
			minutosTrabalhados = 30
		}
		total += servico.Preco * (minutosTrabalhados / 60.0)
	}
	return total
}

func SomarManutencao(manutencoes []Manutencao) float64 {
	total := 0.0
	for _, manutenc := range manutencoes {
		total += manutenc.Preco
	}
	return total
}

func RoundToTwoDecimals(total float64) float64 {
	return math.Round(total * 100) / 100
}

//Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na tela
func main() {
	produtos := []Produto {
		{Nome: "mouse", Preco: 107.50, Quantidade: 1},
		{Nome: "teclado", Preco: 450.00, Quantidade: 3},
	}
	servicos := []Servico {
		{Nome: "empacotagem", Preco: 17.50, MinutosTrabalhados: 34},
		{Nome: "postagem", Preco: 45.00, MinutosTrabalhados: 13},
	}
	manutencoes := []Manutencao {
		{Nome: "reparo", Preco: 19.50},
		{Nome: "troca", Preco: 35.50},
	}

	var wg sync.WaitGroup
	wg.Add(3)

	var totalProdutos, totalServicos, totalManutencao float64

	go func() {
		defer wg.Done()
		totalProdutos = SomarProdutos(produtos)
	}()

	go func() {
		defer wg.Done()
		totalServicos = SomarServicos(servicos)
	}()

	go func() {
		defer wg.Done()
		totalManutencao = SomarManutencao(manutencoes)
	}()

	wg.Wait()

	total := totalProdutos + totalServicos + totalManutencao
	rounded := RoundToTwoDecimals(total)
	fmt.Println("Preço total:", rounded)


}