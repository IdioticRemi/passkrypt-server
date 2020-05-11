package memory

import (
	"github.com/google/uuid"
	"github.com/kizuru/passkrypt-server/pkg/adding"
	"github.com/kizuru/passkrypt-server/pkg/listing"
)

type Storage struct {
	credentials []Account
}

func (s *Storage) AddAccount(credential adding.Account) error {
	for _, existing := range s.credentials {
		if credential.UserID == existing.UserID && credential.Username == existing.Username && credential.Password == existing.Password {
			return adding.ErrDuplicate
		}
	}

	newCredidential := Account{
		ID:       uuid.New().String(),
		UserID:   credential.UserID,
		Username: credential.Username,
		Password: credential.Password,
		Note:     credential.Note,
	}
	s.credentials = append(s.credentials, newCredidential)

	return nil
}

func (s Storage) GetAccount(id string) (listing.Account, error) {
	var credential listing.Account

	for i := range s.credentials {
		if s.credentials[i].ID == id {
			credential.ID = s.credentials[i].ID
			credential.UserID = s.credentials[i].UserID
			credential.Username = s.credentials[i].Username
			credential.Password = s.credentials[i].Password
			credential.Note = s.credentials[i].Note

			return credential, nil
		}
	}

	return credential, listing.ErrNotFound
}

func (s Storage) GetAllAccounts() []listing.Account {
	var credentials []listing.Account

	for i := range s.credentials {
		credentials = append(credentials, listing.Account{
			ID:       s.credentials[i].ID,
			UserID:   s.credentials[i].UserID,
			Username: s.credentials[i].Username,
			Password: s.credentials[i].Password,
			Note:     s.credentials[i].Note,
		})
	}

	return credentials
}
