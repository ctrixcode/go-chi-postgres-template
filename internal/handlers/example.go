package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ctrixcode/go-chi-postgres/internal/models"
	"github.com/ctrixcode/go-chi-postgres/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ExampleHandler struct {
	repo      repository.ExampleRepository
	validator *validator.Validate
}

func NewExampleHandler(repo repository.ExampleRepository) *ExampleHandler {
	return &ExampleHandler{
		repo:      repo,
		validator: validator.New(),
	}
}

func (h *ExampleHandler) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", h.Create)
	r.Get("/", h.List)
	r.Get("/{id}", h.Get)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
	return r
}

func (h *ExampleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateExampleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	example, err := h.repo.Create(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(example)
}

func (h *ExampleHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	example, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Example not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(example)
}

func (h *ExampleHandler) List(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, _ := strconv.ParseUint(limitStr, 10, 64)
	if limit == 0 {
		limit = 10
	}
	offset, _ := strconv.ParseUint(offsetStr, 10, 64)

	examples, err := h.repo.List(r.Context(), limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(examples)
}

func (h *ExampleHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	var req models.UpdateExampleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	example, err := h.repo.Update(r.Context(), id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(example)
}

func (h *ExampleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
