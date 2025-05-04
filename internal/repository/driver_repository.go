package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ARXXIII/f1-api/internal/model"
)

type DriverRepository interface {
	GetAllDrivers() ([]model.Driver, error)
}

type driverRepository struct {
	supabaseURL string
	apiKey      string
}

func NewDriverRepository(supabaseURL, apiKey string) DriverRepository {
	return &driverRepository{supabaseURL: supabaseURL, apiKey: apiKey}
}

func (r *driverRepository) GetAllDrivers() ([]model.Driver, error) {
	url := fmt.Sprintf("%s/driver", r.supabaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("apikey", r.apiKey)
	req.Header.Add("Authorization", "Bearer "+r.apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s, body: %s", resp.Status, string(body))
	}

	var drivers []model.Driver
	if err := json.Unmarshal(body, &drivers); err != nil {
		return nil, err
	}

	return drivers, nil
}
