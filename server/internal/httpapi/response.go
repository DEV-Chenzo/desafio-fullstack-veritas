package httpapi

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
func BadRequest(w http.ResponseWriter, message string) {
	WriteJSON(w, http.StatusBadRequest, map[string]string{"error": message})
}
func NotFound(w http.ResponseWriter) {
	WriteJSON(w, http.StatusNotFound, map[string]string{"error": "Tarefa não encontrada"})
}
