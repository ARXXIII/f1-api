package repository

import (
	"context"

	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/pkg/db"
	"github.com/google/uuid"
)

type DriverRepository interface {
	GetAll(ctx context.Context) ([]model.Driver, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Driver, error)
}

type driverRepo struct{}

func NewDriverRepository() DriverRepository {
	return &driverRepo{}
}

func (r *driverRepo) GetAll(ctx context.Context) ([]model.Driver, error) {
	rows, err := db.Conn.Query(ctx, `SELECT id, first_name, last_name, date_of_birth, place_of_birth, number, debut, team, status FROM drivers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
