package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/ARXXIII/f1-api/internal/service"
)

type DriverHandler struct {
	service service.DriverService
	ctx     context.Context
}

func NewDriverHandler(ctx context.Context, s service.DriverService) *DriverHandler {
	return &DriverHandler{
		service: s,
		ctx:     ctx,
	}
}

func (h *DriverHandler) GetDrivers(w http.ResponseWriter, r *http.Request) {
	drivers, err := h.service.GetDrivers(h.ctx)
	if err != nil {
		http.Error(w, "Failed to fetch drivers", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drivers)
}

func (h *DriverHandler) GetDriverByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	driver, err := h.service.GetDriverByID(h.ctx, id)
	if err != nil {
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(driver)
}
