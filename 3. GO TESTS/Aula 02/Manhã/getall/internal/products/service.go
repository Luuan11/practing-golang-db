package products

import "errors"

type Service interface {
	GetAll() ([]Product, error)
	Store(name, category string, count int, price float64) (Product, error)
	Update(id string, name, category string, count int, price float64) (Product, error)
	UpdateName(id string, name string) (Product, error)
	Delete(id string) error
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	produtos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return produtos, nil

}

func (s *service) Store(name, category string, count int, price float64) (Product, error) {
	product, err := s.repository.Store(name, category, count, price)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Update(id string, name, category string, count int, price float64) (Product, error) {
	if prod, err := s.repository.LastID(string(id)); err != nil {
		return Product{}, errors.New("Product not found")
	} else {
		newProd := Product{
			ID: 		prod.ID,
			Name:     	name,
			Category: 	category,
			Count: 		count,
			Price: 		float64(price),
		}
		pd := s.repository.Update(prod, newProd)
		return pd, nil
	}
}

func (s *service) UpdateName(id string, name string) (Product, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id string) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}