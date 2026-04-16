package dto

import "github.com/gaaandrade/car-log/internal/domain/entity"

type CreateMaintenanceTypeRequestDTO struct {
	Name           string `json:"name"`
	IntervalKM     *int   `json:"interval_km"`
	IntervalMonths *int   `json:"interval_months"`
}

func (req *CreateMaintenanceTypeRequestDTO) ToEntity() (*entity.MaintenanceType, error) {
	return entity.NewMaintenanceType(req.Name, req.IntervalKM, req.IntervalMonths)
}
