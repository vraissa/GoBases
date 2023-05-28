package main

import "fmt"

func main() {

	var idade int
	var tempoDeEmprego int
	var salario float64
	var trabalhaAtualmente bool

	fmt.Println("Por gentileza, digite sua idade:")
	fmt.Scanln(&idade)

	fmt.Println("Esta trabalhando no momento? (true/ false)")
	fmt.Scanln(&trabalhaAtualmente)

	fmt.Println("Esta ha quantos anos nessa empresa?")
	fmt.Scanln(&tempoDeEmprego)

	fmt.Println("Qual seu salario?")
	fmt.Scanln(&salario)

	if idade >= 22 && trabalhaAtualmente && tempoDeEmprego >1 {
		if salario > 100000.0 {
			fmt.Println("Parabens! Seu emprestimo foi aprovado! Sem juros")
		} else {
			fmt.Println("Parabens! Seu emprestimo foi aprovado! Porém, com juros")
		} 

	} else {
		fmt.Println("Desculpe, sua solicitacao de emprestimo foi negada :(")
	}




}

