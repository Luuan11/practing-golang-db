package users

//estrutura Service
type Service interface {
	GetAll() ([]User, error)
	Store(Nome, Sobrenome, Email string, Age uint64, Status bool, DateCreation string) (User, error)
	UpdateName(id uint64, Nome string) (User, error)
	Update(id uint64, Nome, Sobrenome, Email string, Age uint64, Status bool, DateCreation string) (User, error)
	Delete(id uint64) error
}

type service struct{
	repository Repository
}

//metodo GetAll
func (s *service)GetAll() ([]User, error){
	psUser, err:= s.repository.GetAll()
	if err != nil{
		return nil,err
	}
	return psUser, nil
}

//Metodo store
func (s *service) Store(Nome, Sobrenome, Email string, Age uint64, Status bool, DateCreation string)(User, error) {
	
	user ,err := s.repository.Store(Nome, Sobrenome, Email, Age, Status, DateCreation)
	if err != nil{
		return User{}, err
	}

	return user, nil
}


func (s *service) Update(id uint64, Nome, Sobrenome, Email string, Age uint64, Status bool, DateCreation string) (User, error) {
	return s.repository.Update(id, Nome, Sobrenome, Email, Age, Status, DateCreation)
}

func (s *service) UpdateName(id uint64, Nome string) (User, error) {
	return s.repository.UpdateName(id, Nome)
}

func (s *service) Delete(id uint64) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
