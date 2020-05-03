package postgres

import (
	"github.com/BranDebs/Avocado-Backend/account"
)

type postgresRepository struct {
}

func NewPostgresRepository() (account.AccountRepository, error) {
	return &postgresRepository{}, nil
}

func (*postgresRepository) Find(email string) (*account.Account, error) {
	return nil, nil
}

func (*postgresRepository) Store(account *account.Account) error {
	return nil
}

func (*postgresRepository) Delete(email string) (*account.Account, error) {
	return nil, nil
}
