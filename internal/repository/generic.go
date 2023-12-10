package repository

import (
	"context"

	"gorm.io/gorm"
)

type gormRepository[T any] struct {
	db *gorm.DB
}

// NewGormRepository creates a new instance of a generic GORM repository
func NewGormRepository[T any](db *gorm.DB) GenericRepository[T] {
	return &gormRepository[T]{db: db}
}

// ReadAll reads all entities of type T from the database and preloads any related models.
// It uses pagination based on the given page and pageSize.
// It returns a slice of entities and any error encountered.
func (r *gormRepository[T]) ReadAll(ctx context.Context, page int, pageSize int, preloads ...string) ([]T, error) {
	var entities []T
	db := r.db.WithContext(ctx)
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	if err := db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *gormRepository[T]) FindByID(ctx context.Context, id uint) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *gormRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *gormRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r *gormRepository[T]) Delete(ctx context.Context, id uint) error {
	var entity T
	return r.db.WithContext(ctx).Delete(&entity, id).Error
}
