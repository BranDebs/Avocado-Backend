package postgres

import (
	"errors"

	"github.com/BranDebs/Avocado-Backend/account"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(settings ConnSettings) (account.Repository, error) {
	db, err := newGormDB(settings.String())
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&account.Account{}); err != nil {
		return nil, err
	}

	repo := accountRepository{
		db: db,
	}
	return &repo, nil
}

func (repo *accountRepository) Find(email string) (*account.Account, error) {
	var acc account.Account
	res := repo.db.Where(account.Account{Email: email}).First(&acc)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, account.ErrRecordNotFound
		}
		return nil, res.Error
	}
	return &acc, nil
}

func (repo *accountRepository) Store(acc *account.Account) error {
	if _, err := repo.Find(acc.Email); err == nil {
		return account.ErrDuplicateEmail
	}

	if err := repo.db.Create(acc).Error; err != nil {
		return err
	}
	return nil
}

func (repo *accountRepository) Delete(email string) (*account.Account, error) {
	var acc *account.Account
	acc, err := repo.Find(email)
	if err != nil {
		return nil, err
	}

	res := repo.db.Delete(acc)
	if res.Error != nil {
		return nil, res.Error
	}
	return acc, nil
}
