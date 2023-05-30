package main

import (
	"fmt"
)

type ErroSalario struct {
	mensagem string
}

func (e ErroSalario) Error() string {
	return e.mensagem
}


 func main() {
	salarioInformado := 300

	if salarioInformado < 15000 {
		err := ErroSalario{"ERRO! O salario informado nao esta dentro do valor aceito"}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sera necessario o pagamento do imposto")
	}


 }