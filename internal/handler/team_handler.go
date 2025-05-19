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

type TeamHandler struct {
	service service.TeamService
	ctx     context.Context
}

func NewTeamHandler(ctx context.Context, s service.TeamService) *TeamHandler {
	return &TeamHandler{
		service: s,
		ctx:     ctx,
	}
}

func (h *TeamHandler) GetTeam(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page := utils.ParsePage(query.Get("page"))

	switch {
	case query.Has("name"):
		h.getByName(w, query.Get("name"), page)
	case query.Has("engine"):
		h.getByEngine(w, query.Get("engine"), page)
	case query.Has("chassis"):
		h.getByChassis(w, query.Get("chassis"), page)
	default:
		h.getAll(w, page)
	}
}

func (h *TeamHandler) GetTeamByID(w http.ResponseWriter, r *http.Request) {
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

func (h *TeamHandler) getAll(w http.ResponseWriter, page int) {
	teams, err := h.service.GetTeam(h.ctx, page)
	if err != nil {
		http.Error(w, "Failed to fetch teams", http.StatusInternalServerError)
		log.Printf("GetAll error: %v", err)
		return
	}
	h.respond(w, teams)
}

func (h *TeamHandler) getByName(w http.ResponseWriter, name string, page int) {
	team, err := h.service.GetTeamByName(h.ctx, name, page)
	if err != nil {
		http.Error(w, "Failed to fetch teams by name", http.StatusInternalServerError)
		log.Printf("GetByName error: %v", err)
		return
	}
	h.respond(w, team)
}

func (h *TeamHandler) getByEngine(w http.ResponseWriter, engine string, page int) {
	teams, err := h.service.GetTeamByEngine(h.ctx, engine, page)
	if err != nil {
		http.Error(w, "Failed to fetch teams by engine", http.StatusInternalServerError)
		log.Printf("GetByName error: %v", err)
		return
	}
	h.respond(w, teams)
}

func (h *TeamHandler) getByChassis(w http.ResponseWriter, chassis string, page int) {
	teams, err := h.service.GetTeamByChassis(h.ctx, chassis, page)
	if err != nil {
		http.Error(w, "Failed to fetch teams by chassis", http.StatusInternalServerError)
		log.Printf("GetByName error: %v", err)
		return
	}
	h.respond(w, teams)
}

func (h *TeamHandler) getByID(w http.ResponseWriter, id uuid.UUID) {
	team, err := h.service.GetTeamByID(h.ctx, id)
	if err != nil {
		http.Error(w, "Team not found", http.StatusNotFound)
		log.Printf("GetByID error: %v", err)
		return
	}
	h.respond(w, team)
}

func (h *TeamHandler) respond(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
