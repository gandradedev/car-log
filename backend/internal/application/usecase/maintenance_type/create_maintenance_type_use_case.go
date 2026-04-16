package maintenancetype

import (
	"context"

	"github.com/gaaandrade/car-log/internal/application/usecase/maintenance_type/dto"
	domainrepo "github.com/gaaandrade/car-log/internal/domain/repository"
)

type CreateMaintenanceTypeUseCase interface {
	Execute(ctx context.Context, req *dto.CreateMaintenanceTypeRequestDTO) error
}

type createMaintenanceTypeUseCase struct {
	repo domainrepo.MaintenanceTypeRepository
}

func NewCreateMaintenanceTypeUseCase(repo domainrepo.MaintenanceTypeRepository) CreateMaintenanceTypeUseCase {
	return &createMaintenanceTypeUseCase{repo: repo}
}

func (uc *createMaintenanceTypeUseCase) Execute(ctx context.Context, req *dto.CreateMaintenanceTypeRequestDTO) error {
	mt, err := req.ToEntity()
	if err != nil {
		return err
	}
	return uc.repo.Create(ctx, mt)
}
