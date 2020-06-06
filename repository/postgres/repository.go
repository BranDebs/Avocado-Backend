package postgres

import (
	"errors"
	"fmt"

	"github.com/BranDebs/Avocado-Backend/account"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ErrDupAccount error = errors.New("duplicate account error")
)

type postgresRepository struct {
	db *gorm.DB
}

type ConnSettings struct {
	Host     string
	Port     int64
	User     string
	DBName   string
	Password string
}

func (c ConnSettings) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		c.Host, c.Port, c.User, c.DBName, c.Password)
}

func newGormDB(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewRepository(settings ConnSettings) (account.AccountRepository, error) {
	db, err := newGormDB(settings.String())
	if err != nil {
		return nil, err
	}
	repo := postgresRepository{
		db: db,
	}
	return &repo, nil
}

func (repo *postgresRepository) Find(email string) (*account.Account, error) {
	var acc *account.Account
	res := repo.db.Where(account.Account{Email: email}).First(acc)
	if res.Error != nil {
		return nil, res.Error
	}
	return acc, nil
}

func (repo *postgresRepository) Store(account *account.Account) error {
	if repo.db.NewRecord(account) {
		// can store
		repo.db.Create(account)
		return nil
	}
	return ErrDupAccount
}

func (repo *postgresRepository) Delete(email string) (*account.Account, error) {
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
