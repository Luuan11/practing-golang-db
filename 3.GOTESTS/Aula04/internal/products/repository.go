package products

import (
	"github.com/luuan11/middleProducts/pkg/store"
	"github.com/luuan11/middleProducts/internal/entities"
)

type Repository interface {
	GetAll() ([]entities.Product, error)
	GetById(id int) (entities.Product, error)
	Store(name, category string, count int, price float64) (entities.Product, error)
	Update(id uint64, name, category string, count int, price float64) (entities.Product, error)
	UpdateName(id uint64, name string) (entities.Product, error)
	LastID() (uint64, error)
	Delete(id uint64) error
}

func NewRepository(db store.Store) Repository {
	return &FileRepository{db}
}