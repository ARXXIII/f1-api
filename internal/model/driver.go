package model

import (
	"time"

	"github.com/google/uuid"
)

type Driver struct {
	ID          uuid.UUID `json:"id"`
	Ref         string    `json:"ref"`
	Code        *string   `json:"code"`
	Number      *int      `json:"number"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Nationality string    `json:"nationality"`
	Status      string    `json:"status"`
	URL         string    `json:"url"`
}
