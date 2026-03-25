package dto

import (
	"time"

	"github.com/gaaandrade/car-log/internal/domain/entity"
)

// GetVehicleResponseDTO is the output DTO for a single vehicle.
type GetVehicleResponseDTO struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Brand     string    `json:"brand"`
	Model     string    `json:"model"`
	Year      int       `json:"year"`
	Plate     string    `json:"plate"`
	CurrentKM int       `json:"current_km"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d *GetVehicleResponseDTO) FromEntity(v *entity.Vehicle) *GetVehicleResponseDTO {
	return &GetVehicleResponseDTO{
		ID:        v.ID,
		Name:      v.Name,
		Brand:     v.Brand,
		Model:     v.Model,
		Year:      v.Year,
		Plate:     v.Plate,
		CurrentKM: v.CurrentKM,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}
}
