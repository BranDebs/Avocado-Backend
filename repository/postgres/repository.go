package postgres

import (
	"fmt"

	"github.com/BranDebs/Avocado-Backend/secrets"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

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
	if db != nil {
		return db, nil
	}

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
