package products

import (
	"testing"


	"github.com/luuan11/middleProducts/pkg/store"
	"github.com/stretchr/testify/assert"
)

func NewProduct(id string, name, category string, count int, price int) *Product{
	return &Product{
		ID: id,
		Name: name,
		Category: category,
		Count: count,
		Price: float64(price),
	}
}

// atividade 1
func TestStubGetAll(t *testing.T){
	
		var expect []Product
		stubStore := store.StubStore{}
		repository := FileRepository{db: &stubStore}
		product, err := repository.GetAll()
		assert.Equal(t, expect, product)
		assert.Nil(t, err)
	
}

// atividade 2
func TestStubUpdateName(t *testing.T){
	expect := []Product{
		*NewProduct("1", "Keyboard", "Tech", 20, 100),
	}

	myMock := store.MyMockStore{
		CalledRead: 0,
		CalledWrite: 0,
		Data: 	&expect,
		MyNil:	nil,
	}

	repository := FileRepository{db: &myMock}

	repository.Update(expect[0], expect[0])

	assert.Len(t, myMock.Data, 1)
	assert.Equal(t, expect, myMock.Data)
}