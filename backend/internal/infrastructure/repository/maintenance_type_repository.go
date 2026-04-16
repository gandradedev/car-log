package repository

import (
	"context"
	"database/sql"

	"github.com/gaaandrade/car-log/internal/domain/entity"
	domainerrors "github.com/gaaandrade/car-log/internal/domain/errors"
)

type MaintenanceTypeRepository struct {
	db *sql.DB
}

func NewMaintenanceTypeRepository(db *sql.DB) *MaintenanceTypeRepository {
	return &MaintenanceTypeRepository{db: db}
}

func (r *MaintenanceTypeRepository) Create(ctx context.Context, mt *entity.MaintenanceType) error {
	res, err := r.db.ExecContext(ctx, `
		INSERT INTO maintenance_types (name, interval_km, interval_months)
		VALUES (?, ?, ?)
	`, mt.Name, mt.IntervalKM, mt.IntervalMonths)
	if err != nil {
		return err
	}
	mt.ID, err = res.LastInsertId()
	return err
}

func (r *MaintenanceTypeRepository) FindAll(ctx context.Context) ([]*entity.MaintenanceType, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, interval_km, interval_months, created_at, updated_at
		FROM maintenance_types
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []*entity.MaintenanceType
	for rows.Next() {
		mt := &entity.MaintenanceType{}
		if err := rows.Scan(&mt.ID, &mt.Name, &mt.IntervalKM, &mt.IntervalMonths, &mt.CreatedAt, &mt.UpdatedAt); err != nil {
			return nil, err
		}
		types = append(types, mt)
	}
	return types, rows.Err()
}

func (r *MaintenanceTypeRepository) FindByID(ctx context.Context, id int64) (*entity.MaintenanceType, error) {
	mt := &entity.MaintenanceType{}
	err := r.db.QueryRowContext(ctx, `
		SELECT id, name, interval_km, interval_months, created_at, updated_at
		FROM maintenance_types
		WHERE id = ?
	`, id).Scan(&mt.ID, &mt.Name, &mt.IntervalKM, &mt.IntervalMonths, &mt.CreatedAt, &mt.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return mt, err
}

func (r *MaintenanceTypeRepository) Update(ctx context.Context, mt *entity.MaintenanceType) error {
	result, err := r.db.ExecContext(ctx, `
		UPDATE maintenance_types
		SET name=?, interval_km=?, interval_months=?, updated_at=datetime('now')
		WHERE id=?
	`, mt.Name, mt.IntervalKM, mt.IntervalMonths, mt.ID)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domainerrors.NewNotFoundError("Maintenance type not found")
	}
	return nil
}

func (r *MaintenanceTypeRepository) Delete(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM maintenance_types WHERE id=?`, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domainerrors.NewNotFoundError("Maintenance type not found")
	}
	return nil
}
