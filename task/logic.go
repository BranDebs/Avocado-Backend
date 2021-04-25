package task

import (
	"github.com/BranDebs/Avocado-Backend/task/model"
)

type taskService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) Store(t *model.Task) error {
	if err := validateTask(t); err != nil {
		return err
	}

	return nil
}

func (s *taskService) Find(userID uint) ([]*model.Task, error) {
	return nil, nil
}

func (s *taskService) Update(t *model.Task) (*model.Task, error) {
	if err := validateTask(t); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *taskService) Delete(ids ...uint) ([]*model.Task, error) {
	return nil, nil
}

func validateTask(t *model.Task) error {
	if t == nil {
		return ErrNilTask
	}

	if !t.Valid() {
		return ErrInvalidTask
	}

	return nil
}
