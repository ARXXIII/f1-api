package repository

import (
	"context"

	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/internal/utils"
	"github.com/ARXXIII/f1-api/pkg/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TeamRepository interface {
	GetAll(ctx context.Context, page int) ([]model.Team, error)
	GetByName(ctx context.Context, name string, page int) ([]model.Team, error)
	GetByEngine(ctx context.Context, engine string, page int) ([]model.Team, error)
	GetByChassis(ctx context.Context, chassis string, page int) ([]model.Team, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Team, error)
}

type teamRepo struct{}

func NewTeamRepository() TeamRepository {
	return &teamRepo{}
}

func (r *teamRepo) GetAll(ctx context.Context, page int) ([]model.Team, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, name, engine, chassis, debut, founder FROM teams ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := db.Conn.Query(ctx, query, utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanTeam(rows)
}

func (r *teamRepo) GetByName(ctx context.Context, name string, page int) ([]model.Team, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, name, engine, chassis, debut, founder 
	          FROM teams 
	          WHERE LOWER(name) LIKE LOWER($1)
	          ORDER BY id
	          LIMIT $2 OFFSET $3`
	rows, err := db.Conn.Query(ctx, query, "%"+name+"%", utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanTeam(rows)
}

func (r *teamRepo) GetByEngine(ctx context.Context, engine string, page int) ([]model.Team, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, name, engine, chassis, debut, founder  
	          FROM teams 
	          WHERE LOWER(engine) = LOWER($1)
	          ORDER BY id
	          LIMIT $2 OFFSET $3`
	rows, err := db.Conn.Query(ctx, query, engine, utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanTeam(rows)
}

func (r *teamRepo) GetByChassis(ctx context.Context, chassis string, page int) ([]model.Team, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, name, engine, chassis, debut, founder 
	          FROM teams 
	          WHERE LOWER(chassis) = LOWER($1)
	          ORDER BY id
	          LIMIT $2 OFFSET $3`
	rows, err := db.Conn.Query(ctx, query, chassis, utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanTeam(rows)
}

func scanTeam(rows pgx.Rows) ([]model.Team, error) {
	var teams []model.Team
	for rows.Next() {
		var d model.Team
		if err := rows.Scan(
			&d.ID,
			&d.Name,
			&d.Engine,
			&d.Chassis,
			&d.Debut,
			&d.Founder,
		); err != nil {
			return nil, err
		}
		teams = append(teams, d)
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
