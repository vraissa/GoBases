package main	

import (
	"fmt"
	"errors"
)

func calcSalarioMensal(horasTrabalhadas float64, valorHora float64) (float64, error) {
	if horasTrabalhadas < 80 {
		return 0, errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}

	salario := horasTrabalhadas * valorHora
	if salario >= 20000 {
		imposto := salario * 0.1
		salario -= imposto
	}

	return salario, nil

}

func main() {
	var horasTrabalhadas float64
	var valorHora float64

	fmt.Println("Informe a quantidade de horas trabalhadas no mes")
	fmt.Scan(&horasTrabalhadas)

	fmt.Println("Informe valor de horas trabalhadas no mes")
	fmt.Scan(&valorHora)

	salario, err := calcSalarioMensal(horasTrabalhadas, valorHora)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Salario mensal:", salario)

}