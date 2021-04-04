package api

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type JWTResponse struct {
	Token string `json:"token"`
}

type Task struct {
	Description string `json:"description"`
}
