package main

import "fmt"

func main() {
	clients := []map[string]interface{}{
		{
			"Nome":      "Giovana",
			"Idade":     23,
			"Atividade": 2,
			"Salario":   1000000,
		},
		{
			"Nome":      "Lucas",
			"Idade":     19,
			"Atividade": 1,
			"Salario":   3000,
		},
		{
			"Nome":      "Luan",
			"Idade":     21,
			"Atividade": 5,
			"Salario":   8000,
		},
	}

	for _, valor := range clients {
		nome := valor["Nome"].(string)
		fmt.Println("Clientes:", nome)

		idade, idadeOk := valor["Idade"].(int)
		if !idadeOk || idade <= 22 {
			fmt.Println("Não possui empréstimo disponível (idade insuficiente)")
			continue
		}

		atividade, atividadeOk := valor["Atividade"].(int)
		if !atividadeOk || atividade <= 1 {
			fmt.Println("Não possui empréstimo disponível (atividade insuficiente)")
			continue
		}

		salario, salarioOk := valor["Salario"].(int)
		if !salarioOk {
			fmt.Println("Erro: Salário inválido")
			continue
		}

		if salario > 100000 {
			fmt.Println("Possui empréstimo disponível sem juros")
		} else {
			fmt.Println("Possui empréstimo disponível com juros")
		}
	}
}