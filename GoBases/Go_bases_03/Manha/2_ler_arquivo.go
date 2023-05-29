package main

import (
	"fmt"
	"io/ioutil"	
	"strings"
	"strconv"

)

type Produto struct{
	ID string			
	Preco float64
	Quantidade int
}

func main() {

	arquivoData, err := ioutil.ReadFile("Produtos.csv")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	produtos := parseProdutos(string(arquivoData))

	fmt.Printf("%-10s %-10s %-10s\n", "ID", "Preco", "Quantidade")
	for _, p := range produtos {
		fmt.Printf("%-10s %-10.2f %-10d\n", p.ID, p.Preco, p.Quantidade)
	}

	total := calculaTotal(produtos)
	fmt.Println("\nTotal:", total)

}

func parseProdutos(fileContent string) []Produto {
	lines:= strings.Split(fileContent, "\n")
	var produtos []Produto

	for _, line := range lines {
		if line != " " {
			fields := strings.Split(line, ";")
			if len(fields) == 3 {
				preco, err := strconv.ParseFloat(fields[1], 64)
				if err != nil {
					fmt.Println("Erro ao converter o preco", err)
					continue
				}
				quantidade, err := strconv.Atoi(fields[2])
				if err != nil {
					fmt.Println("Erro ao converter a quantidade", err)
					continue
				}
				
				produto := Produto {
					ID: fields[0],
					Preco: preco,
					Quantidade: quantidade,
				}
				produtos = append(produtos, produto)
			}
		}
	}

return produtos


}

func calculaTotal(produtos []Produto) float64 {
	total := 0.0
	for _, p:= range produtos {
		total += p.Preco * float64(p.Quantidade)
	}
	return total
}