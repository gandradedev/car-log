package vehicle

import (
	"context"

	"github.com/gaaandrade/car-log/internal/application/usecase/vehicle/dto"
	domainrepo "github.com/gaaandrade/car-log/internal/domain/repository"
)

type CreateVehicleUseCase interface {
	Execute(ctx context.Context, req *dto.CreateVehicleRequestDTO) error
}

type createVehicleUseCase struct {
	repo domainrepo.VehicleRepository
}

func NewCreateVehicleUseCase(repo domainrepo.VehicleRepository) CreateVehicleUseCase {
	return &createVehicleUseCase{repo: repo}
}

func (uc *createVehicleUseCase) Execute(ctx context.Context, req *dto.CreateVehicleRequestDTO) error {
	v, err := req.ToEntity()
	if err != nil {
		return err
	}
	return uc.repo.Create(ctx, v)
}
