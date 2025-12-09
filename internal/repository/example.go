package repository

import (
	"context"

	"github.com/ctrixcode/go-chi-postgres/internal/models"
	"github.com/google/uuid"
)

type ExampleRepository interface {
	Create(ctx context.Context, req models.CreateExampleRequest) (*models.Example, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Example, error)
	List(ctx context.Context, limit, offset uint64) ([]models.Example, error)
	Update(ctx context.Context, id uuid.UUID, req models.UpdateExampleRequest) (*models.Example, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
