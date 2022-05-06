package storage

import (
	"RestApi/internal/author/model"
	"context"
)

type Repository interface {
	Create(ctx context.Context, author *model.Author) error
	FindOne(ctx context.Context, id string) (model.Author, error)
	Update(ctx context.Context, user model.Author) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, sortOptions SortOptions) ([]model.Author, error)
}

type SortOptions interface {
	GetOrderBy() string
}
