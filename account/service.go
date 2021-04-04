package account

type Service interface {
	Store(acc *Account) error
	Find(email string) (*Account, error)
	Delete(email string) (*Account, error)
	Verify(acc *Account, password string) (string, error)
}
