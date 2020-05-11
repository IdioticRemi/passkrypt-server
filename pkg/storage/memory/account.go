package memory

import (
	"github.com/google/uuid"
	"github.com/kizuru/passkrypt-server/pkg/adding"
	"github.com/kizuru/passkrypt-server/pkg/listing"
)

type Account struct {
	ID       string
	Name     string
	UserID   string
	Username string
	Password string
	Note     string
}

func (s *Storage) AddAccount(account adding.Account) error {
	for _, existing := range s.accounts {
		if account.UserID == existing.UserID && account.Username == existing.Username && account.Password == existing.Password {
			return adding.ErrAccountDuplicate
		}
	}

	newAccount := Account{
		ID:       uuid.New().String(),
		UserID:   account.UserID,
		Username: account.Username,
		Password: account.Password,
		Note:     account.Note,
	}
	s.accounts = append(s.accounts, newAccount)

	return nil
}

func (s Storage) GetAccount(id string) (listing.Account, error) {
	var account listing.Account

	for i := range s.accounts {
		if s.accounts[i].ID == id {
			account.ID = s.accounts[i].ID
			account.UserID = s.accounts[i].UserID
			account.Username = s.accounts[i].Username
			account.Password = s.accounts[i].Password
			account.Note = s.accounts[i].Note

			return account, nil
		}
	}

	return account, listing.ErrAccountNotFound
}

func (s Storage) GetAccounts() []listing.Account {
	var accounts []listing.Account

	for i := range s.accounts {
		accounts = append(accounts, listing.Account{
			ID:       s.accounts[i].ID,
			UserID:   s.accounts[i].UserID,
			Username: s.accounts[i].Username,
			Password: s.accounts[i].Password,
			Note:     s.accounts[i].Note,
		})
	}

	return accounts
}
