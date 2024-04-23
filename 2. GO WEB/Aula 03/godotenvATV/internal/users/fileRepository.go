package users

import (
	"github.com/luuan11/godotENV/pkg/store"
)

type FileRepository struct {
	db storeUser.Store
}

func NewFileRepository(db storeUser.Store) Repository {
	return &FileRepository{
		db: db,
	}
}

func (r *FileRepository) GetAll() ([]User, error) {
	var psUser []User
	r.db.Read(&psUser)
	return psUser, nil
}

func (r *FileRepository) Store(Nome, Sobrenome, Email string, Age uint64, Status bool, DateCreation string) (User, error) {
	u := User{
		Name: Nome,
		Sobrenome: Sobrenome,
		Email: Email,
		Age: Age,
		Status: Status,
		DateCreation: DateCreation,
	}

	var psUser []User

	// primeiro lemos o arquivo
	r.db.Read(&psUser)

	// calculamos qual o próximo ID
	lastIdInserted := len(psUser)
	lastIdInserted++
	u.ID = uint64(lastIdInserted)

	// inserimos o produto a ser cadastrado no slice de produtos
	psUser = append(psUser, u)

	// gravamos no arquivo novamente com o novo produto inserido
	err := r.db.Write(psUser)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (r *FileRepository) Delete(id uint64) error {
	return nil
}

func (r *FileRepository) Update(id uint64, Nome, Sobrenome, Email string, Age uint64, Status bool, DateCreation string) (User, error) {
	return User{}, nil
}
func (r *FileRepository) UpdateName(id uint64, name string) (User, error) {
	return User{}, nil
}

func (r *FileRepository) lastID() (uint64, error) {
	var psUser []User
	if err := r.db.Read(&psUser); err != nil {
		return 0, err
	}

	if len(psUser) == 0 {
		return 0, nil
	}

	return psUser[len(psUser)-1].ID, nil

}