package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luuan11/godotENV/cmd/server/handler"
	"github.com/luuan11/godotENV/internal/users"
)

func main(){
	_ = godotenv.Load()


	repo := users.NewRepository()
	service := users.NewService(repo)
	userHandler := handler.NewUser(service)

	server := gin.Default()
	u := server.Group("/users")
	u.POST("/", userHandler.Store())
	u.GET("/", userHandler.GetAll())

	server.Run()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}