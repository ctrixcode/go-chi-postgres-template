package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ctrixcode/go-chi-postgres/internal/models"
	"github.com/ctrixcode/go-chi-postgres/internal/services"
	"github.com/ctrixcode/go-chi-postgres/pkg/errors"
	"github.com/ctrixcode/go-chi-postgres/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ExampleHandler struct {
	service   services.ExampleService
	validator *validator.Validate
}

func NewExampleHandler(service services.ExampleService) *ExampleHandler {
	return &ExampleHandler{
		service:   service,
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
		response.JSONError(w, errors.BadRequestError(errors.ErrBadRequest, err.Error()))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.JSONError(w, errors.BadRequestError(errors.ErrValidationFailed, err.Error()))
		return
	}

	example, err := h.service.Create(r.Context(), req)
	if err != nil {
		response.JSONError(w, errors.InternalServerError(errors.ErrInternalServerError, err.Error()))
		return
	}

	response.JSONSuccess(w, example, http.StatusCreated, "Example created successfully")
}

func (h *ExampleHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.JSONError(w, errors.BadRequestError(errors.ErrBadRequest, "Invalid UUID"))
		return
	}

	example, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		response.JSONError(w, errors.NotFoundError(errors.ErrNotFound, "Example not found"))
		return
	}

	response.JSONSuccess(w, example, http.StatusOK)
}

func (h *ExampleHandler) List(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, _ := strconv.ParseUint(limitStr, 10, 64)
	if limit == 0 {
		limit = 10
	}
	offset, _ := strconv.ParseUint(offsetStr, 10, 64)

	examples, err := h.service.List(r.Context(), limit, offset)
	if err != nil {
		response.JSONError(w, errors.InternalServerError(errors.ErrInternalServerError, err.Error()))
		return
	}

	response.JSONSuccess(w, examples, http.StatusOK)
}

func (h *ExampleHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.JSONError(w, errors.BadRequestError(errors.ErrBadRequest, "Invalid UUID"))
		return
	}

	var req models.UpdateExampleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSONError(w, errors.BadRequestError(errors.ErrBadRequest, err.Error()))
		return
	}

	example, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		response.JSONError(w, errors.InternalServerError(errors.ErrInternalServerError, err.Error()))
		return
	}

	response.JSONSuccess(w, example, http.StatusOK, "Example updated successfully")
}

func (h *ExampleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.JSONError(w, errors.BadRequestError(errors.ErrBadRequest, "Invalid UUID"))
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		response.JSONError(w, errors.InternalServerError(errors.ErrInternalServerError, err.Error()))
		return
	}

	response.JSONSuccess(w, nil, http.StatusOK, "Example deleted successfully")
}
