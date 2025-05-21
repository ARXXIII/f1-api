package model

import (
	"github.com/google/uuid"
)

type Constructor struct {
	ID          uuid.UUID `json:"id"`
	Ref         string    `json:"ref"`
	Name        string    `json:"name"`
	Nationality string    `json:"nationality"`
	URL         string    `json:"url"`
}
