package postgres

import (
	"errors"
	"fmt"

	"github.com/BranDebs/Avocado-Backend/account"
	"github.com/BranDebs/Avocado-Backend/secrets"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

type ConnSettings struct {
	Host         string `mapstructure:"host"`
	Port         int64  `mapstructure:"port"`
	User         string `mapstructure:"user"`
	PasswordFile string `mapstructure:"password_file"`
	DBName       string `mapstructure:"db_name"`
}

func (c ConnSettings) String() string {
	password, _ := secrets.SingleLineKey(c.PasswordFile)

	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		c.Host, c.Port, c.User, c.DBName, password)
}

func newGormDB(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
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

	if err := db.AutoMigrate(&account.Account{}); err != nil {
		return nil, err
	}

	repo := postgresRepository{
		db: db,
	}
	return &repo, nil
}

func (repo *postgresRepository) Find(email string) (*account.Account, error) {
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

func (repo *postgresRepository) Store(acc *account.Account) error {
	if _, err := repo.Find(acc.Email); err == nil {
		return account.ErrDuplicateEmail
	}

	if err := repo.db.Create(acc).Error; err != nil {
		return err
	}
	return nil
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
