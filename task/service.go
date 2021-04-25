package task

import (
	"github.com/BranDebs/Avocado-Backend/task/model"
)

type Service interface {
	Store(t *model.Task) error
	Find(userID uint) ([]*model.Task, error)
	Update(t *model.Task) (*model.Task, error)
	Delete(ids ...uint) ([]*model.Task, error)
}
