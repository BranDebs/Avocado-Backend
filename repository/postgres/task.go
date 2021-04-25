package postgres

import (
	"github.com/BranDebs/Avocado-Backend/task"
	"github.com/BranDebs/Avocado-Backend/task/model"

	"gorm.io/gorm"
)

type taskEntity struct {
	gorm.Model
	task model.Task `gorm:"embedded"`
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(settings ConnSettings) (task.Repository, error) {
	db, err := newGormDB(settings.String())
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&taskEntity{}); err != nil {
		return nil, err
	}

	repo := taskRepository{
		db: db,
	}
	return &repo, nil
}

func (r *taskRepository) Store(t *model.Task) error {
	return nil
}

func (r *taskRepository) Find(userID uint) ([]*model.Task, error) {
	return nil, nil
}

func (r *taskRepository) Update(t *model.Task) (*model.Task, error) {
	return nil, nil
}

func (r *taskRepository) Delete(ids ...uint) ([]*model.Task, error) {
	return nil, nil
}

func newTaskEntity(t *model.Task) *taskEntity {
	return &taskEntity{
		task: *t,
	}
}
