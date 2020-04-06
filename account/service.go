package account

type AccountService interface {
	Store(*Account) error
	Find(email string) (*Account, error)
	Delete(email string) (*Account, error)
}
