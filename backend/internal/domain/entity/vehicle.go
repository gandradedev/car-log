package entity

import (
	"strings"
	"time"

	domainerrors "github.com/gaaandrade/car-log/internal/domain/errors"
)

// Vehicle is the core domain entity. It has no JSON tags as it is not exposed directly via HTTP.
type Vehicle struct {
	ID        int64
	Name      string
	Brand     string
	Model     string
	Year      int
	Plate     string
	CurrentKM int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewVehicle creates and validates a new vehicle.
func NewVehicle(name, brand, model, plate string, year, currentKM int) (*Vehicle, error) {
	v := &Vehicle{
		Name:      strings.TrimSpace(name),
		Brand:     strings.TrimSpace(brand),
		Model:     strings.TrimSpace(model),
		Year:      year,
		Plate:     strings.ToUpper(strings.TrimSpace(plate)),
		CurrentKM: currentKM,
	}
	if err := v.Validate(); err != nil {
		return nil, err
	}
	return v, nil
}

// Validate checks whether the vehicle has all required fields.
func (v *Vehicle) Validate() error {
	if v.Name == "" {
		return domainerrors.NewValidationError("name_required", "Name is required", nil)
	}
	if v.Brand == "" {
		return domainerrors.NewValidationError("brand_required", "Brand is required", nil)
	}
	if v.Model == "" {
		return domainerrors.NewValidationError("model_required", "Model is required", nil)
	}
	if v.Year < 1900 || v.Year > time.Now().Year()+1 {
		return domainerrors.NewValidationError("invalid_year", "Year is invalid", nil)
	}
	if v.Plate == "" {
		return domainerrors.NewValidationError("plate_required", "Plate is required", nil)
	}
	if v.CurrentKM < 0 {
		return domainerrors.NewValidationError("invalid_km", "Current KM must be >= 0", nil)
	}
	return nil
}
