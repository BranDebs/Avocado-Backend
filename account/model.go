package account

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email        string
	Password     []byte
	PasswordSalt []byte
}

func (Account) TableName() string {
	return "account"
}
