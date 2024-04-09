package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	id          int
	nome        string
	sobrenome   string
	email       string
	idade       int
	altura      float64
	ativo       bool
	dataCriacao string
}

func main() {
	jsonData := `{
        "id": 1,
        "nome": "Luan",
        "sobrenome": "Chaves",
        "email": "luan.fschaves@mercadolivre.com",
        "idade": 21,
        "altura": 170,
        "ativo": true,
        "dataCriacao": "18-03-2024"
    }`

	var user user
	err := json.Unmarshal([]byte (jsonData), &user)
	if err != nil{
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	fmt.Println("ID:", user.id)
    fmt.Println("Nome:", user.nome)
    fmt.Println("Sobrenome:", user.sobrenome)
}