package dto

import "github.com/gaaandrade/car-log/internal/domain/entity"

type UpdateVehicleRequestDTO struct {
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	Model     string `json:"model"`
	Year      int    `json:"year"`
	Plate     string `json:"plate"`
	CurrentKM int    `json:"current_km"`
}

func (req *UpdateVehicleRequestDTO) ToEntity(id int64) (*entity.Vehicle, error) {
	v, err := entity.NewVehicle(req.Name, req.Brand, req.Model, req.Plate, req.Year, req.CurrentKM)
	if err != nil {
		return nil, err
	}
	v.ID = id
	return v, nil
}
