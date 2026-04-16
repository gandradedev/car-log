package dto

import (
	"time"

	"github.com/gaaandrade/car-log/internal/domain/entity"
)

// GetMaintenanceTypeResponseDTO is the output DTO for a single maintenance type.
type GetMaintenanceTypeResponseDTO struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	IntervalKM     *int      `json:"interval_km"`
	IntervalMonths *int      `json:"interval_months"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (d *GetMaintenanceTypeResponseDTO) FromEntity(mt *entity.MaintenanceType) *GetMaintenanceTypeResponseDTO {
	return &GetMaintenanceTypeResponseDTO{
		ID:             mt.ID,
		Name:           mt.Name,
		IntervalKM:     mt.IntervalKM,
		IntervalMonths: mt.IntervalMonths,
		CreatedAt:      mt.CreatedAt,
		UpdatedAt:      mt.UpdatedAt,
	}
}
