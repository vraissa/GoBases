package main

import (
	"fmt"
	"errors"
	"os"
)

type Cliente struct {
	Arquivo     string
	Nome        string
	RG          string
	NumeroTelefone int
	Endereco    string
}

func gerarID() (string, error) {

	ID := "12345"

	if ID == " " {
		return " ", errors.New("Erro ID vazio")
	}
	return ID, nil
	
}

func cadastrarCliente(cliente Cliente) {
	fmt.Println("Cliente cadastrado com sucesso.")
}

func lerCliente() error {
	fileName := "customers.txt"
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Erro: o arquivo %s não foi encontrado ou está danificado", fileName)
	}
	defer file.Close()

	fmt.Println("Dados dos clientes lidos com sucesso.")
	return nil
}

func main() {
	defer func() { //adia a execução de uma função anônima até o final da função 
		if r := recover(); r != nil {
			fmt.Println("erro recuprado", r)
		}
	}

}
