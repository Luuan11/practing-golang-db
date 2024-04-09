package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Id          string `json:"id"`
	Nome        string `json:"nome"`
	Sobrenome   string `json:"sobrenome"`
	Email       string `json:"email"`
	Idade       int    `json:"idade"`
	Altura      int    `json:"altura"`
	Ativo       bool   `json:"ativo"`
	DataCriacao string `json:"datacriacao"`
}

var users = []User{
	{"1", "Luan", "Chaves", "luan.fschaves@mercadolivre.com", 21, 170, true, "18-03-2024"},
	{"2", "Giovana", "Siqueira", "teste@mercadolivre.com", 21, 170, true, "18-03-2024"},
}

func GetUserID(c *gin.Context) {

	id := c.Param("id")

	for _, item := range users {
		if item.Id == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	//erro 404
	c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()

	router.GET("/users", GetAllUsers)
	router.GET("/users/:id", GetUserID)

	router.Run()
}
