package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuan11/godotENV/internal/users"
)

//implementando request
// type request struct{
// 	Name 			string   	`json:"name"`
// 	Sobrenome		string		`json:"lastname"`
// 	Email			string		`json:"email"`
// 	Age 			uint64	    `json:"age"`
// 	Status 			bool	  	`json:"status"`
// 	DateCreation 	string 		`json:"dateCreation"`
// }

type CreateRequestDto struct {
	Name 			string   	`json:"name"`
	Sobrenome		string		`json:"lastname"`
	Email			string		`json:"email"`
	Age 			uint64	    `json:"age"`
	Status 			bool	  	`json:"status"`
	DateCreation 	string 		`json:"dateCreation"`
}

type UpdateRequestDto struct {
	Name 			string   	`json:"name"`
	Sobrenome		string		`json:"lastname"`
	Email			string		`json:"email"`
	Age 			uint64	    `json:"age"`
	Status 			bool	  	`json:"status"`
	DateCreation 	string 		`json:"dateCreation"`
}

type UpdateNameRequestDto struct {
	Name string `json:"name"`
}

type UserHandler struct {
	service users.Service
}

func NewUser(p users.Service) *UserHandler {
	return &UserHandler{
		service: p,
	}
}

func (c *UserHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if len(p) == 0 {
			ctx.Status(http.StatusNoContent)
			return
		}

		ctx.JSON(http.StatusOK, p)
	}
}

func (c *UserHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateRequestDto
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		// quando chamamos a service, os dados já estarão tratados
		fmt.Println(req.Name, req.Sobrenome, req.Email, req.Age,req.Status, req.DateCreation)
		p, err := c.service.Store(req.Name, req.Sobrenome, req.Email, req.Age, req.Status, req.DateCreation)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, p)
	}
}

func (c *UserHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("userId"), 10, 0)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		var req UpdateRequestDto
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(req.Name, req.Sobrenome, req.Email, req.Age,req.Status, req.DateCreation)

		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "o nome do user é obrigatório"})
			return
		}

		if req.Sobrenome == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "o sobrenome do user é obrigatório"})
			return
		}

		if req.Email == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "o email do user é obrigatória"})
			return
		}

		if req.Age > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "a idade do user tem que ser maior de 6 anos"})
			return
		}

		if req.Age > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "a idade do user tem que ser maior de 6 anos"})
			return
		}

		if req.DateCreation == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "a data de crição precisa está preenchida"})
			return
		}

		p, err := c.service.Update(id, req.Name, req.Sobrenome, req.Email, req.Age,req.Status, req.DateCreation)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *UserHandler) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("userId"), 10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		var req UpdateNameRequestDto
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O nome do user é obrigatório"})
			return
		}
		p, err := c.service.UpdateName(id, req.Name)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *UserHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("userId"), 10, 0)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("O user %d foi removido", id)})
	}
}