package model

type Driver struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	DateOfBirth  string `json:"date_of_birth"`
	PlaceOfBirth string `json:"place_of_birth"`
	Number       int    `json:"number"`
	Debut        string `json:"debut"`
	Team         string `json:"team"`
	Status       string `json:"status"`
}
