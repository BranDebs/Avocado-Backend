package account

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"github.com/BranDebs/Avocado-Backend/internal/jwt"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 64

	n = 1 << 14
	r = 8
	p = 1
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
	ErrDuplicateEmail = errors.New("email already exists")
	ErrNotVerified    = errors.New("credentials cannot be verified")
)

type accountService struct {
	accountRepo Repository

	jwtSettings *jwt.Settings
}

func NewService(repo Repository, jwtSettings *jwt.Settings) Service {
	jwtSettings.Init()

	return &accountService{
		accountRepo: repo,
		jwtSettings: jwtSettings,
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

	hash, err := scrypt.Key([]byte(account.Password), salt, n, r, p, PW_HASH_BYTES)
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

func (s *accountService) Verify(acc *Account, password string) (string, error) {
	hash, err := scrypt.Key([]byte(password), acc.PasswordSalt, n, r, p, PW_HASH_BYTES)
	if err != nil {
		return "", fmt.Errorf("AccountService.Verify: generating password hash: %w", err)
	}

	if !bytes.Equal(hash, acc.Password) {
		return "", ErrNotVerified
	}

	jwt := jwt.New(acc.Email, s.jwtSettings.TTL)

	return jwt.Token(s.jwtSettings.SigningKey), nil
}
