package service

import (
	"context"

	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/internal/repository"
	"github.com/google/uuid"
)

type TeamService interface {
	GetTeams(ctx context.Context) ([]model.Team, error)
	GetTeamByID(ctx context.Context, id uuid.UUID) (*model.Team, error)
}

type teamService struct {
	repo repository.TeamRepository
}

func NewTeamService(r repository.TeamRepository) TeamService {
	return &teamService{repo: r}
}

func (s *teamService) GetTeams(ctx context.Context) ([]model.Team, error) {
	return s.repo.GetAll(ctx)
}

func (s *teamService) GetTeamByID(ctx context.Context, id uuid.UUID) (*model.Team, error) {
	return s.repo.GetByID(ctx, id)
}
