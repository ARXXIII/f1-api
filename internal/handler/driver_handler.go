package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ARXXIII/f1-api/internal/service"
	"github.com/ARXXIII/f1-api/internal/utils"
	"github.com/google/uuid"
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

func (h *DriverHandler) GetDriver(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page := utils.ParsePage(query.Get("page"))

	switch {
	case query.Has("name"):
		h.getByName(w, query.Get("name"), page)
	case query.Has("team"):
		h.getByTeam(w, query.Get("team"), page)
	case query.Has("status"):
		h.getByStatus(w, query.Get("status"), page)
	default:
		h.getAll(w, page)
	}
}

func (h *DriverHandler) GetDriverByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		log.Printf("Error: Invalid ID format: %v", err)
		return
	}

	h.getByID(w, id)
}

// ------------------------
// Приватные методы
// ------------------------

func (h *DriverHandler) getAll(w http.ResponseWriter, page int) {
	drivers, err := h.service.GetDriver(h.ctx, page)
	if err != nil {
		http.Error(w, "Failed to fetch drivers", http.StatusInternalServerError)
		log.Printf("GetAll error: %v", err)
		return
	}
	h.respond(w, drivers)
}

func (h *DriverHandler) getByName(w http.ResponseWriter, name string, page int) {
	drivers, err := h.service.GetDriverByName(h.ctx, name, page)
	if err != nil {
		http.Error(w, "Failed to fetch drivers by name", http.StatusInternalServerError)
		log.Printf("GetByName error: %v", err)
		return
	}
	h.respond(w, drivers)
}

func (h *DriverHandler) getByTeam(w http.ResponseWriter, team string, page int) {
	drivers, err := h.service.GetDriverByTeam(h.ctx, team, page)
	if err != nil {
		http.Error(w, "Failed to fetch drivers by team", http.StatusInternalServerError)
		log.Printf("GetByTeam error: %v", err)
		return
	}
	h.respond(w, drivers)
}

func (h *DriverHandler) getByStatus(w http.ResponseWriter, status string, page int) {
	drivers, err := h.service.GetDriverByStatus(h.ctx, status, page)
	if err != nil {
		http.Error(w, "Failed to fetch drivers by status", http.StatusInternalServerError)
		log.Printf("GetByStatus error: %v", err)
		return
	}
	h.respond(w, drivers)
}

func (h *DriverHandler) getByID(w http.ResponseWriter, id uuid.UUID) {
	driver, err := h.service.GetDriverByID(h.ctx, id)
	if err != nil {
		http.Error(w, "Driver not found", http.StatusNotFound)
		log.Printf("GetByID error: %v", err)
		return
	}
	h.respond(w, driver)
}

func (h *DriverHandler) respond(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
