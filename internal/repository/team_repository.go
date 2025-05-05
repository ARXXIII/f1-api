package repository

import (
	"context"

	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/pkg/db"
	"github.com/google/uuid"
)

type TeamRepository interface {
	GetAll(ctx context.Context) ([]model.Team, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Team, error)
}

type teamRepo struct{}

func NewTeamRepository() TeamRepository {
	return &teamRepo{}
}

func (r *teamRepo) GetAll(ctx context.Context) ([]model.Team, error) {
	rows, err := db.Conn.Query(ctx, `SELECT id, name, engine, chassis, debut, founder FROM teams`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []model.Team
	for rows.Next() {
		var t model.Team
		if err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Engine,
			&t.Chassis,
			&t.Debut,
			&t.Founder,
		); err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}

	return teams, nil
}

func (r *teamRepo) GetByID(ctx context.Context, id uuid.UUID) (*model.Team, error) {
	row := db.Conn.QueryRow(ctx, `SELECT id, name, engine, chassis, debut, founder FROM teams WHERE id = $1`, id)

	var t model.Team
	if err := row.Scan(
		&t.ID,
		&t.Name,
		&t.Engine,
		&t.Chassis,
		&t.Debut,
		&t.Founder,
	); err != nil {
		return nil, err
	}

	return &t, nil
}
