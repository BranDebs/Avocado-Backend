package api

type Error struct {
	Message string      `json:"message"`
	Detail  interface{} `json:"detail,omitempty"`
	Cause   error       `json:"-"`
}
