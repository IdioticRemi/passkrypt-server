package memory

import (
	"github.com/google/uuid"
	"github.com/kizuru/passkrypt-server/pkg/registering"
)

type User struct {
	ID       string
	Email    string
	Username string
	Password string
}

func (s *Storage) AddUser(user registering.User) error {
	for _, existing := range s.users {
		if user.Email == existing.Email && user.Password == existing.Password {
			return registering.ErrUserDuplicate
		}
	}

	newUser := User{
		ID:       uuid.New().String(),
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}
	s.users = append(s.users, newUser)

	return nil
}
