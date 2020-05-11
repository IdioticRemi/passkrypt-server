package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kizuru/passkrypt-server/pkg/listing"
	"github.com/kizuru/passkrypt-server/pkg/registering"
	"github.com/kizuru/passkrypt-server/pkg/unregistering"
)

type User struct {
	ID       string
	Email    string
	Username string
	Password string
}

func (s *Storage) RegisterUser(user registering.User) error {
	for _, existing := range s.users {
		if user.Email == existing.Email {
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

func (s *Storage) UnregisterUser(user unregistering.User) error {
	index := s.findUserIndex(user.ID)
	if index < 0 {
		return unregistering.ErrUserNotFound
	}

	s.users = append(s.users[:index], s.users[+1:]...)

	fmt.Println(s.users)

	return nil
}

func (s *Storage) findUserIndex(id string) int {
	for i, user := range s.users {
		if user.ID == id {
			return i
		}
	}

	return -1
}

func (s Storage) GetUser(id string) (listing.User, error) {
	var user listing.User

	for i := range s.users {
		if s.users[i].ID == id {
			user.ID = s.users[i].ID
			user.Email = s.users[i].Email
			user.Username = s.users[i].Username

			return user, nil
		}
	}

	return user, listing.ErrUserNotFound
}

func (s Storage) GetUsers() []listing.User {
	var users []listing.User

	for i := range s.users {
		users = append(users, listing.User{
			ID:       s.users[i].ID,
			Email:    s.users[i].Email,
			Username: s.users[i].Username,
		})
	}

	return users
}
