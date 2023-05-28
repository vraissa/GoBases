package main

import "fmt"

func calcularSalario(categoria string, horasTrabalhadas int) float64 {
	salarioBase := 0.0
	salario := 0.0

	switch categoria {
	case "C":
		salarioBase = 1000.0
	case "B":
		salarioBase = 1500.0
		if horasTrabalhadas > 160 {
			salarioBase += salarioBase * 0.2
		}
	case "A":
		salarioBase = 3000.0
		if horasTrabalhadas > 160 {
			salarioBase += salarioBase * 0.5
		}
	default:
		fmt.Println("Categoria invalida")
		return 0.0	

	}

	salario = float64(horasTrabalhadas) * salarioBase
	return salario

}

func main() {
	categoriaA := "A"
	horasA := 172
	salarioA := calcularSalario(categoriaA, horasA)
	fmt.Println("Salario do funcionario A e de :", salarioA)

	categoriaB := "A"
	horasB := 176
	salarioB := calcularSalario(categoriaB, horasB)
	fmt.Println("Salario do funcionario A e de :", salarioB)

	categoriaC := "A"
	horasC := 162
	salarioC := calcularSalario(categoriaC, horasC)
	fmt.Println("Salario do funcionario A e de :", salarioC)


}