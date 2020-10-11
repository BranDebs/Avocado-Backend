package account

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 64
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
	ErrDuplicateEmail = errors.New("email already exists")
)

type accountService struct {
	accountRepo AccountRepository
}

func NewAccountService(repo AccountRepository) AccountService {
	return &accountService{
		accountRepo: repo,
	}
}

func (s *accountService) Find(email string) (*Account, error) {
	return s.accountRepo.Find(email)
}

func (s *accountService) Store(account *Account) error {
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return fmt.Errorf("AccountService.Store: generating salt: %w", err)
	}

	hash, err := scrypt.Key([]byte(account.Password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		return fmt.Errorf("AccountService.Store: generating password hash: %w", err)
	}

	account.Password = hash
	account.PasswordSalt = salt

	return s.accountRepo.Store(account)
}

func (s *accountService) Delete(email string) (*Account, error) {
	return s.accountRepo.Delete(email)
}

func (s *accountService) Verify(acc *Account, password string) (bool, error) {
	hash, err := scrypt.Key([]byte(password), acc.PasswordSalt, 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		return false, fmt.Errorf("AccountService.Verify: generating password hash: %w", err)
	}

	return bytes.Equal(hash, acc.Password), nil
}
