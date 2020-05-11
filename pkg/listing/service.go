package listing

type Repository interface {
	GetAccount(string) (Account, error)
	GetAllAccounts() []Account
}

type Service interface {
	GetAccount(string) (Account, error)
	GetAccounts() []Account
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAccounts() []Account {
	return s.r.GetAllAccounts()
}

func (s *service) GetAccount(id string) (Account, error) {
	return s.r.GetAccount(id)
}
