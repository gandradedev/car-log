package vehicle

import (
	"context"

	"github.com/gaaandrade/car-log/internal/application/usecase/vehicle/dto"
	domainrepo "github.com/gaaandrade/car-log/internal/domain/repository"
)

type ListVehiclesUseCase interface {
	Execute(ctx context.Context) (*dto.ListVehiclesResponseDTO, error)
}

type listVehiclesUseCase struct {
	repo domainrepo.VehicleRepository
}

func NewListVehiclesUseCase(repo domainrepo.VehicleRepository) ListVehiclesUseCase {
	return &listVehiclesUseCase{repo: repo}
}

func (uc *listVehiclesUseCase) Execute(ctx context.Context) (*dto.ListVehiclesResponseDTO, error) {
	vehicles, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.GetVehicleResponseDTO, len(vehicles))
	for i, v := range vehicles {
		responses[i] = *(&dto.GetVehicleResponseDTO{}).FromEntity(v)
	}

	return &dto.ListVehiclesResponseDTO{
		Total:    len(vehicles),
		Vehicles: responses,
	}, nil
}
