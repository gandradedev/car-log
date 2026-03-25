package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	vehicleusecase "github.com/gaaandrade/car-log/internal/application/usecase/vehicle"
	"github.com/gaaandrade/car-log/internal/application/usecase/vehicle/dto"
	domainerrors "github.com/gaaandrade/car-log/internal/domain/errors"
)

// VehicleHandler defines the HTTP contract for vehicle operations.
type VehicleHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	UpdateKM(w http.ResponseWriter, r *http.Request)
}

type vehicleHandler struct {
	createUseCase   vehicleusecase.CreateVehicleUseCase
	listUseCase     vehicleusecase.ListVehiclesUseCase
	getUseCase      vehicleusecase.GetVehicleUseCase
	updateUseCase   vehicleusecase.UpdateVehicleUseCase
	deleteUseCase   vehicleusecase.DeleteVehicleUseCase
	updateKMUseCase vehicleusecase.UpdateVehicleKMUseCase
}

func NewVehicleHandler(
	createUseCase vehicleusecase.CreateVehicleUseCase,
	listUseCase vehicleusecase.ListVehiclesUseCase,
	getUseCase vehicleusecase.GetVehicleUseCase,
	updateUseCase vehicleusecase.UpdateVehicleUseCase,
	deleteUseCase vehicleusecase.DeleteVehicleUseCase,
	updateKMUseCase vehicleusecase.UpdateVehicleKMUseCase,
) VehicleHandler {
	return &vehicleHandler{
		createUseCase:   createUseCase,
		listUseCase:     listUseCase,
		getUseCase:      getUseCase,
		updateUseCase:   updateUseCase,
		deleteUseCase:   deleteUseCase,
		updateKMUseCase: updateKMUseCase,
	}
}

func respond(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		json.NewEncoder(w).Encode(v)
	}
}

func respondError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if domainerrors.IsNotFoundError(err) {
		status = http.StatusNotFound
	} else if domainerrors.IsValidationError(err) || domainerrors.IsAlreadyExistsError(err) {
		status = http.StatusBadRequest
	}

	ce := &domainerrors.CustomError{}
	if domainerrors.IsCustomError(err, ce) {
		respond(w, status, ce)
		return
	}
	respond(w, status, map[string]string{"error": err.Error()})
}

func parseID(r *http.Request) (int64, error) {
	return strconv.ParseInt(r.PathValue("id"), 10, 64)
}

// Create registers a new vehicle.
//
//	@Summary	Register a new vehicle
//	@Tags		vehicles
//	@Accept		json
//	@Param		vehicle	body	dto.CreateVehicleRequestDTO	true	"Vehicle data"
//	@Success	201
//	@Failure	400	{object}	domainerrors.CustomError
//	@Failure	500	{object}	domainerrors.CustomError
//	@Router		/vehicles [post]
func (h *vehicleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateVehicleRequestDTO
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

// List returns all registered vehicles.
//
//	@Summary	List all vehicles
//	@Tags		vehicles
//	@Produce	json
//	@Success	200	{object}	dto.ListVehiclesResponseDTO
//	@Failure	500	{object}	domainerrors.CustomError
//	@Router		/vehicles [get]
func (h *vehicleHandler) List(w http.ResponseWriter, r *http.Request) {
	resp, err := h.listUseCase.Execute(r.Context())
	if err != nil {
		respondError(w, err)
		return
	}
	respond(w, http.StatusOK, resp)
}

// GetByID fetches a vehicle by ID.
//
//	@Summary	Fetch a vehicle by ID
//	@Tags		vehicles
//	@Produce	json
//	@Param		id	path		int	true	"Vehicle ID"
//	@Success	200	{object}	dto.GetVehicleResponseDTO
//	@Failure	404	{object}	domainerrors.CustomError
//	@Failure	500	{object}	domainerrors.CustomError
//	@Router		/vehicles/{id} [get]
func (h *vehicleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	resp, err := h.getUseCase.Execute(r.Context(), id)
	if err != nil {
		respondError(w, err)
		return
	}

	respond(w, http.StatusOK, resp)
}

// Update updates an existing vehicle's data.
//
//	@Summary	Update a vehicle
//	@Tags		vehicles
//	@Accept		json
//	@Param		id		path	int							true	"Vehicle ID"
//	@Param		vehicle	body	dto.UpdateVehicleRequestDTO	true	"Updated data"
//	@Success	204
//	@Failure	400	{object}	domainerrors.CustomError
//	@Failure	404	{object}	domainerrors.CustomError
//	@Failure	500	{object}	domainerrors.CustomError
//	@Router		/vehicles/{id} [put]
func (h *vehicleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	var req dto.UpdateVehicleRequestDTO
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

// Delete removes a vehicle.
//
//	@Summary	Remove a vehicle
//	@Tags		vehicles
//	@Param		id	path	int	true	"Vehicle ID"
//	@Success	204
//	@Failure	404	{object}	domainerrors.CustomError
//	@Failure	500	{object}	domainerrors.CustomError
//	@Router		/vehicles/{id} [delete]
func (h *vehicleHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

// UpdateKM updates the current KM of a vehicle.
//
//	@Summary	Update vehicle km
//	@Tags		vehicles
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int							true	"Vehicle ID"
//	@Param		body	body		dto.UpdateVehicleKMRequestDTO	true	"KM data"
//	@Success	200		{object}	dto.GetVehicleResponseDTO
//	@Failure	400		{object}	domainerrors.CustomError
//	@Failure	404		{object}	domainerrors.CustomError
//	@Failure	500		{object}	domainerrors.CustomError
//	@Router		/vehicles/{id}/km [patch]
func (h *vehicleHandler) UpdateKM(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}

	var req dto.UpdateVehicleKMRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	resp, err := h.updateKMUseCase.Execute(r.Context(), id, &req)
	if err != nil {
		respondError(w, err)
		return
	}

	respond(w, http.StatusOK, resp)
}
