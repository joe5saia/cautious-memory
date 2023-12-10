package repository

import (
	"context"
)

// GenericRepository is a generic interface for CRUD operations
type GenericRepository[T any] interface {
	ReadAll(ctx context.Context, page int, pageSize int, preloads ...string) ([]T, error)
	FindByID(ctx context.Context, id uint) (*T, error)
	Create(ctx context.Context, entity *T) error
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id uint) error
}
