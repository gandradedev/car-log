package handler

import (
	"encoding/json"
	"net/http"

	maintenancetypeusecase "github.com/gaaandrade/car-log/internal/application/usecase/maintenance_type"
	"github.com/gaaandrade/car-log/internal/application/usecase/maintenance_type/dto"
	domainerrors "github.com/gaaandrade/car-log/internal/domain/errors"
)

// compile-time reference so swag can resolve domainerrors.CustomError in annotations.
var _ domainerrors.CustomError

// MaintenanceTypeHandler defines the HTTP contract for maintenance type operations.
type MaintenanceTypeHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type maintenanceTypeHandler struct {
	createUseCase maintenancetypeusecase.CreateMaintenanceTypeUseCase
	listUseCase   maintenancetypeusecase.ListMaintenanceTypesUseCase
	updateUseCase maintenancetypeusecase.UpdateMaintenanceTypeUseCase
	deleteUseCase maintenancetypeusecase.DeleteMaintenanceTypeUseCase
}

func NewMaintenanceTypeHandler(
	createUseCase maintenancetypeusecase.CreateMaintenanceTypeUseCase,
	listUseCase maintenancetypeusecase.ListMaintenanceTypesUseCase,
	updateUseCase maintenancetypeusecase.UpdateMaintenanceTypeUseCase,
	deleteUseCase maintenancetypeusecase.DeleteMaintenanceTypeUseCase,
) MaintenanceTypeHandler {
	return &maintenanceTypeHandler{
		createUseCase: createUseCase,
		listUseCase:   listUseCase,
		updateUseCase: updateUseCase,
		deleteUseCase: deleteUseCase,
	}
}

// Create registers a new maintenance type.
//
//	@Summary	Register a new maintenance type
//	@Tags		maintenance types
//	@Accept		json
//	@Param		maintenance_type	body	dto.CreateMaintenanceTypeRequestDTO	true	"Maintenance type data"
//	@Success	201
//	@Failure	400	{object}	domainerrors.CustomError
//	@Failure	500	{object}	domainerrors.CustomError
//	@Router		/maintenance-types [post]
func (h *maintenanceTypeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateMaintenanceTypeRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	if err := h.createUseCase.Execute(r.Context(), &req); err != nil {
		respondError(w, err)
		return
	}

	respond(w, http.StatusCreated, nil)
}

// List returns all registered maintenance types.
//
//	@Summary	List all maintenance types
//	@Tags		maintenance types
//	@Produce	json
//	@Success	200	{object}	dto.ListMaintenanceTypesResponseDTO
//	@Failure	500	{object}	domainerrors.CustomError
//	@Router		/maintenance-types [get]
func (h *maintenanceTypeHandler) List(w http.ResponseWriter, r *http.Request) {
	resp, err := h.listUseCase.Execute(r.Context())
	if err != nil {
		respondError(w, err)
		return
	}
	respond(w, http.StatusOK, resp)
}

// Update updates an existing maintenance type.
//
//	@Summary	Update a maintenance type
//	@Tags		maintenance types
//	@Accept		json
//	@Param		id				path	int									true	"Maintenance type ID"
//	@Param		maintenance_type	body	dto.UpdateMaintenanceTypeRequestDTO	true	"Updated data"
//	@Success	204
//	@Failure	400	{object}	domainerrors.CustomError
//	@Failure	404	{object}	domainerrors.CustomError
//	@Failure	500	{object}	domainerrors.CustomError
//	@Router		/maintenance-types/{id} [put]
func (h *maintenanceTypeHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	var req dto.UpdateMaintenanceTypeRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	if err := h.updateUseCase.Execute(r.Context(), id, &req); err != nil {
		respondError(w, err)
		return
	}

	respond(w, http.StatusNoContent, nil)
}

// Delete removes a maintenance type.
//
//	@Summary	Remove a maintenance type
//	@Tags		maintenance types
//	@Param		id	path	int	true	"Maintenance type ID"
//	@Success	204
//	@Failure	404	{object}	domainerrors.CustomError
//	@Failure	500	{object}	domainerrors.CustomError
//	@Router		/maintenance-types/{id} [delete]
func (h *maintenanceTypeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	if err := h.deleteUseCase.Execute(r.Context(), id); err != nil {
		respondError(w, err)
		return
	}

	respond(w, http.StatusNoContent, nil)
}
