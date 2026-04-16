package repository

import (
	"context"

	"github.com/gaaandrade/car-log/internal/domain/entity"
)

// MaintenanceTypeRepository defines the persistence contract for maintenance types.
// The interface lives in the domain to ensure infrastructure depends on the domain, not the other way around.
type MaintenanceTypeRepository interface {
	Create(ctx context.Context, mt *entity.MaintenanceType) error
	FindAll(ctx context.Context) ([]*entity.MaintenanceType, error)
	FindByID(ctx context.Context, id int64) (*entity.MaintenanceType, error)
	Update(ctx context.Context, mt *entity.MaintenanceType) error
	Delete(ctx context.Context, id int64) error
}
