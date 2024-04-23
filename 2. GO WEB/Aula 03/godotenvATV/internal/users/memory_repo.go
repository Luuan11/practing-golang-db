package users

import (
	"fmt"
)

var psUser []User
var lastID uint64 = 0

type MemoryRepository struct {
}

//metodos do repository
func (r *MemoryRepository) GetAll() ([]User, error){
	return psUser , nil
}

func (r *MemoryRepository)Store(Nome, Sobrenome, Email string, Age uint64, Status bool, DateCreation string) (User, error) {
	lastID++
	u := User{
		ID: lastID,
		Name: Nome,
		Sobrenome: Sobrenome,
		Email: Email,
		Age: Age,
		Status: Status,
		DateCreation: DateCreation,
	}

	psUser = append(psUser, u)
	return u,nil
}

func (r *MemoryRepository) Update(Id uint64, Nome, Sobrenome, Email string, Age uint64, Status bool, DateCreation string) (User, error) {
	u := User{
		Name: Nome,
		Sobrenome: Sobrenome,
		Email: Email,
		Age: Age,
		Status: Status,
		DateCreation: DateCreation,

	}
	updated := false
	for i := range psUser {
		if psUser[i].ID == Id {
			u.ID = Id
			psUser[i] = u
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("product %d not found", Id)
	}
	return u, nil
}

func (r *MemoryRepository) UpdateName(id uint64, name string) (User, error) {
	var u User
	updated := false
	for i := range psUser {
		if psUser[i].ID == id {
			psUser[i].Name = name
			updated = true
			u = psUser[i]
		}
	}
	if !updated {
		return User{}, fmt.Errorf("product %d not found", id)
	}

	return u, nil
}

func (r *MemoryRepository) Delete(id uint64) error {
	deleted := false
	var index int
	for i := range psUser {
		if psUser[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("product %d not found", id)
	}

	psUser = append(psUser[:index], psUser[index+1:]...)
	return nil
}

func (r *MemoryRepository)lastID()(uint64, error){
	return lastID,nil
}