package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ARXXIII/f1-api/internal/service"
)

type DriverHandler struct {
	service service.DriverService
}

func NewDriverHandler(service service.DriverService) *DriverHandler {
	return &DriverHandler{service: service}
}

func (h *DriverHandler) GetDrivers(w http.ResponseWriter, r *http.Request) {
	drivers, err := h.service.GetDrivers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drivers)
}
