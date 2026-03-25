package vehicle

import (
	"context"

	"github.com/gaaandrade/car-log/internal/application/usecase/vehicle/dto"
	domainerrors "github.com/gaaandrade/car-log/internal/domain/errors"
	domainrepo "github.com/gaaandrade/car-log/internal/domain/repository"
)

type UpdateVehicleKMUseCase interface {
	Execute(ctx context.Context, id int64, req *dto.UpdateVehicleKMRequestDTO) (*dto.GetVehicleResponseDTO, error)
}

type updateVehicleKMUseCase struct {
	repo domainrepo.VehicleRepository
}

func NewUpdateVehicleKMUseCase(repo domainrepo.VehicleRepository) UpdateVehicleKMUseCase {
	return &updateVehicleKMUseCase{repo: repo}
}

func (uc *updateVehicleKMUseCase) Execute(ctx context.Context, id int64, req *dto.UpdateVehicleKMRequestDTO) (*dto.GetVehicleResponseDTO, error) {
	current, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if current == nil {
		return nil, domainerrors.NewNotFoundError("Vehicle not found")
	}
	if req.CurrentKM < current.CurrentKM {
		return nil, domainerrors.NewValidationError("invalid_km", "Current KM must be >= vehicle's current KM", nil)
	}

	if err := uc.repo.UpdateKM(ctx, id, req.CurrentKM); err != nil {
		return nil, err
	}

	updated, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return (&dto.GetVehicleResponseDTO{}).FromEntity(updated), nil
}
