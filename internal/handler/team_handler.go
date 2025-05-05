package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ARXXIII/f1-api/internal/service"
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

func (h *TeamHandler) GetTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := h.service.GetTeams(h.ctx)
	if err != nil {
		http.Error(w, "Failed to fetch teams", http.StatusInternalServerError)
		log.Printf("Error fetching teams: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
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

	team, err := h.service.GetTeamByID(h.ctx, id)
	if err != nil {
		http.Error(w, "Team not found", http.StatusNotFound)
		log.Printf("Error: Team not found: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(team)
}
