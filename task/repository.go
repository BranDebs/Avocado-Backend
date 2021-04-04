package task

type Repository interface {
	Store(t *Task) error
	Find(userID uint) ([]*Task, error)
	Update(t *Task) (*Task, error)
	Delete(ids ...uint) ([]*Task, error)
}
