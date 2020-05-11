package listing

type Repository interface {
	GetAccount(string) (Account, error)
	GetAccounts() []Account
	GetUser(string) (User, error)
	GetUsers() []User
}

type Service interface {
	GetAccount(string) (Account, error)
	GetAccounts() []Account
	GetUser(string) (User, error)
	GetUsers() []User
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAccounts() []Account {
	return s.r.GetAccounts()
}

func (s *service) GetAccount(id string) (Account, error) {
	return s.r.GetAccount(id)
}

func (s *service) GetUsers() []User {
	return s.r.GetUsers()
}

func (s *service) GetUser(id string) (User, error) {
	return s.r.GetUser(id)
}
