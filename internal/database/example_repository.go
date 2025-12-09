package database

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/ctrixcode/go-chi-postgres/internal/models"
	"github.com/ctrixcode/go-chi-postgres/internal/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type exampleRepository struct {
	db *sqlx.DB
}

func NewExampleRepository(db *sqlx.DB) repository.ExampleRepository {
	return &exampleRepository{db: db}
}

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func (r *exampleRepository) Create(ctx context.Context, req models.CreateExampleRequest) (*models.Example, error) {
	query := psql.Insert("examples").
		Columns("name", "lucky_number", "is_premium").
		Values(req.Name, req.LuckyNumber, req.IsPremium).
		Suffix("RETURNING id, name, lucky_number, is_premium, created_at, updated_at")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var example models.Example
	err = r.db.GetContext(ctx, &example, sql, args...)
	if err != nil {
		return nil, err
	}

	return &example, nil
}

func (r *exampleRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Example, error) {
	query := psql.Select("*").From("examples").Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var example models.Example
	err = r.db.GetContext(ctx, &example, sql, args...)
	if err != nil {
		return nil, err
	}

	return &example, nil
}

func (r *exampleRepository) List(ctx context.Context, limit, offset uint64) ([]models.Example, error) {
	query := psql.Select("*").From("examples").Limit(limit).Offset(offset)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var examples []models.Example
	err = r.db.SelectContext(ctx, &examples, sql, args...)
	if err != nil {
		return nil, err
	}

	return examples, nil
}

func (r *exampleRepository) Update(ctx context.Context, id uuid.UUID, req models.UpdateExampleRequest) (*models.Example, error) {
	updateBuilder := psql.Update("examples").Where(squirrel.Eq{"id": id}).Set("updated_at", time.Now())

	if req.Name != nil {
		updateBuilder = updateBuilder.Set("name", *req.Name)
	}
	if req.LuckyNumber != nil {
		updateBuilder = updateBuilder.Set("lucky_number", *req.LuckyNumber)
	}
	if req.IsPremium != nil {
		updateBuilder = updateBuilder.Set("is_premium", *req.IsPremium)
	}

	updateBuilder = updateBuilder.Suffix("RETURNING id, name, lucky_number, is_premium, created_at, updated_at")

	sql, args, err := updateBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var example models.Example
	err = r.db.GetContext(ctx, &example, sql, args...)
	if err != nil {
		return nil, err
	}

	return &example, nil
}

func (r *exampleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := psql.Delete("examples").Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("example not found")
	}

	return nil
}
