package account

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
	return s.accountRepo.Store(account)
}

func (s *accountService) Delete(email string) (*Account, error) {
	return s.accountRepo.Delete(email)
}
