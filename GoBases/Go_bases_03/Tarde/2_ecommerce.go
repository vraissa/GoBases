package main

import "fmt"

type Usuario struct {
	Nome string	
	Sobrenome string
	Email string
	Produtos []*Produto
}

type Produto struct {
	Nome string
	Preco float64
	Quantidade int
}

func novoProduto(nome string, preco float64) *Produto {
	return &Produto {
		Nome: nome,	
		Preco: preco,
	}
}

func adicionaProduto(usuario *Usuario, produto *Produto, quantidade int) {
	produto.Quantidade = quantidade
	usuario.Produtos = append(usuario.Produtos, produto)

}
func deletarProduto(usuario *Usuario) {

	//como e para apagar tudo, pude utilizar desta forma, caso contrario seria melhor usar um array
	usuario.Produtos = nil
}

func main() {
	usuario := &Usuario {
		Nome: "Vitoria",
		Sobrenome: "Raissa",
		Email: "vitoria@teste.com",
		Produtos: nil,
	}

	//criar produto
	produto := novoProduto("teclado", 500.0)

	//adicionar produto para o usuario
	adicionaProduto(usuario, produto, 2)

	//verificar se produtos foram adicionados
	for _, p := range usuario.Produtos {
		fmt.Printf("Nome: %s, Preço: %.2f, Quantidade: %d\n", p.Nome, p.Preco, p.Quantidade)
	}

	//deletar produtos
	deletarProduto(usuario)

	//verificar se foi excluido corretamente
	fmt.Println("Produtos do usuário após exclusão:")
	fmt.Println(usuario.Produtos)
}