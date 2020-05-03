package postgres

import (
	"fmt"

	"github.com/BranDebs/Avocado-Backend/account"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

func NewPostgresRepository(settings ConnSettings) (account.AccountRepository, error) {
	db, err := newGormDB(settings.String())
	if err != nil {
		return nil, err
	}
	repo := postgresRepository{
		db: db,
	}
	return &repo, nil
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
