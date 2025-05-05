package service

import (
	"context"

	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/internal/repository"
	"github.com/google/uuid"
)

type DriverService interface {
	GetDrivers(ctx context.Context) ([]model.Driver, error)
	GetDriverByID(ctx context.Context, id uuid.UUID) (*model.Driver, error)
}

type driverService struct {
	repo repository.DriverRepository
}

func NewDriverService(r repository.DriverRepository) DriverService {
	return &driverService{repo: r}
}

func (s *driverService) GetDrivers(ctx context.Context) ([]model.Driver, error) {
	return s.repo.GetAll(ctx)
}

func (s *driverService) GetDriverByID(ctx context.Context, id uuid.UUID) (*model.Driver, error) {
	return s.repo.GetByID(ctx, id)
}
