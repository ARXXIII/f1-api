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

type ConstructorHandler struct {
	service service.ConstructorService
	ctx     context.Context
}

func NewConstructorHandler(ctx context.Context, s service.ConstructorService) *ConstructorHandler {
	return &ConstructorHandler{
		service: s,
		ctx:     ctx,
	}
}

func (h *ConstructorHandler) GetConstructor(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page := utils.ParsePage(query.Get("page"))

	switch {
	case query.Has("name"):
		h.getByName(w, query.Get("name"), page)
	case query.Has("nationality"):
		h.getByNationality(w, query.Get("nationality"), page)
	default:
		h.getAll(w, page)
	}
}

func (h *ConstructorHandler) GetConstructorByID(w http.ResponseWriter, r *http.Request) {
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

func (h *ConstructorHandler) getAll(w http.ResponseWriter, page int) {
	constructors, err := h.service.GetConstructor(h.ctx, page)
	if err != nil {
		http.Error(w, "Failed to fetch constructors", http.StatusInternalServerError)
		log.Printf("GetAll error: %v", err)
		return
	}
	h.respond(w, constructors)
}

func (h *ConstructorHandler) getByName(w http.ResponseWriter, name string, page int) {
	constructor, err := h.service.GetConstructorByName(h.ctx, name, page)
	if err != nil {
		http.Error(w, "Failed to fetch constructor by name", http.StatusInternalServerError)
		log.Printf("GetByName error: %v", err)
		return
	}
	h.respond(w, constructor)
}

func (h *ConstructorHandler) getByNationality(w http.ResponseWriter, nationality string, page int) {
	constructors, err := h.service.GetConstructorByNationality(h.ctx, nationality, page)
	if err != nil {
		http.Error(w, "Failed to fetch constructors by nationality", http.StatusInternalServerError)
		log.Printf("GetByNationality error: %v", err)
		return
	}
	h.respond(w, constructors)
}

func (h *ConstructorHandler) getByID(w http.ResponseWriter, id uuid.UUID) {
	constructor, err := h.service.GetConstructorByID(h.ctx, id)
	if err != nil {
		http.Error(w, "Constructor not found", http.StatusNotFound)
		log.Printf("GetByID error: %v", err)
		return
	}
	h.respond(w, constructor)
}

func (h *ConstructorHandler) respond(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
