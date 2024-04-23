package users

type User struct{
	ID 				int		    `json:"id"`
	Name 			string   	`json:"name"`
	Sobrenome		string		`json:"lastname"`
	Email			string		`json:"email"`
	Age 			int	    	`json:"age"`
	Status 			bool 		`json:"status"`
	DateCreation 	string 		`json:"dateCreation"`
}

var psUser []User
var lastID int

//estrutura do repository
type Repository interface{
	GetAll() ([]User, error)
	Store(Id int, Nome, Sobrenome, Email string, Age int, Status bool, DateCreation string) (User, error)
	lastID()(int, error)
}

type repository struct{}

func NewRepository() Repository{
	return &repository{}
}

//metodos do repository
func (r *repository) GetAll() ([]User, error){
	return psUser , nil
}

func (r *repository)Store(Id int, Nome, Sobrenome, Email string, Age int, Status bool, DateCreation string) (User, error) {
	p := User{Id, Nome, Sobrenome, Email, Age, Status, DateCreation}
	psUser = append(psUser, p)
	lastID = p.ID
	return p,nil
}

func (r *repository)lastID()(int, error){
	return lastID,nil
}