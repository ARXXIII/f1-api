package repository

import (
	"context"

	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/internal/utils"
	"github.com/ARXXIII/f1-api/pkg/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ConstructorRepository interface {
	GetAll(ctx context.Context, page int) ([]model.Constructor, error)
	GetByName(ctx context.Context, name string, page int) ([]model.Constructor, error)
	GetByNationality(ctx context.Context, nationality string, page int) ([]model.Constructor, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Constructor, error)
}

type constructorRepo struct{}

func NewConstructorRepository() ConstructorRepository {
	return &constructorRepo{}
}

func (r *constructorRepo) GetAll(ctx context.Context, page int) ([]model.Constructor, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, ref, name, nationality, url FROM constructors ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := db.Conn.Query(ctx, query, utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanConstructor(rows)
}

func (r *constructorRepo) GetByName(ctx context.Context, name string, page int) ([]model.Constructor, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, ref, name, nationality, url
	          FROM constructors 
	          WHERE LOWER(name) LIKE LOWER($1)
	          ORDER BY id
	          LIMIT $2 OFFSET $3`
	rows, err := db.Conn.Query(ctx, query, "%"+name+"%", utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanConstructor(rows)
}

func (r *constructorRepo) GetByNationality(ctx context.Context, nationality string, page int) ([]model.Constructor, error) {
	offset := (page - 1) * utils.DEFAULT_PAGE_SIZE
	query := `SELECT id, ref, name, nationality, url
	          FROM constructors 
	          WHERE LOWER(nationality) = LOWER($1)
	          ORDER BY id
	          LIMIT $2 OFFSET $3`
	rows, err := db.Conn.Query(ctx, query, nationality, utils.DEFAULT_PAGE_SIZE, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanConstructor(rows)
}

func scanConstructor(rows pgx.Rows) ([]model.Constructor, error) {
	var constructors []model.Constructor
	for rows.Next() {
		var d model.Constructor
		if err := rows.Scan(
			&d.ID,
			&d.Ref,
			&d.Name,
			&d.Nationality,
			&d.URL,
		); err != nil {
			return nil, err
		}
		constructors = append(constructors, d)
	}
	return constructors, nil
}

func (r *constructorRepo) GetByID(ctx context.Context, id uuid.UUID) (*model.Constructor, error) {
	row := db.Conn.QueryRow(ctx, `SELECT id, ref, name, nationality, url FROM constructors WHERE id = $1`, id)

	var t model.Constructor
	if err := row.Scan(
		&t.ID,
		&t.Ref,
		&t.Name,
		&t.Nationality,
		&t.URL,
	); err != nil {
		return nil, err
	}

	return &t, nil
}
