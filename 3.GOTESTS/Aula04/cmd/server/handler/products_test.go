package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/luuan11/middleProducts/cmd/server/handler"
	"github.com/luuan11/middleProducts/internal/products"
	"github.com/luuan11/middleProducts/pkg/store"
	"github.com/luuan11/middleProducts/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_UpdateProducts(t *testing.T) {
	t.Run("Test_UpdateProducts_OK", func(t *testing.T) {
		var response interface{}
		r := createServerProduct()

		req, rr := createRequestTest(http.MethodPost, "/products/",
			`{
			"name": "Keyboard",
			"category": "tech",
			"count": 20,
			"price": 10
		  }`)

		  
		  r.ServeHTTP(rr, req)
		  assert.Equal(t, http.StatusCreated, rr.Code)
		  err := json.Unmarshal(rr.Body.Bytes(), &response)
		  assert.Nil(t, err)
		  
		  var dataStored products.Product
		  jsonData, err := json.Marshal(response)
		  assert.Nil(t, err)
		  
		  fmt.Println("---")
		  fmt.Println(dataStored)
		  fmt.Println("---")

		err = json.Unmarshal(jsonData, &dataStored)
		assert.Nil(t, err)

		id := dataStored.ID

		req, rr = createRequestTest(http.MethodPut, fmt.Sprintf("/products/%d", id),
		`{
			"name": "Keyboard",
			"category": "techalterado",
			"count": 20,
			"price": 10
		  }`)

		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)

		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.Nil(t, err)

		var pd products.Product
		jsonData, err = json.Marshal(response)
		assert.Nil(t, err)

		err = json.Unmarshal(jsonData, &pd)
		assert.Nil(t, err)

		req, rr = createRequestTest(http.MethodGet, fmt.Sprintf("/products/%d", id), "")
		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)

		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.Nil(t, err)

		var expectedResponse products.Product
		jsonData, err = json.Marshal(response)
		assert.Nil(t, err)

		err = json.Unmarshal(jsonData, &expectedResponse)
		assert.Nil(t, err)

		assert.Equal(t, expectedResponse, pd)
	})
}

func Test_DeleteUsers(t *testing.T) {
	t.Run("Test_UpdateUsers_OK", func(t *testing.T) {
		var response web.Response

		r := createServerProduct()

		req, rr := createRequestTest(http.MethodPost, "/products/",
		`{
			"Name": "Keyboard",
			"Category": "tech",
			"Count": 20,
			"Price": 10
		  }`)

		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		assert.Nil(t, err)

		var dataStored products.Product
		jsonData, err := json.Marshal(response.Data)
		assert.Nil(t, err)

		err = json.Unmarshal(jsonData, &dataStored)
		assert.Nil(t, err)
		id := dataStored.ID

		req, rr = createRequestTest(http.MethodDelete, fmt.Sprintf("/products/%d", id), "")
		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)

		req, rr = createRequestTest(http.MethodGet, fmt.Sprintf("/products/%d", id), "")
		r.ServeHTTP(rr, req)
		fmt.Println(rr.Body.String())
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func createServerProduct() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.PUT("/:id", p.Update())
	pr.GET("/", p.GetAll())
	pr.DELETE("/:id", p.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}