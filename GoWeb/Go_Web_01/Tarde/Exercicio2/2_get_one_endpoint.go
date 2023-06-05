package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Theme struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var themes = []Theme{
	{ID: "121", Name: "luar"},
	{ID: "2", Name: "Theme 2"},
	{ID: "3", Name: "Theme 3"},
}

func GetThemeByID(c *gin.Context) {
	id := c.Param("id") // Obter o ID do parâmetro de caminho

	// Lógica para procurar o item pelo ID
	for _, theme := range themes {
		if theme.ID == id {
			c.JSON(http.StatusOK, theme) // Retorna o tema encontrado
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Tema não encontrado",
	})
}

func main() {
	router := gin.Default()

	// Rota para buscar um tema pelo ID
	router.GET("/themes/:id", GetThemeByID) // endpoint

	router.Run(":8080")
}