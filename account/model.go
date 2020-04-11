package account

import (
	"time"
)

type Account struct {
	Email     string     `json:"email"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
}
