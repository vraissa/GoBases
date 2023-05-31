package main

import (
	"fmt"
	"os"
)

//desenvolver a funcionalidade para poder ler o arquivo .txt
//Desenvolva o código necessário para ler os dados do arquivo chamado customers.txt
//exibir um panic ao tentar ler um arquivo que não existe
func main() {

	fileName := "customers.txt"
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("O arquivo %s não foi encontrado ou está danificado.", fileName))
	}

	defer file.Close()

	// Código para processar os dados do arquivo aqui
	fmt.Println("Arquivo lido com sucesso.")
}





