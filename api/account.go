package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/BranDebs/Avocado-Backend/account"
	"github.com/go-chi/chi"
)

const (
	CreateAccountErrMsg = "Failed to create an account."
	LoginAccountErrMsg  = "Failed to login into account."
	DeleteAccountErrMsg = "Failed to delete account."
)

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

	signedJWT, err := h.acctSvc.Verify(foundAcc, acc.Password)
	if err != nil {
		detail := DefaultErrDetail
		if errors.Is(err, account.ErrNotVerified) {
			detail = "Invalid login credentials."
		}

		setupError(w, ContentTypeJSON, http.StatusInternalServerError, &Error{
			Message: LoginAccountErrMsg,
			Detail:  detail,
			Cause:   err,
		})
		return
	}

	jwtRes := &JWTResponse{
		Token: signedJWT,
	}

	resp, err := json.Marshal(jwtRes)
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
