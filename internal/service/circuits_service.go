package service

import (
	"context"

	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/internal/repository"
	"github.com/google/uuid"
)

type CircuitService interface {
	GetCircuit(ctx context.Context, page int) ([]model.Circuit, error)
	GetCircuitByCurrent(ctx context.Context, current string, page int) ([]model.Circuit, error)
	GetCircuitByCountry(ctx context.Context, country string, page int) ([]model.Circuit, error)
	GetCircuitByID(ctx context.Context, id uuid.UUID) (*model.Circuit, error)
	GetCircuitByName(ctx context.Context, name string) (*model.Circuit, error)
}

type circuitService struct {
	repo repository.CircuitRepository
}

func NewCircuitService(r repository.CircuitRepository) CircuitService {
	return &circuitService{repo: r}
}

func (s *circuitService) GetCircuit(ctx context.Context, page int) ([]model.Circuit, error) {
	return s.repo.GetAll(ctx, page)
}

func (s *circuitService) GetCircuitByName(ctx context.Context, name string) (*model.Circuit, error) {
	return s.repo.GetByName(ctx, name)
}

func (s *circuitService) GetCircuitByCurrent(ctx context.Context, current string, page int) ([]model.Circuit, error) {
	return s.repo.GetByCurrent(ctx, current, page)
}

func (s *circuitService) GetCircuitByCountry(ctx context.Context, country string, page int) ([]model.Circuit, error) {
	return s.repo.GetByCountry(ctx, country, page)
}

func (s *circuitService) GetCircuitByID(ctx context.Context, id uuid.UUID) (*model.Circuit, error) {
	return s.repo.GetByID(ctx, id)
}
