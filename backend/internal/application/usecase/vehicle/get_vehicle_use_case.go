package vehicle

import (
	"context"

	"github.com/gaaandrade/car-log/internal/application/usecase/vehicle/dto"
	domainerrors "github.com/gaaandrade/car-log/internal/domain/errors"
	domainrepo "github.com/gaaandrade/car-log/internal/domain/repository"
)

type GetVehicleUseCase interface {
	Execute(ctx context.Context, id int64) (*dto.GetVehicleResponseDTO, error)
}

type getVehicleUseCase struct {
	repo domainrepo.VehicleRepository
}

func NewGetVehicleUseCase(repo domainrepo.VehicleRepository) GetVehicleUseCase {
	return &getVehicleUseCase{repo: repo}
}

func (uc *getVehicleUseCase) Execute(ctx context.Context, id int64) (*dto.GetVehicleResponseDTO, error) {
	v, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, domainerrors.NewNotFoundError("Vehicle not found")
	}
	return (&dto.GetVehicleResponseDTO{}).FromEntity(v), nil
}
