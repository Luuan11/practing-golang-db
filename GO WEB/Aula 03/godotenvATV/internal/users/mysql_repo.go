package users

import "database/sql"

type MySqlRepository struct {
	db *sql.DB
}

func (m *MySqlRepository) GetAll() ([]User, error) {
	return nil, nil
}

func (m *MySqlRepository) Store(Nome, Sobrenome, Email string, Age uint64, Status bool, DateCreation string) (User, error) {
	return User{}, nil
}

func (m *MySqlRepository) LastID() (uint64, error) {
	return 0, nil
}