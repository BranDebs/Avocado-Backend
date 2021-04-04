package account

type Repository interface {
	Find(email string) (*Account, error)
	Store(*Account) error
	Delete(email string) (*Account, error)
}
