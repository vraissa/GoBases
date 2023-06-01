package main

import (
	"encoding/json" // para trabalhar com JSON
	"fmt" // para imprimir mensagens
	"io/ioutil" // para ler e escrever arquivos
	
)

func main() {
	theme := "produtos" // Tema escolhido 

	var data interface{}

	switch theme { // swithc verifica o valor da variavel
	case "produtos":
		data = []struct {
			ID         int    `json:"id"`
			Nome       string `json:"nome"`
			Cor        string `json:"cor"`
			Preco      float64 `json:"preco"`
			Estoque    int    `json:"estoque"`
			Codigo     string `json:"codigo"`
			Publicacao bool   `json:"publicacao"`
			DataCriacao string `json:"data_criacao"`
		}{
			// Dados de exemplo para o tema "produtos"
			{1, "Mouse", "Vermelho", 19.99, 10, "ABC123", true, "2023-05-25"},
			{2, "Teclado", "Azul", 24.99, 5, "DEF456", false, "2023-05-26"},
		}
	case "usuarios":
		data = []struct {
			ID          int    `json:"id"`
			Nome        string `json:"nome"`
			Sobrenome   string `json:"sobrenome"`
			Email       string `json:"email"`
			Idade       int    `json:"idade"`
			Altura      float64 `json:"altura"`
			Ativo       bool   `json:"ativo"`
			DataCriacao string `json:"data_criacao"`
		}{
			{1, "Gabriel", "Campos", "gabriel@example.com", 28, 1.85, true, "2023-05-25"},
			{2, "Vitoria", "Raissa", "vitoria@example.com", 29, 1.56, false, "2023-05-26"},
		}		
	
case "transacoes":
	data = []struct {
		ID            int    `json:"id"`
		Codigo        string `json:"codigo"`
		Moeda         string `json:"moeda"`
		Valor         float64 `json:"valor"`
		Emissor       string `json:"emissor"`
		Receptor      string `json:"receptor"`
		DataTransacao string `json:"data_transacao"`
	}{
		{1, "ABC123", "USD", 100.50, "Emissor 1", "Receptor 1", "2023-05-25"},
		{2, "DEF456", "EUR", 200.75, "Emissor 2", "Receptor 2", "2023-05-26"},
	}
default:
	fmt.Println("Tema invalido")
	return
	}
	
// Convertendo os dados para JSON
jsonData, err := json.Marshal(data) // json.Marshal converte os dados de data para formato json
if err != nil {
	fmt.Println("Erro ao gerar JSON:", err)
	return
}

// Escrevendo o JSON em um arquivo
fileName := theme + ".json"
err = ioutil.WriteFile(fileName, jsonData, 0644)// função ioutil.WriteFile é usada para escrever o conteúdo do JSON na forma de bytes
if err != nil {
	fmt.Println("Erro ao escrever arquivo:", err)
	return
}

fmt.Println("Arquivo", fileName, "gerado")

}



