package users

//estrutura Service
type Service interface {
	GetAll() ([]User, error)
	Store(Nome, Sobrenome, Email string, Age, Status int, DateCreation string) (User, error)
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
func (s *service) Store(Nome, Sobrenome, Email string, Age, Status int, DateCreation string)(User, error){
	lastID, err := s.repository.lastID()
	if err != nil {
		return User{}, err
	}

	lastID++

	user ,err := s.repository.Store(lastID, Nome, Sobrenome, Email, Age, Status, DateCreation)
	if err != nil{
		return User{}, err
	}

	return user, nil
}

func NewService(r Repository) Service{
	return &service{
		repository: r,
	}
}