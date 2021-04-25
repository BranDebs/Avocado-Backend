package task

import (
	"github.com/BranDebs/Avocado-Backend/task/model"
)

type taskService struct {
	repo Repository
}

// NewService returns a service in the task domain to the caller.
func NewService(repo Repository) Service {
	return &taskService{
		repo: repo,
	}
}

// Store stores a valid task for a user based on user ID.
func (s *taskService) Store(t *model.Task) error {
	if err := validateTask(t); err != nil {
		return err
	}

	return nil
}

// Find obtains a task using the userID.
func (s *taskService) Find(userID uint) ([]*model.Task, error) {
	return nil, nil
}

// Update performs a partial update to a valid task of a user.
func (s *taskService) Update(t *model.Task) (*model.Task, error) {
	if err := validateTask(t); err != nil {
		return nil, err
	}

	return nil, nil
}

// Delete removes tasks based on task IDs.
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
