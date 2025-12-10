package services

import (
	"context"

	"github.com/ctrixcode/go-chi-postgres/internal/models"
	"github.com/ctrixcode/go-chi-postgres/internal/repository"
	"github.com/google/uuid"
)

type ExampleService interface {
	Create(ctx context.Context, req models.CreateExampleRequest) (*models.Example, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Example, error)
	List(ctx context.Context, limit, offset uint64) ([]models.Example, error)
	Update(ctx context.Context, id uuid.UUID, req models.UpdateExampleRequest) (*models.Example, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type exampleService struct {
	repo repository.ExampleRepository
}

func NewExampleService(repo repository.ExampleRepository) ExampleService {
	return &exampleService{
		repo: repo,
	}
}

func (s *exampleService) Create(ctx context.Context, req models.CreateExampleRequest) (*models.Example, error) {
	// In a real application, you might have business logic here.
	// For example, validating the request, calling other services, etc.
	return s.repo.Create(ctx, req)
}

func (s *exampleService) GetByID(ctx context.Context, id uuid.UUID) (*models.Example, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *exampleService) List(ctx context.Context, limit, offset uint64) ([]models.Example, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *exampleService) Update(ctx context.Context, id uuid.UUID, req models.UpdateExampleRequest) (*models.Example, error) {
	return s.repo.Update(ctx, id, req)
}

func (s *exampleService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
