package api

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
