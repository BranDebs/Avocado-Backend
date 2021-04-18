package api

import "net/http"

func (h *handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, ContentTypeJSON, http.StatusOK, nil)
}

func (h *handler) FindTasks(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, ContentTypeJSON, http.StatusOK, nil)
}

func (h *handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, ContentTypeJSON, http.StatusOK, nil)
}

func (h *handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, ContentTypeJSON, http.StatusOK, nil)
}
