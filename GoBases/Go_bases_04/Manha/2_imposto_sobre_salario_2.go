package main

import (
	"fmt"
	"errors"
)

func verificaSalario(salario int) error {
	if salario < 15000 {
		return errors.New("ERRO! O salario digitado nao esta dentro do valor minimo")
	}
	return nil

}

 func main() {
	salarioInformado := 300

	if err := verificaSalario(salarioInformado); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Necessario realizar o pagamento de imposto")
	}


 }