package main

import (
	"github.com/gin-gonic/gin" //framework gin.,
)

// Estrutura para o tema "products"
type Product struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Preco float64 `json:"preco"`
}

// Estrutura para o tema "users"
type User struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
}

// Estrutura para o tema "transactions"
type Transaction struct {
	ID     int    `json:"id"`
	Codigo string `json:"codigo"`
	Valor  float64 `json:"valor"`
}

func main() {
	router := gin.Default() // cria um objeto que e um roteador HTTP que permite definir rotas e
	// handles para construir uma aplicacao web

	router.GET("/products", GetProducts) // define uma rota HTTP especificando o path e o handler associados a rota
	router.GET("/users", GetUsers)
	router.GET("/transactions", GetTransactions)

	router.Run(":8080") //para iniciar o servidor web na porta 8080
}

//handler para retornar a lista de produtos
func GetProducts(c *gin.Context) {
	products := []Product{
		{1, "Produto 1", 19.99},
		{2, "Produto 2", 24.99},
		{3, "Produto 3", 29.99},
	}

	c.JSON(200, products)
}

func GetUsers(c *gin.Context) {
	users := []User{
		{1, "Usuário 1", "usuario1@example.com"},
		{2, "Usuário 2", "usuario2@example.com"},
		{3, "Usuário 3", "usuario3@example.com"},
	}

	c.JSON(200, users)
}

func GetTransactions(c *gin.Context) {
	transactions := []Transaction{
		{1, "ABC123", 100.50},
		{2, "DEF456", 200.75},
		{3, "GHI789", 150.25},
	}

	c.JSON(200, transactions)
}

func GetAll(c *gin.Context) {
	allData := struct {
		Products     []Product     `json:"products"`
		Users        []User        `json:"users"`
		Transactions []Transaction `json:"transactions"`
	}{
		GetProductsData(),
		GetUsersData(),
		GetTransactionsData(),
	}

	c.JSON(200, allData)
}

// Função para obter os dados dos produtos
func GetProductsData() []Product {
	products := []Product{
		{1, "Produto 1", 19.99},
		{2, "Produto 2", 24.99},
		{3, "Produto 3", 29.99},
	}

	return products
}

// Função para obter os dados dos usuários
func GetUsersData() []User {
	users := []User{
		{1, "Usuário 1", "usuario1@example.com"},
		{2, "Usuário 2", "usuario2@example.com"},
		{3, "Usuário 3", "usuario3@example.com"},
	}

	return users
}

// Função para obter os dados das transações
func GetTransactionsData()