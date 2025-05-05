package model

import (
	"time"

	"github.com/google/uuid"
)

type Driver struct {
	ID           uuid.UUID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	PlaceOfBirth string    `json:"place_of_birth"`
	Number       int       `json:"number"`
	Debut        time.Time `json:"debut"`
	Team         uuid.UUID `json:"team"`
	Status       string    `json:"status"`
}
