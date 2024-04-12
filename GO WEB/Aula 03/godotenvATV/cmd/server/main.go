package main

import (
	"log"
	"os"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luuan11/godotENV/cmd/server/handler"
	"github.com/luuan11/godotENV/internal/users"
	"github.com/luuan11/godotENV/pkg/store"
)

func LoggerMiddleware(ctx *gin.Context) {
	fmt.Printf("[%s] %s\n", ctx.Request.Method, ctx.Request.URL)
	ctx.Next()
}

func TokenMiddleware(ctx *gin.Context) {
	tokenEnvironment := os.Getenv("TOKEN")
	token := ctx.GetHeader("token")
	if token != tokenEnvironment {
		// status StatusUnauthorized equivalente ao 401
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token inválido",
		})
		return
	}
	ctx.Next()
}

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	db := storeUser.NewFileStore("file", "users.json")

	user := os.Getenv("MY_USER")
	password := os.Getenv("MY_PASS")

	fmt.Println("user", user, "pass", password)

	repo := users.NewRepository(db)
	service := users.NewService(repo)
	userHandler := handler.NewUser(service)

	server := gin.Default()
	u := server.Group("/users")
	u.POST("/", userHandler.Store())
	u.GET("/", userHandler.GetAll())
	server.Use(LoggerMiddleware)

	u.Use(TokenMiddleware)
	u.PUT("/:productId", userHandler.Update())
	u.PATCH("/:productId", userHandler.UpdateName())
	u.DELETE("/:productId", userHandler.Delete())

	server.Run()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
