package main

import (
	"fmt"
	"errors"
)

const(
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

type operationFunc func(...int) (float64, error)

func operacao(tipoDeOp string) (operationFunc, error) {

	switch tipoDeOp {
	case minimum:
		return minimumFunc, nil
	case average:
		return averageFunc, nil
	case maximum:
		return maximoFunc, nil
	default:
		return nil, errors.New("calculo com erro")
	}
}

func minimumFunc(numeros...int)(float64, error) {
	if len(numeros) == 0 {
		return 0.0, errors.New("nenhum numero fornecido, tente novamente")
	}

	min := numeros[0]
	for _, num := range numeros {
		if num < min {
		min = num
		}
	}
	return float64(min), nil
}

func averageFunc(numeros...int)(float64, error) {
	if len(numeros) == 0 {
		return 0.0, errors.New("nenhum numero fornecido, tente novamente")
	}

	soma := 0
	for _, num := range numeros{
		soma += num
	}

	average := float64(soma) / float64(len(numeros))
	return average, nil

}

func maximoFunc(numeros...int)(float64, error) {
	if len(numeros) == 0 {
		return 0.0, errors.New("nenhum numero fornecido, tente novamente")
	}

	max:= numeros[0]
	for _, num := range numeros {
		if num > max {
			max = num
		}
	}
	return float64(max), nil

}

func main() {
	minimoFunc, err := operacao(minimum)
	if err != nil {
		fmt.Println(err)
	} else {
		minValor, _ := minimoFunc(2,3,3,4,10,2,4,5)
		fmt.Println("Valor minimo e: ", minValor)
	}

	avgFunc, err := operacao(average)
	if err != nil {
		fmt.Println(err)
	}else {
		avgValor, _ := avgFunc(2,3,3,4,1,2,4,5)
		fmt.Println("O valor medio e", avgValor)
	}

	maxFunc, err := operacao(maximum)
	if err != nil {
		fmt.Println(err) 
		} else {
			maxValor, _ := maxFunc(2,3,3,4,1,2,4,5)
			fmt.Println("O valor maximo e", maxValor)
		}
	}



