package dto

import "github.com/gaaandrade/car-log/internal/domain/entity"

type CreateVehicleRequestDTO struct {
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	Model     string `json:"model"`
	Year      int    `json:"year"`
	Plate     string `json:"plate"`
	CurrentKM int    `json:"current_km"`
}

func (req *CreateVehicleRequestDTO) ToEntity() (*entity.Vehicle, error) {
	return entity.NewVehicle(req.Name, req.Brand, req.Model, req.Plate, req.Year, req.CurrentKM)
}
