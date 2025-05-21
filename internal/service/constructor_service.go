package service

import (
	"context"

	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/internal/repository"
	"github.com/google/uuid"
)

type ConstructorService interface {
	GetConstructor(ctx context.Context, page int) ([]model.Constructor, error)
	GetConstructorByName(ctx context.Context, name string, page int) ([]model.Constructor, error)
	GetConstructorByNationality(ctx context.Context, nationality string, page int) ([]model.Constructor, error)
	GetConstructorByID(ctx context.Context, id uuid.UUID) (*model.Constructor, error)
}

type constructorService struct {
	repo repository.ConstructorRepository
}

func NewConstructorService(r repository.ConstructorRepository) ConstructorService {
	return &constructorService{repo: r}
}

func (s *constructorService) GetConstructor(ctx context.Context, page int) ([]model.Constructor, error) {
	return s.repo.GetAll(ctx, page)
}

func (s *constructorService) GetConstructorByName(ctx context.Context, name string, page int) ([]model.Constructor, error) {
	return s.repo.GetByName(ctx, name, page)
}

func (s *constructorService) GetConstructorByNationality(ctx context.Context, nationality string, page int) ([]model.Constructor, error) {
	return s.repo.GetByNationality(ctx, nationality, page)
}

func (s *constructorService) GetConstructorByID(ctx context.Context, id uuid.UUID) (*model.Constructor, error) {
	return s.repo.GetByID(ctx, id)
}
