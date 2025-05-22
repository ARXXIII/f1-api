package model

import (
	"github.com/google/uuid"
)

type Circuit struct {
	ID       uuid.UUID `json:"id"`
	Ref      string    `json:"ref"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Country  string    `json:"country"`
	Current  bool      `json:"current"`
	URL      string    `json:"url"`
}
