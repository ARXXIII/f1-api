package repository

import (
	"context"

	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/internal/utils"
	"github.com/ARXXIII/f1-api/pkg/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CircuitRepository interface {
	GetAll(ctx context.Context, page int) ([]model.Circuit, error)
	GetByCurrent(ctx context.Context, current string, page int) ([]model.Circuit, error)
	GetByCountry(ctx context.Context, country string, page int) ([]model.Circuit, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Circuit, error)
	GetByName(ctx context.Context, name string) (*model.Circuit, error)
}

type circuitRepo struct{}

func NewCircuitRepository() CircuitRepository {
	return &circuitRepo{}
}

func (r *circuitRepo) GetAll(ctx context.Context, page int) ([]model.Circuit, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, ref, name, location, country, current, url FROM circuits ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := db.Conn.Query(ctx, query, utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanCircuit(rows)
}

func (r *circuitRepo) GetByName(ctx context.Context, name string) (*model.Circuit, error) {
	row := db.Conn.QueryRow(ctx, `SELECT id, ref, name, location, country, current, url FROM circuits WHERE LOWER(name) LIKE LOWER($1)`, name)

	var d model.Circuit
	if err := row.Scan(
		&d.ID,
		&d.Ref,
		&d.Name,
		&d.Location,
		&d.Country,
		&d.Current,
		&d.URL,
	); err != nil {
		return nil, err
	}

	return &d, nil
}

func (r *circuitRepo) GetByCurrent(ctx context.Context, current string, page int) ([]model.Circuit, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, ref, name, location, country, current, url 
	          FROM circuits 
	          WHERE current = $1
	          ORDER BY name
	          LIMIT $2 OFFSET $3`
	rows, err := db.Conn.Query(ctx, query, current, utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanCircuit(rows)
}

func (r *circuitRepo) GetByCountry(ctx context.Context, country string, page int) ([]model.Circuit, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, ref, name, location, country, current, url 
	          FROM circuits 
	          WHERE LOWER(country) = LOWER($1)
	          ORDER BY name
	          LIMIT $2 OFFSET $3`
	rows, err := db.Conn.Query(ctx, query, country, utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanCircuit(rows)
}

func scanCircuit(rows pgx.Rows) ([]model.Circuit, error) {
	var circuits []model.Circuit
	for rows.Next() {
		var d model.Circuit
		if err := rows.Scan(
			&d.ID,
			&d.Ref,
			&d.Name,
			&d.Location,
			&d.Country,
			&d.Current,
			&d.URL,
		); err != nil {
			return nil, err
		}
		circuits = append(circuits, d)
	}
	return circuits, nil
}

func (r *circuitRepo) GetByID(ctx context.Context, id uuid.UUID) (*model.Circuit, error) {
	row := db.Conn.QueryRow(ctx, `SELECT id, ref, name, location, country, current, url FROM circuits WHERE id = $1`, id)

	var d model.Circuit
	if err := row.Scan(
		&d.ID,
		&d.Ref,
		&d.Name,
		&d.Location,
		&d.Country,
		&d.Current,
		&d.URL,
	); err != nil {
		return nil, err
	}

	return &d, nil
}
