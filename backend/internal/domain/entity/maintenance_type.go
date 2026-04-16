package entity

import (
	"strings"
	"time"

	domainerrors "github.com/gaaandrade/car-log/internal/domain/errors"
)

// MaintenanceType represents a category of vehicle maintenance (e.g. oil change, tire rotation).
type MaintenanceType struct {
	ID             int64
	Name           string
	IntervalKM     *int
	IntervalMonths *int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// NewMaintenanceType creates and validates a new MaintenanceType.
func NewMaintenanceType(name string, intervalKM, intervalMonths *int) (*MaintenanceType, error) {
	mt := &MaintenanceType{
		Name:           strings.TrimSpace(name),
		IntervalKM:     intervalKM,
		IntervalMonths: intervalMonths,
	}
	if err := mt.Validate(); err != nil {
		return nil, err
	}
	return mt, nil
}

// Validate checks whether the maintenance type has all required fields.
func (mt *MaintenanceType) Validate() error {
	if mt.Name == "" {
		return domainerrors.NewValidationError("name_required", "Name is required", nil)
	}
	if mt.IntervalKM != nil && *mt.IntervalKM <= 0 {
		return domainerrors.NewValidationError("invalid_interval_km", "Interval KM must be greater than 0", nil)
	}
	if mt.IntervalMonths != nil && *mt.IntervalMonths <= 0 {
		return domainerrors.NewValidationError("invalid_interval_months", "Interval months must be greater than 0", nil)
	}
	return nil
}
