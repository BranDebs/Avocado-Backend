package model

// Task represents a user's task.
type Task struct {
	UserID      uint
	Status      Status
	Description string
}

// Valid return true if a `Task` contains:
// 	* UserID
//	* Status
func (t *Task) Valid() bool {
	if t.UserID == 0 {
		return false
	}

	if !t.Status.Valid() {
		return false
	}

	return true
}
