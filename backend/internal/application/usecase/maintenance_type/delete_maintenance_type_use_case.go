package maintenancetype

import (
	"context"

	domainrepo "github.com/gaaandrade/car-log/internal/domain/repository"
)

type DeleteMaintenanceTypeUseCase interface {
	Execute(ctx context.Context, id int64) error
}

type deleteMaintenanceTypeUseCase struct {
	repo domainrepo.MaintenanceTypeRepository
}

func NewDeleteMaintenanceTypeUseCase(repo domainrepo.MaintenanceTypeRepository) DeleteMaintenanceTypeUseCase {
	return &deleteMaintenanceTypeUseCase{repo: repo}
}

func (uc *deleteMaintenanceTypeUseCase) Execute(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
