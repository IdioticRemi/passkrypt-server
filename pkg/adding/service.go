package adding

type Repository interface {
	AddAccount(Account) error
}

type Service interface {
	AddAccount(...Account) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddAccount(accounts ...Account) error {
	for _, account := range accounts {
		if err := s.r.AddAccount(account); err != nil {
			return err
		}
	}

	return nil
}
