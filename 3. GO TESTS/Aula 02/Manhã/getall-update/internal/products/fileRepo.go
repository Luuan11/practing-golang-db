package products

import (
	"errors"
	"fmt"

	"github.com/luuan11/middleProducts/internal/entities"
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

func (r *FileRepository) GetAll() ([]entities.Product, error) {
	var ps []entities.Product
	err := r.db.Read(&ps)
	if err != nil {
		return []entities.Product{}, errors.New("error for GetAll")
	}
	return ps, nil
}

func (r *FileRepository) Store(name, category string, count int, price float64) (entities.Product, error) {
	p := entities.Product{
		Name:     name,
		Category: category,
		Count:    count,
		Price:    price,
	}

	var ps []entities.Product

	// primeiro lemos o arquivo
	err := r.db.Read(&ps)
	if err != nil {
		return entities.Product{}, errors.New("error for Storage")
	}
	// calculamos qual o próximo ID
	lastIdInserted := len(ps)
	lastIdInserted++
	p.ID = uint64(lastIdInserted)

	// inserimos o produto a ser cadastrado no slice de produtos
	ps = append(ps, p)

	// gravamos no arquivo novamente com o novo produto inserido
	err = r.db.Write(ps)
	if err != nil {
		return entities.Product{}, err
	}
	return p, nil
}

func (r *FileRepository) Delete(id uint64) error {
	return nil
}

func (r *FileRepository) Update(id uint64, name, productType string, count int, price float64) (entities.Product, error) {
	return entities.Product{}, nil
}
func (r *FileRepository) UpdateName(id uint64, name string) (entities.Product, error) {
	var u entities.Product
	updated := false

	var pd []entities.Product
	if err := r.db.Read(&pd); err != nil {
		return entities.Product{}, err
	}
	for i := range pd {
		if pd[i].ID == id {
			pd[i].Name = name
			updated = true
			u = pd[i]
		}
	}
	if !updated {
		return entities.Product{}, fmt.Errorf("produto nome %d não encontrado", id)
	}

	if err := r.db.Write(pd); err != nil {
		return entities.Product{}, err
	}
	return u, nil
}

func (r *FileRepository) LastID() (uint64, error) {
	var ps []entities.Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].ID, nil

}
