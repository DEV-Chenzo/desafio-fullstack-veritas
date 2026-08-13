package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"veritas-kanban/internal/task"
)

type TaskHandler struct{ service *task.Service }

func NewTaskHandler(service *task.Service) *TaskHandler { return &TaskHandler{service: service} }

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.List(r.Context())
	if err != nil {
		InternalError(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, tasks)
}
func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	input, ok := decodeInput(w, r)
	if !ok {
		return
	}
	item, err := h.service.Create(r.Context(), input)
	if err != nil {
		BadRequest(w, err.Error())
		return
	}
	WriteJSON(w, http.StatusCreated, item)
}
func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, ok := taskID(w, r)
	if !ok {
		return
	}
	item, err := h.service.Get(r.Context(), id)
	h.respondTask(w, item, err, http.StatusOK)
}
func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, ok := taskID(w, r)
	if !ok {
		return
	}
	input, ok := decodeInput(w, r)
	if !ok {
		return
	}
	item, err := h.service.Update(r.Context(), id, input)
	h.respondTask(w, item, err, http.StatusOK)
}
func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, ok := taskID(w, r)
	if !ok {
		return
	}
	err := h.service.Delete(r.Context(), id)
	if errors.Is(err, task.ErrNotFound) {
		NotFound(w)
		return
	}
	if err != nil {
		InternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (h *TaskHandler) respondTask(w http.ResponseWriter, item task.Task, err error, status int) {
	if errors.Is(err, task.ErrNotFound) {
		NotFound(w)
		return
	}
	if err != nil {
		InternalError(w, err)
		return
	}
	WriteJSON(w, status, item)
}

func decodeInput(w http.ResponseWriter, r *http.Request) (task.Input, bool) {
	defer r.Body.Close()
	var input task.Input
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		BadRequest(w, "JSON inválido")
		return input, false
	}
	return input, true
}
func taskID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		BadRequest(w, "ID inválido")
		return 0, false
	}
	return id, true
}
