package main

import "github.com/gin-gonic/gin"

func main() {

	//criando um router
	router := gin.Default()

	//capturando o request get "/hello-world"
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//executando o nosso servidor na porta padrao 8080
	router.Run()
}