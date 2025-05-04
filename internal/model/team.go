package model

type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Engine  string `json:"engine"`
	Chassis string `json:"chassis"`
	Debut   string `json:"debut"`
	Founder string `json:"founder"`
}
