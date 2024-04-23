package users

import (
	 "github.com/luuan11/godotENV/pkg/store"
)

type User struct{
	ID 				uint64		`json:"id"`
	Name 			string   	`json:"name"`
	Sobrenome		string		`json:"lastname"`
	Email			string		`json:"email"`
	Age 			uint64	    `json:"age"`
	Status 			bool 		`json:"status"`
	DateCreation 	string 		`json:"dateCreation"`
}

//estrutura do repository
type Repository interface{
	GetAll() ([]User, error)
	Store(Nome, Sobrenome, Email string, Age uint64, Status bool, DateCreation string) (User, error)
	lastID()(uint64, error)
	Update(Id uint64, Nome, Sobrenome, Email string, Age uint64, Status bool, DateCreation string) (User, error)
	UpdateName(id uint64, Nome string) (User, error)
	Delete(id uint64) error
}

func NewRepository(db storeUser.Store) Repository{
	return &FileRepository{db}
}

