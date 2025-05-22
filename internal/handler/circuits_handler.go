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

type CircuitHandler struct {
	service service.CircuitService
	ctx     context.Context
}

func NewCircuitHandler(ctx context.Context, s service.CircuitService) *CircuitHandler {
	return &CircuitHandler{
		service: s,
		ctx:     ctx,
	}
}

func (h *CircuitHandler) GetCircuit(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page := utils.ParsePage(query.Get("page"))

	switch {
	case query.Has("name"):
		h.getByName(w, query.Get("name"))
	case query.Has("country"):
		h.getByCountry(w, query.Get("country"), page)
	case query.Has("current"):
		h.getByCurrent(w, query.Get("current"), page)
	default:
		h.getAll(w, page)
	}
}

func (h *CircuitHandler) GetCircuitByID(w http.ResponseWriter, r *http.Request) {
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
// Private methods
// ------------------------

func (h *CircuitHandler) getAll(w http.ResponseWriter, page int) {
	circuits, err := h.service.GetCircuit(h.ctx, page)
	if err != nil {
		http.Error(w, "Failed to fetch circuits", http.StatusInternalServerError)
		log.Printf("GetAll error: %v", err)
		return
	}
	h.respond(w, circuits)
}

func (h *CircuitHandler) getByName(w http.ResponseWriter, name string) {
	circuit, err := h.service.GetCircuitByName(h.ctx, name)
	if err != nil {
		http.Error(w, "Failed to fetch circuit by name", http.StatusInternalServerError)
		log.Printf("GetByName error: %v", err)
		return
	}
	h.respond(w, circuit)
}

func (h *CircuitHandler) getByCountry(w http.ResponseWriter, country string, page int) {
	circuits, err := h.service.GetCircuitByCountry(h.ctx, country, page)
	if err != nil {
		http.Error(w, "Failed to fetch circuits by country", http.StatusInternalServerError)
		log.Printf("GetByCountry error: %v", err)
		return
	}
	h.respond(w, circuits)
}

func (h *CircuitHandler) getByCurrent(w http.ResponseWriter, current string, page int) {
	circuits, err := h.service.GetCircuitByCurrent(h.ctx, current, page)
	if err != nil {
		http.Error(w, "Failed to fetch circuits by current", http.StatusInternalServerError)
		log.Printf("GetByCurrent error: %v", err)
		return
	}
	h.respond(w, circuits)
}

func (h *CircuitHandler) getByID(w http.ResponseWriter, id uuid.UUID) {
	circuit, err := h.service.GetCircuitByID(h.ctx, id)
	if err != nil {
		http.Error(w, "Circuit not found", http.StatusNotFound)
		log.Printf("GetByID error: %v", err)
		return
	}
	h.respond(w, circuit)
}

func (h *CircuitHandler) respond(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
