package task

import (
	"gorm.io/gorm"
)

type Status uint

func (s *Status) String() string {
	switch *s {
	case 1:
		return "todo"
	case 2:
		return "in_progress"
	case 3:
		return "completed"
	}
	return "unknown"
}

type Task struct {
	gorm.Model
	UserID      uint
	Status      Status
	Description string
}

func (Task) TableName() string {
	return "task"
}
