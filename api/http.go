package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/BranDebs/Avocado-Backend/account"
	"github.com/go-chi/chi"
)

const (
	ContentTypeJSON = "application/json"

	CreateAccountErrMsg = "Failed to create an account."
	LoginAccountErrMsg  = "Failed to login into account."
	DeleteAccountErrMsg = "Failed to delete account."

	DefaultErrDetail  = "Something wrong happened on our end, try again in 30 minutes time."
	JSONBodyErrDetail = "Ensure that request body is a valid JSON object."
)

type Handler interface {
	CreateAccount(http.ResponseWriter, *http.Request)
	LoginAccount(http.ResponseWriter, *http.Request)
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

func (h *handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var bbuf bytes.Buffer

	if _, err := io.Copy(&bbuf, r.Body); err != nil {
		setupError(w, ContentTypeJSON, http.StatusUnprocessableEntity, &Error{
			Message: CreateAccountErrMsg,
			Detail:  JSONBodyErrDetail,
			Cause:   err,
		})
		return
	}
	defer r.Body.Close()

	var acc Account
	if err := json.Unmarshal(bbuf.Bytes(), &acc); err != nil {
		setupError(w, ContentTypeJSON, http.StatusInternalServerError, &Error{
			Message: CreateAccountErrMsg,
			Detail:  DefaultErrDetail,
			Cause:   err,
		})
		return
	}

	if err := h.acctSvc.Store(&account.Account{Email: acc.Email, Password: []byte(acc.Password)}); err != nil {
		if errors.Is(err, account.ErrDuplicateEmail) {
			setupError(w, ContentTypeJSON, http.StatusUnprocessableEntity, &Error{
				Message: CreateAccountErrMsg,
				Detail:  "You already have an account with us.",
				Cause:   err,
			})
			return
		}

		setupError(w, ContentTypeJSON, http.StatusInternalServerError, &Error{
			Message: CreateAccountErrMsg,
			Detail:  DefaultErrDetail,
			Cause:   err,
		})
		return
	}

	acc.Password = ""

	resp, err := json.Marshal(&acc)
	if err != nil {
		setupError(w, ContentTypeJSON, http.StatusInternalServerError, &Error{
			Message: CreateAccountErrMsg,
			Detail:  DefaultErrDetail,
			Cause:   err,
		})
		return
	}

	setupResponse(w, ContentTypeJSON, http.StatusOK, resp)
}

func (h *handler) LoginAccount(w http.ResponseWriter, r *http.Request) {
	var bbuf bytes.Buffer

	if _, err := io.Copy(&bbuf, r.Body); err != nil {
		setupError(w, ContentTypeJSON, http.StatusUnprocessableEntity, &Error{
			Message: LoginAccountErrMsg,
			Detail:  JSONBodyErrDetail,
		})
		return
	}
	defer r.Body.Close()

	var acc Account
	if err := json.Unmarshal(bbuf.Bytes(), &acc); err != nil {
		setupError(w, ContentTypeJSON, http.StatusUnprocessableEntity, &Error{
			Message: LoginAccountErrMsg,
			Detail:  JSONBodyErrDetail,
			Cause:   err,
		})
		return
	}

	foundAcc, err := h.acctSvc.Find(acc.Email)
	if err != nil {
		errCode := http.StatusInternalServerError
		detail := DefaultErrDetail
		if errors.Is(err, account.ErrRecordNotFound) {
			errCode = http.StatusNotFound
			detail = "Email does not exists."
		}
		setupError(w, ContentTypeJSON, errCode, &Error{
			Message: LoginAccountErrMsg,
			Detail:  detail,
			Cause:   err,
		})
		return
	}

	success, err := h.acctSvc.Verify(foundAcc, acc.Password)
	if err != nil {
		setupError(w, ContentTypeJSON, http.StatusInternalServerError, &Error{
			Message: LoginAccountErrMsg,
			Detail:  DefaultErrDetail,
			Cause:   err,
		})
		return
	}

	if !success {
		setupError(w, ContentTypeJSON, http.StatusUnauthorized, &Error{
			Message: LoginAccountErrMsg,
			Detail:  "Invalid login credentials.",
			Cause:   err,
		})
		return
	}

	resp, err := json.Marshal(&acc)
	if err != nil {
		setupError(w, ContentTypeJSON, http.StatusInternalServerError, &Error{
			Message: LoginAccountErrMsg,
			Detail:  DefaultErrDetail,
			Cause:   err,
		})
		return
	}

	setupResponse(w, ContentTypeJSON, http.StatusOK, resp)
}

func (h *handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	if _, err := h.acctSvc.Delete(email); err != nil {
		errCode := http.StatusInternalServerError
		detail := DefaultErrDetail
		if errors.Is(err, account.ErrRecordNotFound) {
			errCode = http.StatusNotFound
			detail = "Email does not exists."
		}
		setupError(w, ContentTypeJSON, errCode, &Error{
			Message: DeleteAccountErrMsg,
			Detail:  detail,
			Cause:   err,
		})
		return
	}

	setupResponse(w, ContentTypeJSON, http.StatusNoContent, nil)
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
	log.Printf("%s: %s", err.Message, err.Cause)
	setupResponse(w, contentType, statusCode, resp)
}
