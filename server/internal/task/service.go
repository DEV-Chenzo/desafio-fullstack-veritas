package task

import (
	"context"
	"fmt"
	"strings"
)

type Service struct{ repository *Repository }

func NewService(repository *Repository) *Service                   { return &Service{repository: repository} }
func (s *Service) List(ctx context.Context) ([]Task, error)        { return s.repository.List(ctx) }
func (s *Service) Get(ctx context.Context, id int64) (Task, error) { return s.repository.Get(ctx, id) }
func (s *Service) Delete(ctx context.Context, id int64) error      { return s.repository.Delete(ctx, id) }
func (s *Service) Create(ctx context.Context, input Input) (Task, error) {
	if err := normalize(&input); err != nil {
		return Task{}, err
	}
	return s.repository.Create(ctx, input)
}
func (s *Service) Update(ctx context.Context, id int64, input Input) (Task, error) {
	if err := normalize(&input); err != nil {
		return Task{}, err
	}
	return s.repository.Update(ctx, id, input)
}

func normalize(input *Input) error {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)
	input.Status = Status(strings.ToLower(strings.TrimSpace(string(input.Status))))
	if input.Status == "" {
		input.Status = StatusTodo
	}
	if input.Title == "" || len(input.Title) > 120 {
		return fmt.Errorf("título é obrigatório e deve ter no máximo 120 caracteres")
	}
	if input.Status != StatusTodo && input.Status != StatusDoing && input.Status != StatusDone {
		return fmt.Errorf("status deve ser todo, doing ou done")
	}
	return nil
}
