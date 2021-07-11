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

// NewTaskRepository takes in a valid postgres connection setting and returns a task repository in postgres dialect.
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

// Store takes in a valid task and stores it in postgres.
func (r *taskRepository) Store(t *model.Task) error {
	return nil
}

// Find obtains a task from postgres using a vald user ID.
func (r *taskRepository) Find(userID uint) ([]*model.Task, error) {
	return nil, nil
}

// Update does a partial update using a valid task into postgres.
func (r *taskRepository) Update(t *model.Task) (*model.Task, error) {
	return nil, nil
}

// Delete removes tasks using valid task IDs.
func (r *taskRepository) Delete(ids ...uint) ([]*model.Task, error) {
	return nil, nil
}

func newTaskEntity(t *model.Task) *taskEntity {
	return &taskEntity{
		task: *t,
	}
}
