package products

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/luuan11/middleProducts/internal/entities"
	"github.com/luuan11/middleProducts/pkg/store"
	"github.com/stretchr/testify/assert"
)

func Test_GetAll(t *testing.T){
	input := []entities.Product{
		{
			ID:       1,
			Name:     "CellPhone",
			Category: "tech",
			Count:    3,
			Price:    250,
		},
		{
			ID:       1,
			Name:     "Notbook",
			Category: "tech",
			Count:    10,
			Price:    1750.5,
		},
	}
	dataJson, _ := json.Marshal(input)

	dbMock := store.Mock{
		Data: dataJson,
	}

	storeStub := store.FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	resp, _ := myRepo.GetAll()
	assert.Equal(t, input, resp)
}

type myStoreProduct struct {
	ReadCalled bool
}

func (s *myStoreProduct) Read(data interface{}) error {
	s.ReadCalled = true
	sFile := `[{
		"ID": 1,
		"Name": "Before Update",
		"Category": "tech",
		"Count":    20,     
		"Price":    100
		}]`
	return json.Unmarshal([]byte(sFile), &data)
}

func (s *myStoreProduct) Write(data interface{}) error {
	return nil
}

func Test_update(t *testing.T){
	myStoreProduct := myStoreProduct{}
	repo := NewRepository(&myStoreProduct)

	var expectedResponse []Product
	_ = myStoreProduct.Read(&expectedResponse)

	result, err := repo.UpdateName(expectedResponse[0].ID, "After Update")
	assert.Nil(t, err)
	
	fmt.Println(result)
	assert.Equal(t, "After Update", result.Name)
	assert.True(t, myStoreProduct.ReadCalled)
}
