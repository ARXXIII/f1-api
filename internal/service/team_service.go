package service

import (
	"context"

	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/internal/repository"
	"github.com/google/uuid"
)

type TeamService interface {
	GetTeam(ctx context.Context, page int) ([]model.Team, error)
	GetTeamByName(ctx context.Context, name string, page int) ([]model.Team, error)
	GetTeamByEngine(ctx context.Context, chassis string, page int) ([]model.Team, error)
	GetTeamByChassis(ctx context.Context, engine string, page int) ([]model.Team, error)
	GetTeamByID(ctx context.Context, id uuid.UUID) (*model.Team, error)
}

type teamService struct {
	repo repository.TeamRepository
}

func NewTeamService(r repository.TeamRepository) TeamService {
	return &teamService{repo: r}
}

func (s *teamService) GetTeam(ctx context.Context, page int) ([]model.Team, error) {
	return s.repo.GetAll(ctx, page)
}

func (s *teamService) GetTeamByName(ctx context.Context, name string, page int) ([]model.Team, error) {
	return s.repo.GetByName(ctx, name, page)
}

func (s *teamService) GetTeamByEngine(ctx context.Context, engine string, page int) ([]model.Team, error) {
	return s.repo.GetByEngine(ctx, engine, page)
}

func (s *teamService) GetTeamByChassis(ctx context.Context, chassis string, page int) ([]model.Team, error) {
	return s.repo.GetByChassis(ctx, chassis, page)
}

func (s *teamService) GetTeamByID(ctx context.Context, id uuid.UUID) (*model.Team, error) {
	return s.repo.GetByID(ctx, id)
}
