package main

import (
	"fmt"
	"strings" )

func main() {

palavra := "teste"

numeroLetras := len(palavra)
fmt.Println("A palavra", palavra, "tem", numeroLetras, "letras")


letras := strings.Split(palavra, "")
fmt.Println("Total de letras:")
for i := 0; i < len(letras); i++ {
	letra := letras[i]
	fmt.Println(letra)
}


}