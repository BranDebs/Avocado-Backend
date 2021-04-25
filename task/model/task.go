package model

type Task struct {
	UserID      uint
	Status      Status
	Description string
}

func (t *Task) Valid() bool {
	if t.UserID == 0 {
		return false
	}

	if !t.Status.Valid() {
		return false
	}

	return true
}
