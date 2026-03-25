package vehicle

import (
	"context"

	domainrepo "github.com/gaaandrade/car-log/internal/domain/repository"
)

type DeleteVehicleUseCase interface {
	Execute(ctx context.Context, id int64) error
}

type deleteVehicleUseCase struct {
	repo domainrepo.VehicleRepository
}

func NewDeleteVehicleUseCase(repo domainrepo.VehicleRepository) DeleteVehicleUseCase {
	return &deleteVehicleUseCase{repo: repo}
}

func (uc *deleteVehicleUseCase) Execute(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
