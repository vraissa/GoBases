package main

import "fmt"

func calcularImposto(salario float64) float64 {

	var imposto float64

		if salario <= 50000 {
			imposto = salario * 0.17
		} else {
			imposto = salario * 0.27
		}
		return imposto

}

func main() {

	funcionario1 := 50000.0
	impostoFuncionario1 := calcularImposto(funcionario1)
	fmt.Println("Funcionario 1 tem o salario de", funcionario1, "e o imposto de", impostoFuncionario1)

	funcionario2 := 150000.0
	impostoFuncionario2 := calcularImposto(funcionario2)
	fmt.Println("Funcionario 2 tem o salario de", funcionario2, "e o imposto de", impostoFuncionario2)

}