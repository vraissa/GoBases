package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("Benjamin tem", employees["Benjamin"])


	contador := 0
	for _, age := range employees {
		if age > 21 {
			contador++
		}
	}
	fmt.Println("Funcionarios com mais de 21 anos:", contador)

	employees["Federico"] = 25

	delete(employees, "Pedro")

	fmt.Println("Funcionários:", employees)


}
