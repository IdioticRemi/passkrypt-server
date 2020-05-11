package registering

type Repository interface {
}

type Service interface {
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}
