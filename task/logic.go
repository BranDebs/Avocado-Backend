package task

type taskService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) Store(t *Task) error {
	return nil
}

func (s *taskService) Find(userID uint) ([]*Task, error) {
	return nil, nil
}

func (s *taskService) Update(t *Task) (*Task, error) {
	return nil, nil
}

func (s *taskService) Delete(ids ...uint) ([]*Task, error) {
	return nil, nil
}
