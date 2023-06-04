package main

import(
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

// Dentro do manipulador de endpoint, recebi os valores para filtrar do contexto.
// Em seguida, ele gera a lógica do filtro para nossa matriz.
// Retorne a matriz filtrada por meio do endpoint.

type Product struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Preco float64 `json:"preco"`
}

var products = []Product{
	{1, "Produto 1", 19.99},
	{2, "Produto 2", 24.99},
	{3, "Produto 3", 29.99},
}

func GetProducts(c *gin.Context) {
	// Lógica para obter produtos

	c.JSON(http.StatusOK, gin.H{
		"message": "Lista de produtos",
		"data":    products,
	})
}

func main() {
router := gin.Default()

router.GET("/products", GetProducts) //define a rota GET e associa o manipulador GETProduts a essa rota
router.Run(":8080")// inicia o servidor na prta 8080, permitindo que o servidor receba e processe as solicitacoes HTTP

}

func GETProduts(c *gin.Context) {
	filtros := c.QueryMap("Filtros")//

	filtroProducts := make([]Product,0)
	for _, product := range products {
		if matchFiltros(product, filtros) {
			filtroProducts = append(filtroProducts, product)
		}
	}
	c.JSON(http.StatusOK, filtroProducts)

}

func matchFiltros(product Product, filtros map[string]string)bool {

	for field, value := range filtros {
		// Verificar se o valor do campo corresponde ao valor do filtro
		switch field {
		case "id":
			if fmt.Sprintf("%d", product.ID) != value {
				return false
			}
		case "nome":
			if product.Nome != value {
				return false
			}
		case "preco":
			if fmt.Sprintf("%.2f", product.Preco) != value {
				return false
			}
		default:
			// Campo desconhecido, ignorar ou tratar como necessário
		}
	}

	return true

}
