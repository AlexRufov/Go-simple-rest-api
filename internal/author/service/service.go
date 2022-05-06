package service

import (
	"RestApi/internal/author/model"
	"RestApi/internal/author/storage"
	"RestApi/pkg/api/sort"
	"RestApi/pkg/logging"
	"context"
)

type Service struct {
	repository storage.Repository
	logger     *logging.Logger
}

func NewService(repository storage.Repository, logger *logging.Logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) GetAll(ctx context.Context, sortOptions sort.Options) ([]model.Author, error) {
	options := storage.NewSortOptions(sortOptions.Field, sortOptions.Order)
	all, err := s.repository.FindAll(ctx, options)
	if err != nil {
		return nil, err
	}
	return all, nil
}
