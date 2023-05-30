package main 

import "fmt"

type ErroSalario struct {
	Salario int
}

func (e ErroSalario) Error() string {
	return fmt.Sprintf("erro: o mínimo tributável é 15.000 e o salário informado é: %d", e.Salario)
}


 func main() {
	salarioInformado := 12000

	if salarioInformado < 15000 {
		err := ErroSalario{Salario: salarioInformado}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sera necessario o pagamento do imposto")
	}


 }