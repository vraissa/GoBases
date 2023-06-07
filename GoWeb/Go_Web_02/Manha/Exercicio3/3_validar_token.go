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

func autenticarMiddleware(c *gin.Context) {
	token := c.GetHeader("token")

	if token != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "token invalido!",
		})
		c.Abort()//utilizada para interromper o fluxo de middlewares e handlers, retorna uma resposta imediata
		return
	}

	c.Next()//indica que o controle deve ser passado para o próximo middleware ou handler na cadeia de execução
}

func main() {
	r := gin.Default()

	r.POST("/enviar-produto", autenticarMiddleware, enviarProduto)

	r.Run(":8080")
}

func enviarProduto(c *gin.Context) {
	var produto Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar o JSON da requisição"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pedido enviado com sucesso!"})
}