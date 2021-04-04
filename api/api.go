package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/BranDebs/Avocado-Backend/account"
	"github.com/BranDebs/Avocado-Backend/task"
)

const (
	ContentTypeJSON = "application/json"

	DefaultErrDetail  = "Something wrong happened on our end, try again in 30 minutes time."
	JSONBodyErrDetail = "Ensure that request body is a valid JSON object."
)

type Handler interface {
	CreateAccount(http.ResponseWriter, *http.Request)
	LoginAccount(http.ResponseWriter, *http.Request)
	DeleteAccount(http.ResponseWriter, *http.Request)

	CreateTask(http.ResponseWriter, *http.Request)
	FindTasks(http.ResponseWriter, *http.Request)
	UpdateTask(http.ResponseWriter, *http.Request)
	DeleteTask(http.ResponseWriter, *http.Request)

	Ping(http.ResponseWriter, *http.Request)
}

type handler struct {
	acctSvc account.Service
	taskSvc task.Service
}

func NewHandler(acctSvc account.Service, taskSvc task.Service) Handler {
	return &handler{
		acctSvc: acctSvc,
		taskSvc: taskSvc,
	}
}

func (*handler) Ping(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, ContentTypeJSON, http.StatusOK, []byte(`{"message": "Pong!"}`))
}

func setupResponse(w http.ResponseWriter, contentType string, statusCode int, body []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	if _, err := w.Write(body); err != nil {
		log.Println(err)
	}
}

func setupError(w http.ResponseWriter, contentType string, statusCode int, err *Error) {
	resp, _ := json.Marshal(err)
	setupResponse(w, contentType, statusCode, resp)
}
