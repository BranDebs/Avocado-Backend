package model

type Status uint

const (
	StatusUnknown Status = iota
	StatusTodo
	StatusInProgress
	StatusCompleted
)

func (s *Status) String() string {
	switch *s {
	case StatusTodo:
		return "todo"
	case StatusInProgress:
		return "in_progress"
	case StatusCompleted:
		return "completed"
	}
	return "unknown"
}

func (t *Status) Valid() bool {
	return t.String() != "unknown"
}
