package main

import (
	"fmt"
	"io/ioutil"
)

type Produto struct{
	ID string			
	Preco float64
	Quantidade int
}

func main() {

	produtos := []Produto{
		{ID: "fone", Preco: 10.50, Quantidade: 2},
		{ID: "mouse", Preco: 32.40, Quantidade: 4},
		{ID: "caneca", Preco: 15.34, Quantidade: 6},
	}

	var data string	
	for _, p := range produtos {
		produtosData := fmt.Sprintf("%s;%.2f;%d\n", p.ID, p.Preco, p.Quantidade)
		data += produtosData
	}

	err := ioutil.WriteFile("Produtos.csv", []byte(data), 0644)
	if err != nil {
		fmt.Println("Erro ao gravar o arquivo:", err)
		return
	}

	fmt.Println("Arquivo gravado com sucesso!")

}