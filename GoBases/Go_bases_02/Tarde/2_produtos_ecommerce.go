package main

import (
	"fmt"
)

type Produto struct {
	Tipo  string
	Nome  string
	Preco float64
}

func (p Produto) CalcularCusto() float64 {
	switch p.Tipo {
	case "Pequeno":
		return p.Preco
	case "Médio":
		return p.Preco * 1.03
	case "Grande":
		return p.Preco*1.06 + 2500
	default:
		return p.Preco
	}
}

type ProdutoInterface interface {
	CalcularCusto() float64
}

type Loja struct {
	Produtos []ProdutoInterface
}

func (l *Loja) Adicionar(p ProdutoInterface) {
	l.Produtos = append(l.Produtos, p)
}

func (l Loja) Total() float64 {
	total := 0.0
	for _, p := range l.Produtos {
		total += p.CalcularCusto()
	}
	return total
}

func NovoProduto(tipo, nome string, preco float64) ProdutoInterface {
	return &Produto{
		Tipo:  tipo,
		Nome:  nome,
		Preco: preco,
	}
}

func NovaLoja() *Loja {
	return &Loja{}
}

func main() {
	loja := NovaLoja()

	produto1 := NovoProduto("Pequeno", "Produto 1", 100)
	produto2 := NovoProduto("Médio", "Produto 2", 200)
	produto3 := NovoProduto("Grande", "Produto 3", 300)
	loja.Adicionar(produto1)
	loja.Adicionar(produto2)
	loja.Adicionar(produto3)

	precoTotal := loja.Total()
	fmt.Println("Preço Total:", precoTotal)
}
