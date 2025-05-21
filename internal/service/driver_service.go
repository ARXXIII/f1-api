package service

import (
	"context"

	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/internal/repository"
	"github.com/google/uuid"
)

type DriverService interface {
	GetDriver(ctx context.Context, page int) ([]model.Driver, error)
	GetDriverByName(ctx context.Context, name string, page int) ([]model.Driver, error)
	GetDriverByStatus(ctx context.Context, status string, page int) ([]model.Driver, error)
	GetDriverByNationality(ctx context.Context, nationality string, page int) ([]model.Driver, error)
	GetDriverByID(ctx context.Context, id uuid.UUID) (*model.Driver, error)
}

type driverService struct {
	repo repository.DriverRepository
}

func NewDriverService(r repository.DriverRepository) DriverService {
	return &driverService{repo: r}
}

func (s *driverService) GetDriver(ctx context.Context, page int) ([]model.Driver, error) {
	return s.repo.GetAll(ctx, page)
}

func (s *driverService) GetDriverByName(ctx context.Context, name string, page int) ([]model.Driver, error) {
	return s.repo.GetByName(ctx, name, page)
}

func (s *driverService) GetDriverByStatus(ctx context.Context, status string, page int) ([]model.Driver, error) {
	return s.repo.GetByStatus(ctx, status, page)
}

func (s *driverService) GetDriverByNationality(ctx context.Context, nationality string, page int) ([]model.Driver, error) {
	return s.repo.GetByNationality(ctx, nationality, page)
}

func (s *driverService) GetDriverByID(ctx context.Context, id uuid.UUID) (*model.Driver, error) {
	return s.repo.GetByID(ctx, id)
}
