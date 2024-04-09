package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Id          int
	Nome        string
	Sobrenome   string
	Email       string
	Idade       int
	Altura      float64
	Ativo       bool
	DataCriacao string
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

	fmt.Println("ID:", user.Id)
    fmt.Println("Nome:", user.Nome)
    fmt.Println("Sobrenome:", user.Sobrenome)
}