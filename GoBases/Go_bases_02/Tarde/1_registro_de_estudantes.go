package main

import "fmt"

type Aluno struct {
	Nome string
	Sobrenome string		
	RG string
	DataAdmissao string
}

func (aluno Aluno) Detalhar() {
	fmt.Println("Nome:", aluno.Nome)
	fmt.Println("Sobrenome:", aluno.Sobrenome)
	fmt.Println("RG:", aluno.RG)
	fmt.Println("Data de admissao:", aluno.DataAdmissao)

}

func main() {
	aluno := Aluno {
		Nome: "Vitoria",
		Sobrenome: "Zeferino",
		RG: "12345",
		DataAdmissao: "10/02/1994",
	}

	aluno.Detalhar()

}