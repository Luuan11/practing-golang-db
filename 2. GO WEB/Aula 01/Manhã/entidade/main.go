package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type user struct{
	Id          int
	Nome        string
	Sobrenome   string
	Email       string
	Idade       int
	Altura      float64
	Ativo       bool
	DataCriacao string
}

func GetAll(c *gin.Context){
	jsonData := `[{
        "id": 1,
        "nome": "Luan",
        "sobrenome": "Chaves",
        "email": "luan.fschaves@mercadolivre.com",
        "idade": 21,
        "altura": 170,
        "ativo": true,
        "dataCriacao": "18-03-2024"
    }]`

	var users []user

	err := json.Unmarshal([]byte(jsonData), &users)
	if err != nil{
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}
	c.JSON(200, users)

}

func main() {
	router := gin.Default()

	router.GET("/users", GetAll)
	router.Run()
}