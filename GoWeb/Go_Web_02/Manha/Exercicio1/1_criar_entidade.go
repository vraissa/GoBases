package main

import (
	"fmt" // formata entrada e saida
	"net/http" // criar um servidor http
	"sync" // fornece primitivas de sincronização que ajudam a controlar o acesso a recursos compartilhados entre goroutines
	"encoding/json" // decodificar o corpo JSON para a estrutura da entidade
)

 type Entity struct {
	 ID int `json:"id"`
	 Data string `json:"data"`
 }

 var (
	entities []Entity
	mutex    sync.Mutex //Essa variável é usada para sincronizar o acesso ao slice 'entities'
)

func main() {
	http.HandleFunc("/create_entity", createEntityHandler) //sada para registrar um manipulador de função para um padrão de rota específico em um servidor HTTP
	http.ListenAndServe(":8080", nil) //é usada para iniciar um servidor HTTP e começar a ouvir as solicitações recebidas.

}

func createEntityHandler(w http.ResponseWriter, r *http.Request) { //w http.ResponseWriter objeto que implementa a interface, usado para escrever a resposta HTTP de volta ao cliente que fez a solicitação. No contexto dessa função w é usado para enviar a resposta de sucesso quando a entidade é criada com sucesso
	var entity Entity
	err := json.NewDecoder(r.Body).Decode(&entity)
	//metodo newDecoder cria um novo decodificador json a partir do corpo da solicitacao r.body
	//metodo decode e usado para decodificar o json jo objeto entity
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entity.ID = getNextID()//Aqui, estamos atribuindo um ID à entidade que foi decodificada do corpo da solicitação. O ID é gerado chamando a função getNextID

	mutex.Lock()//usado para garantir que apenas uma goroutines acsse o slice por ve o metodo lock e usado para suspende-la ate que o mutex seja desbloqueaDO
	entities = append(entities, entity)
	mutex.Unlock()//apos adicionado os slices entites, o mutex e debloqueado para que as outras goroutines possam acessar

	w.WriteHeader(http.StatusCreated)//definindo o código de status HTTP da resposta como 201 (Created). Isso indica que a entidade foi criada com sucesso
	fmt.Fprint(w, "Entity created successfully")
}

func getNextID() int {
	mutex.Lock()
	defer mutex.Unlock()

	if len(entities) == 0 {
		return 1
	}
	lastID := entities[len(entities)-1].ID
	return lastID + 1
}
