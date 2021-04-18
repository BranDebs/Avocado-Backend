package postgres

import (
	"github.com/BranDebs/Avocado-Backend/task"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(settings ConnSettings) (task.Repository, error) {
	db, err := newGormDB(settings.String())
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&task.Task{}); err != nil {
		return nil, err
	}

	repo := taskRepository{
		db: db,
	}
	return &repo, nil
}

func (r *taskRepository) Store(t *task.Task) error {
	return nil
}

func (r *taskRepository) Find(userID uint) ([]*task.Task, error) {
	return nil, nil
}

func (r *taskRepository) Update(t *task.Task) (*task.Task, error) {
	return nil, nil
}

func (r *taskRepository) Delete(ids ...uint) ([]*task.Task, error) {
	return nil, nil
}
