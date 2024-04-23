package products

import (
	"errors"

	"fmt"
	"github.com/luuan11/middleProducts/pkg/store"
)


type FileRepository struct {
	db store.Store
}

func NewFileRepository(db store.Store) Repository {
	return &FileRepository{
		db: db,
	}
}

func (r *FileRepository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *FileRepository) Store(name, category string, count int, price float64) (Product, error) {
	p := Product{
		Name:     name,
		Category: category,
		Count:    count,
		Price:    price,
	}

	var ps []Product

	// primeiro lemos o arquivo
	r.db.Read(&ps)

	// calculamos qual o próximo ID
	lastIdInserted := len(ps)
	lastIdInserted++
	p.ID = fmt.Sprint(lastIdInserted)

	// inserimos o produto a ser cadastrado no slice de produtos
	ps = append(ps, p)

	// gravamos no arquivo novamente com o novo produto inserido
	if _, err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *FileRepository) Delete(id string) error {
	return nil
}

func (r *FileRepository) Update(pd, produto Product) Product {
	return Product{}
}
func (r *FileRepository) UpdateName(id string, name string) (Product, error) {
	return Product{}, nil
}

func (r *FileRepository) LastID(id string) (Product, error) {
	if products, err := r.GetAll(); err != nil {
		return Product{}, err
	} else {
		for _, p := range products {
			if p.ID == id {
				return p, nil
			}
		}
		return Product{}, errors.New("Nao foi possivel achar o produto com o Id: " + id)
	}
}
