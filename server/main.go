package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Task representa uma tarefa individual no kanban
type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Desc      string    `json:"description"`
	Status    string    `json:"status"` // "todo", "doing", "done"
	CreatedAt time.Time `json:"createdAt"`
}

// TaskStore gerencia as tarefas em memória
type TaskStore struct {
	tasks map[string]Task
	mu    sync.RWMutex
}

// Instância global do store
var store = &TaskStore{
	tasks: make(map[string]Task),
}

// ============== HANDLERS ==============

// GetTasks retorna todas as tarefas
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	store.mu.RLock()
	defer store.mu.RUnlock()
	
	tasks := make([]Task, 0, len(store.tasks))
	for _, task := range store.tasks {
		tasks = append(tasks, task)
	}
	
	json.NewEncoder(w).Encode(tasks)
}

// CreateTask cria uma nova tarefa
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// Validação básica
	if task.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	
	if task.Status == "" {
		task.Status = "todo"
	}
	
	// Gerar ID único
	task.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	task.CreatedAt = time.Now()
	
	store.mu.Lock()
	store.tasks[task.ID] = task
	store.mu.Unlock()
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// UpdateTask atualiza uma tarefa existente
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	id := vars["id"]
	
	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	store.mu.Lock()
	defer store.mu.Unlock()
	
	task, exists := store.tasks[id]
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	
	// Atualizar campos
	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}
	if updatedTask.Desc != "" {
		task.Desc = updatedTask.Desc
	}
	if updatedTask.Status != "" {
		task.Status = updatedTask.Status
	}
	
	store.tasks[id] = task
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// DeleteTask deleta uma tarefa
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	id := vars["id"]
	
	store.mu.Lock()
	defer store.mu.Unlock()
	
	if _, exists := store.tasks[id]; !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	
	delete(store.tasks, id)
	
	w.WriteHeader(http.StatusNoContent)
}

// GetTask retorna uma tarefa específica
func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	id := vars["id"]
	
	store.mu.RLock()
	defer store.mu.RUnlock()
	
	task, exists := store.tasks[id]
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	
	json.NewEncoder(w).Encode(task)
}

// HealthCheck endpoint para verificar se o servidor está online
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// ============== MAIN ==============

func main() {
	router := mux.NewRouter()
	
	// Rotas da API
	router.HandleFunc("/api/health", HealthCheck).Methods("GET")
	router.HandleFunc("/api/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", GetTask).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", UpdateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", DeleteTask).Methods("DELETE")
	
	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           300,
	})
	
	handler := c.Handler(router)
	
	port := ":8080"
	log.Printf("🚀 Servidor iniciado em http://localhost%s\n", port)
	log.Printf("📝 API disponível em http://localhost%s/api\n", port)
	
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
