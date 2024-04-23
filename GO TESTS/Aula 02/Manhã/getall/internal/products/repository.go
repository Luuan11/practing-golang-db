package products

import "github.com/luuan11/middleProducts/pkg/store"

type Repository interface {
	GetAll() ([]Product, error)
	Store(name, category string, count int, price float64) (Product, error)
	Update(pd, produto Product) Product
	UpdateName(id string, name string) (Product, error)
	LastID(id string) (Product, error)
	Delete(id string) error
}

func NewRepository(db store.Store) Repository {
	return &FileRepository{db}
}