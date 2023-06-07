package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Produto struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Tipo       string  `json:"tipo"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
}

func enviarProduto(c *gin.Context) {
	var produto Produto 
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao decodificar o JSON da requisição",
		})
		return
	}
	if produto.Nome == "" {
		retornarErro(c, "nome")
		return
	}
	if produto.Tipo == "" {
		retornarErro(c, "tipo")
		return
	}
	if produto.Quantidade == 0{
		retornarErro(c, "quantidade")
		return
	}
	if produto.Preco == 0.0 {
		retornarErro(c, "preco")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Produtos enviados com sucesso!"})

}

func retornarErro(c *gin.Context, campo string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Campo " + campo + " faltando", "message": "Campo " + campo + " é obrigatório"})
}

func main() {
	r := gin.Default()

	r.POST("/enviar-produtos", enviarProduto) 
		
	r.Run(":8080")
}