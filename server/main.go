package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"veritas-kanban/internal/config"
	"veritas-kanban/internal/database"
	"veritas-kanban/internal/httpapi"
	"veritas-kanban/internal/task"
)

func main() {
	ctx := context.Background()
	cfg := config.Load()

	db, err := database.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("falha ao conectar ao PostgreSQL: %v", err)
	}
	defer db.Close()

	if err := database.Migrate(ctx, db); err != nil {
		log.Fatalf("falha ao preparar banco de dados: %v", err)
	}

	repository := task.NewRepository(db)
	service := task.NewService(repository)
	router := httpapi.NewRouter(service, db, cfg.CORSOrigin)

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("API disponível em http://localhost%s/api", server.Addr)
	log.Fatal(server.ListenAndServe())
}
