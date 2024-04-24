package products

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/luuan11/middleProducts/internal/entities"
	"github.com/luuan11/middleProducts/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceGetAll(t *testing.T) {
	input := []entities.Product{
		{
			ID:       1,
			Name:     "CellPhone",
			Category: "Tech",
			Count:    3,
			Price:    250,
		}, {
			ID:       2,
			Name:     "Notebook",
			Category: "Tech",
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
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, input, result)
	assert.Nil(t, err)
}

func TestServiceGetAllError(t *testing.T) {
	expectedError := errors.New("error for GetAll")

	dbMock := store.Mock{
		Err: expectedError,
	}

	storeStub := store.FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)

}

func TestStore(t *testing.T) {
	testProduct := entities.Product{
		Name:     "CellPhone",
		Category: "Tech",
		Count:    3,
		Price:    52.0,
	}

	dbMock := store.Mock{
		Data: []byte("[]"),
	}

	storeStub := store.FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, _ := myService.Store(testProduct.Name, testProduct.Category, testProduct.Count, testProduct.Price)

	assert.Equal(t, testProduct.Name, result.Name)
	assert.Equal(t, testProduct.Category, result.Category)
	assert.Equal(t, testProduct.Price, result.Price)
	assert.Equal(t, uint64(1), result.ID)
}

func TestStoreError(t *testing.T) {
	testProduct := entities.Product{
		Name:     "CellPhone",
		Category: "Tech",
		Count:    3,
		Price:    52.0,
	}
	expectedError := errors.New("error for Storage")

	dbMock := store.Mock{
		Err: expectedError,
	}

	storeStub := store.FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.Store(testProduct.Name, testProduct.Category, testProduct.Count, testProduct.Price)

	assert.Equal(t, expectedError, err)
	assert.Equal(t, entities.Product{}, result)

}

func TestDbUpdate(t *testing.T) {

	testProduct := []entities.Product{
		{
			Name:     "Keyboard",
			Category: "Tech",
			Count:    3,
			Price:    52.0,
		},
	}

	dataJson, _ := json.Marshal(testProduct)
	_ = json.Unmarshal(dataJson, &testProduct)

	dbMock := store.Mock{
		Data: dataJson,
	}

	storeMock := store.FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}

	expectedResult := testProduct[0]
	expectedResult.Name = "Teste"

	repo := NewRepository(&storeMock)
	serv := NewService(repo)

	pd, _ := serv.Update(expectedResult.ID, expectedResult.Name, expectedResult.Category, expectedResult.Count, expectedResult.Price)

	assert.Equal(t, expectedResult, pd)
	assert.True(t, storeMock.ReadCalled)
}

func TestDbDelete(t *testing.T) {

}
