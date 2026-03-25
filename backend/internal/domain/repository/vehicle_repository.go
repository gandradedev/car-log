package repository

import (
	"context"

	"github.com/gaaandrade/car-log/internal/domain/entity"
)

// VehicleRepository defines the persistence contract for vehicles.
// The interface lives in the domain to ensure infrastructure depends on the domain, not the other way around.
type VehicleRepository interface {
	Create(ctx context.Context, v *entity.Vehicle) error
	FindAll(ctx context.Context) ([]*entity.Vehicle, error)
	FindByID(ctx context.Context, id int64) (*entity.Vehicle, error)
	Update(ctx context.Context, v *entity.Vehicle) error
	Delete(ctx context.Context, id int64) error
	UpdateKM(ctx context.Context, id int64, km int) error
}
