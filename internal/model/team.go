package model

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Engine  string    `json:"engine"`
	Chassis string    `json:"chassis"`
	Debut   time.Time `json:"debut"`
	Founder string    `json:"founder"`
}
