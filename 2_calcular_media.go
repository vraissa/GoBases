package main

import "fmt"
import "errors"

func mediaNotas(nota1 float64, nota2 float64, nota3 float64) (float64, error) {

	notaFinal := (nota1 + nota2 + nota3) / 3

	if notaFinal > 6 {
		fmt.Println("Parabens! Voce foi aprovado! Sua media final foi de", notaFinal)
	} else {
		fmt.Println("Voce nao foi aprovado! Sua media final foi de", notaFinal)
	} 

	if notaFinal < 0 {
		return 0, errors.New("A nota nao pode ser inferior a zero")
	} 
	
	return notaFinal, nil

}

func main() {

	var nota1 float64
	var nota2 float64
	var nota3 float64

	fmt.Println("Informe suas notas:")
	fmt.Scanln(&nota1, &nota2, &nota3)

	media, _ := mediaNotas(nota1, nota2, nota3)
	fmt.Println("Média Final:", media)


}