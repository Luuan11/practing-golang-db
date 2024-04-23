package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuan11/middleProducts/pkg/web"
	"github.com/luuan11/middleProducts/internal/products"
)

type CreateRequestDto struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Count    int     `json:"count"`
	Price    float64 `json:"price"`
}

type UpdateRequestDto struct {
	Name     string  `json:"name" `
	Category string  `json:"category" binding:"required"`
	Count    int     `json:"count"`
	Price    float64 `json:"price"`
}

type UpdateNameRequestDto struct {
	Name string `json:"name"`
}

type ProductHandler struct {
	service products.Service
}

func NewProduct(p products.Service) *ProductHandler {
	return &ProductHandler{
		service: p,
	}
}

func GetBody[T any](body T, ctx *gin.Context) (T, error) {
	jsonbBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return body, errors.New("body is missing")
	}
	if err := json.Unmarshal(jsonbBody, &body); err != nil {
		return body, errors.New("not possible to parse the body verify json")
	}
	return body, err
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *ProductHandler) GetAll() gin.HandlerFunc {
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

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body CreateRequestDto true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (c *ProductHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateRequestDto
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		// quando chamamos a service, os dados já estarão tratados
		fmt.Println(req.Name, req.Category, req.Count, req.Price)
		p, err := c.service.Store(req.Name, req.Category, req.Count, req.Price)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, p)
	}
}

func (c *ProductHandler) Update(ctx *gin.Context) {
	var productReq UpdateRequestDto
	id := ctx.Param("id")
	productReq, err := GetBody(productReq, ctx)
	if err != nil {
		web.NewResponse(http.StatusBadRequest, nil, err.Error()).Response(ctx)
	}
	user, err := c.service.Update(id, productReq.Name, productReq.Category, productReq.Count, productReq.Price)
	if err != nil {
		web.NewResponse(http.StatusBadRequest, nil, err.Error()).Response(ctx)
	}
	web.NewResponse(http.StatusCreated, user, "").Response(ctx)
}

func (c *ProductHandler) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("productId"), 10, 0)
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
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O nome do produto é obrigatório"})
			return
		}
		p, err := c.service.UpdateName(fmt.Sprint(id), req.Name)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("productId"), 10, 0)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		err = c.service.Delete(fmt.Sprint(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		// poderiamos usar o http.StatusNoContent -> 204 tbm!
		ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("O produto %d foi removido", id)})
	}
}
