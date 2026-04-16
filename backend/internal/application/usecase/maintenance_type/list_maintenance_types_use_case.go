package maintenancetype

import (
	"context"

	"github.com/gaaandrade/car-log/internal/application/usecase/maintenance_type/dto"
	domainrepo "github.com/gaaandrade/car-log/internal/domain/repository"
)

type ListMaintenanceTypesUseCase interface {
	Execute(ctx context.Context) (*dto.ListMaintenanceTypesResponseDTO, error)
}

type listMaintenanceTypesUseCase struct {
	repo domainrepo.MaintenanceTypeRepository
}

func NewListMaintenanceTypesUseCase(repo domainrepo.MaintenanceTypeRepository) ListMaintenanceTypesUseCase {
	return &listMaintenanceTypesUseCase{repo: repo}
}

func (uc *listMaintenanceTypesUseCase) Execute(ctx context.Context) (*dto.ListMaintenanceTypesResponseDTO, error) {
	types, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := &dto.ListMaintenanceTypesResponseDTO{
		Total:            len(types),
		MaintenanceTypes: make([]dto.GetMaintenanceTypeResponseDTO, 0, len(types)),
	}
	for _, mt := range types {
		resp.MaintenanceTypes = append(resp.MaintenanceTypes, *(&dto.GetMaintenanceTypeResponseDTO{}).FromEntity(mt))
	}
	return resp, nil
}
