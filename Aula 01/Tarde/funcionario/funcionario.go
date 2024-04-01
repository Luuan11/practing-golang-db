package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println(employees)
	fmt.Println("O funcionário Benjamim tem:", employees["Benjamin"], "anos de idade")

	maioresDeVinteUm := 0
	for _, i := range employees{
		if i > 0{
			maioresDeVinteUm++
		}
	}

	fmt.Printf("Na empresa temos um total de: %d Funcionarios", maioresDeVinteUm)

	employees["Frederico"] = 25
	delete(employees, "Pedro")
	fmt.Println("Após a saida de Pedro, temos esses funcionários", employees)
}