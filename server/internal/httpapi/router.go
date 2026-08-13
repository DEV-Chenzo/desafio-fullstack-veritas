package httpapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"veritas-kanban/internal/task"
)

func NewRouter(service *task.Service, db *pgxpool.Pool, corsOrigin string) http.Handler {
	handler := NewTaskHandler(service)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(r.Context()); err != nil {
			InternalError(w, err)
			return
		}
		WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /api/tasks", handler.List)
	mux.HandleFunc("POST /api/tasks", handler.Create)
	mux.HandleFunc("GET /api/tasks/{id}", handler.Get)
	mux.HandleFunc("PUT /api/tasks/{id}", handler.Update)
	mux.HandleFunc("DELETE /api/tasks/{id}", handler.Delete)
	return CORS(corsOrigin, Recover(mux))
}

func CORS(origin string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if value := recover(); value != nil {
				InternalError(w, fmt.Errorf("panic: %v", value))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
func InternalError(w http.ResponseWriter, err error) {
	log.Printf("erro interno: %v", err)
	WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Erro interno do servidor"})
}
