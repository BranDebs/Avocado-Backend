package task

import (
	"errors"

	"github.com/BranDebs/Avocado-Backend/task/model"
)

var (
	ErrNilTask     = errors.New("task: task cannot be nil")
	ErrInvalidTask = errors.New("task: task is not valid")
)

type Repository interface {
	Store(t *model.Task) error
	Find(userID uint) ([]*model.Task, error)
	Update(t *model.Task) (*model.Task, error)
	Delete(ids ...uint) ([]*model.Task, error)
}
