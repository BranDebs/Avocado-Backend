package model

// Status represents a user's status for their task.
type Status uint

const (
	StatusUnknown Status = iota
	StatusTodo
	StatusInProgress
	StatusCompleted
)

// String returns `Status` enum type into a readable string.
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

// Valid returns true if `Status` is a valid enum type.
func (t *Status) Valid() bool {
	return t.String() != "unknown"
}
