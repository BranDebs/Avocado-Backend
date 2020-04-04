package account

import (
	"time"
)

type Account struct {
	Email     string
	Name      string
	CreatedAt *time.Time
}
