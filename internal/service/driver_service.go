package service

import (
	"github.com/ARXXIII/f1-api/internal/model"
	"github.com/ARXXIII/f1-api/internal/repository"
)

type DriverService interface {
	GetDrivers() ([]model.Driver, error)
}

type driverService struct {
	repo repository.DriverRepository
}

func NewDriverService(repo repository.DriverRepository) DriverService {
	return &driverService{repo: repo}
}

func (s *driverService) GetDrivers() ([]model.Driver, error) {
	drivers, err := s.repo.GetAllDrivers()
	if err != nil {
		return nil, err
	}
	return drivers, nil
}
