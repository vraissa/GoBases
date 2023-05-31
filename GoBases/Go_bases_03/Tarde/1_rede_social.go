package main

import "fmt"

type Usuario struct {
	Nome string
	Sobrenome string
	Idade int
	Email string
	Senha string
}

func (u * Usuario) MudarNome(nome, sobrenome string) {
	u.Nome = nome
	u.Sobrenome = sobrenome
}
func (u *Usuario) MudarIdade(idade int) {
	u.Idade = idade
}
func(u * Usuario) MudarEmail(email string) {
	u.Email = email
}
func(u * Usuario) MudarSenha(senha string) {
	u.Senha = senha
}

func main() {
	usuario := Usuario {
		Nome: "Vitoria",
		Sobrenome: "Zeferino",
		Idade: 29,
		Email: "vitoria@exemplo.com",
		Senha: "teste01",
	}

	// Exibir informações iniciais do usuário
	fmt.Println("Informações do usuário:")
	fmt.Println("Nome:", usuario.Nome, usuario.Sobrenome)
	fmt.Println("Idade:", usuario.Idade)
	fmt.Println("Email:", usuario.Email)
	fmt.Println("Senha:", usuario.Senha)

	// Alterar informações do usuário
	usuario.MudarNome("Gabriel", "Campos")
	usuario.MudarIdade(28)
	usuario.MudarEmail("gabriel@teste.com")
	usuario.MudarSenha("1234teste")

	// Exibir alteracoes
	fmt.Println("Informações do usuário:")
	fmt.Println("Nome:", usuario.Nome, usuario.Sobrenome)
	fmt.Println("Idade:", usuario.Idade)
	fmt.Println("Email:", usuario.Email)
	fmt.Println("Senha:", usuario.Senha)

}