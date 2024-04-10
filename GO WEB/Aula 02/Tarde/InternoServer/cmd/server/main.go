package main

import(
	"log"
	"github.com/gin-gonic/gin"
	"github.com/luuan11/internServer/internal/users"
	"github.com/luuan11/internServer/cmd/server/handler"
)

func main(){
	repo := users.NewRepository()
	service := users.NewService(repo)
	userHandler := handler.NewUser(service)

	server := gin.Default()
	u := server.Group("/users")
	u.POST("/", userHandler.Store())
	u.GET("/", userHandler.GetAll())
	u.GET("/:id", userHandler.GetAll())
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}