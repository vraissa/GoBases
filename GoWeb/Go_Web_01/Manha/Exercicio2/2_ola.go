package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Rota para retornar o JSON com a mensagem
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensagem": "Ola, Harry Potter, me chamo Tom Riddle",
		})
			
	})

	router.Run(":8080")
	
}