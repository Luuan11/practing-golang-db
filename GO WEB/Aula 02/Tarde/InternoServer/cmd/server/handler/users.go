package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luuan11/internServer/internal/users"
)

//implementando request
type request struct{
	ID 				int		    `json:"id"`
	Name 			string   	`json:"name"`
	Sobrenome		string		`json:"lastname"`
	Email			string		`json:"email"`
	Age 			int	    	`json:"age"`
	Status 			int	  		`json:"status"`
	DateCreation 	string 		`json:"dateCreation"`
}

//estrutura controller
type UserHandler struct{
	service users.Service
}

func NewUser(u users.Service) *UserHandler{
	return &UserHandler{
		service: u,
	}
}

func (c * User)GetAll() gin.HandlerFunc{
	return func(ctx *gin.Context){
		token := ctx.Request.Header.Get("token")
			if token != "123456"{
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"Error": "Token Invalido",
				})
				return
			}
			
			u, err := c.service.GetAll()
			if err != nil{
				ctx.JSON(http.StatusNotFound, gin.H{
					"error" : err.Error(),
				})
				return
			}

			if len(u) == 0{
				ctx.Status(http.StatusNoContent)
			}
		}
}