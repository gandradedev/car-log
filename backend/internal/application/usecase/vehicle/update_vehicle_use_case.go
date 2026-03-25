package vehicle

import (
	"context"

	"github.com/gaaandrade/car-log/internal/application/usecase/vehicle/dto"
	domainrepo "github.com/gaaandrade/car-log/internal/domain/repository"
)

type UpdateVehicleUseCase interface {
	Execute(ctx context.Context, id int64, req *dto.UpdateVehicleRequestDTO) error
}

type updateVehicleUseCase struct {
	repo domainrepo.VehicleRepository
}

func NewUpdateVehicleUseCase(repo domainrepo.VehicleRepository) UpdateVehicleUseCase {
	return &updateVehicleUseCase{repo: repo}
}

func (uc *updateVehicleUseCase) Execute(ctx context.Context, id int64, req *dto.UpdateVehicleRequestDTO) error {
	v, err := req.ToEntity(id)
	if err != nil {
		return err
	}
	return uc.repo.Update(ctx, v)
}
