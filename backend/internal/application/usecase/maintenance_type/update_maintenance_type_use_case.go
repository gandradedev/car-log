package maintenancetype

import (
	"context"

	"github.com/gaaandrade/car-log/internal/application/usecase/maintenance_type/dto"
	domainrepo "github.com/gaaandrade/car-log/internal/domain/repository"
)

type UpdateMaintenanceTypeUseCase interface {
	Execute(ctx context.Context, id int64, req *dto.UpdateMaintenanceTypeRequestDTO) error
}

type updateMaintenanceTypeUseCase struct {
	repo domainrepo.MaintenanceTypeRepository
}

func NewUpdateMaintenanceTypeUseCase(repo domainrepo.MaintenanceTypeRepository) UpdateMaintenanceTypeUseCase {
	return &updateMaintenanceTypeUseCase{repo: repo}
}

func (uc *updateMaintenanceTypeUseCase) Execute(ctx context.Context, id int64, req *dto.UpdateMaintenanceTypeRequestDTO) error {
	mt, err := req.ToEntity(id)
	if err != nil {
		return err
	}
	return uc.repo.Update(ctx, mt)
}
