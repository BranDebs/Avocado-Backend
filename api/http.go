package api

import (
	"log"
	"net/http"

	"github.com/BranDebs/Avocado-Backend/account"
)

const (
	ContentTypeJSON = "application/json"
)

func setupResponse(w http.ResponseWriter, contentType string, statusCode int, body []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	if _, err := w.Write(body); err != nil {
		log.Println(err)
	}
}

type Handler interface {
	PostAccount(http.ResponseWriter, *http.Request)
	GetAccount(http.ResponseWriter, *http.Request)
	DeleteAccount(http.ResponseWriter, *http.Request)

	Ping(http.ResponseWriter, *http.Request)
}

type handler struct {
	acctSvc account.AccountService
}

func NewHandler(acctSvc account.AccountService) Handler {
	return &handler{
		acctSvc: acctSvc,
	}
}

func (*handler) Ping(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, ContentTypeJSON, http.StatusOK, []byte(`{"message": "Pong!"}`))
}

func (h *handler) PostAccount(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, ContentTypeJSON, http.StatusOK, []byte(`{"message":"Post Account"}`))
}

func (h *handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, ContentTypeJSON, http.StatusOK, []byte(`{"message":"Get Account"}`))
}

func (h *handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, ContentTypeJSON, http.StatusOK, []byte(`{"message":"Delete Account"}`))
}
