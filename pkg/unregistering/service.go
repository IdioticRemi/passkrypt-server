package unregistering

type Service interface {
	UnregisterUser(...User) error
}

type Repository interface {
	UnregisterUser(User) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UnregisterUser(users ...User) error {
	for _, user := range users {
		if err := s.r.UnregisterUser(user); err != nil {
			return err
		}
	}

	return nil
}
