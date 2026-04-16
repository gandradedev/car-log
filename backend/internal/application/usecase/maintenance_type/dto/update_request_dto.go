package dto

import "github.com/gaaandrade/car-log/internal/domain/entity"

type UpdateMaintenanceTypeRequestDTO struct {
	Name           string `json:"name"`
	IntervalKM     *int   `json:"interval_km"`
	IntervalMonths *int   `json:"interval_months"`
}

func (req *UpdateMaintenanceTypeRequestDTO) ToEntity(id int64) (*entity.MaintenanceType, error) {
	mt, err := entity.NewMaintenanceType(req.Name, req.IntervalKM, req.IntervalMonths)
	if err != nil {
		return nil, err
	}
	mt.ID = id
	return mt, nil
}
