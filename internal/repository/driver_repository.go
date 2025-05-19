package repository

import (
	"context"

	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/internal/utils"
	"github.com/ARXXIII/f1-api/pkg/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type DriverRepository interface {
	GetAll(ctx context.Context, page int) ([]model.Driver, error)
	GetByName(ctx context.Context, name string, page int) ([]model.Driver, error)
	GetByTeam(ctx context.Context, team string, page int) ([]model.Driver, error)
	GetByStatus(ctx context.Context, status string, page int) ([]model.Driver, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Driver, error)
}

type driverRepo struct{}

func NewDriverRepository() DriverRepository {
	return &driverRepo{}
}

func (r *driverRepo) GetAll(ctx context.Context, page int) ([]model.Driver, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, first_name, last_name, date_of_birth, place_of_birth, number, debut, team, status FROM drivers ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := db.Conn.Query(ctx, query, utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanDriver(rows)
}

func (r *driverRepo) GetByName(ctx context.Context, name string, page int) ([]model.Driver, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, first_name, last_name, date_of_birth, place_of_birth, number, debut, team, status 
	          FROM drivers 
	          WHERE LOWER(first_name) LIKE LOWER($1) OR LOWER(last_name) LIKE LOWER($1)
	          ORDER BY last_name
	          LIMIT $2 OFFSET $3`
	rows, err := db.Conn.Query(ctx, query, "%"+name+"%", utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanDriver(rows)
}

func (r *driverRepo) GetByTeam(ctx context.Context, team string, page int) ([]model.Driver, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, first_name, last_name, date_of_birth, place_of_birth, number, debut, team, status 
	          FROM drivers 
	          WHERE LOWER(team) = LOWER($1)
	          ORDER BY last_name
	          LIMIT $2 OFFSET $3`
	rows, err := db.Conn.Query(ctx, query, team, utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanDriver(rows)
}

func (r *driverRepo) GetByStatus(ctx context.Context, status string, page int) ([]model.Driver, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, first_name, last_name, date_of_birth, place_of_birth, number, debut, team, status 
	          FROM drivers 
	          WHERE LOWER(status) = LOWER($1)
	          ORDER BY last_name
	          LIMIT $2 OFFSET $3`
	rows, err := db.Conn.Query(ctx, query, status, utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanDriver(rows)
}

func scanDriver(rows pgx.Rows) ([]model.Driver, error) {
	var drivers []model.Driver
	for rows.Next() {
		var d model.Driver
		if err := rows.Scan(
			&d.ID,
			&d.FirstName,
			&d.LastName,
			&d.DateOfBirth,
			&d.PlaceOfBirth,
			&d.Number,
			&d.Debut,
			&d.Team,
			&d.Status,
		); err != nil {
			return nil, err
		}
		drivers = append(drivers, d)
	}
	return drivers, nil
}

func (r *driverRepo) GetByID(ctx context.Context, id uuid.UUID) (*model.Driver, error) {
	row := db.Conn.QueryRow(ctx, `SELECT id, first_name, last_name, date_of_birth, place_of_birth, number, debut, team, status FROM drivers WHERE id = $1`, id)

	var d model.Driver
	if err := row.Scan(
		&d.ID,
		&d.FirstName,
		&d.LastName,
		&d.DateOfBirth,
		&d.PlaceOfBirth,
		&d.Number,
		&d.Debut,
		&d.Team,
		&d.Status,
	); err != nil {
		return nil, err
	}

	return &d, nil
}
