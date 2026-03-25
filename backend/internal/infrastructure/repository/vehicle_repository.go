package repository

import (
	"context"
	"database/sql"

	"github.com/gaaandrade/car-log/internal/domain/entity"
	domainerrors "github.com/gaaandrade/car-log/internal/domain/errors"
)

type VehicleRepository struct {
	db *sql.DB
}

func NewVehicleRepository(db *sql.DB) *VehicleRepository {
	return &VehicleRepository{db: db}
}

func (r *VehicleRepository) Create(ctx context.Context, v *entity.Vehicle) error {
	res, err := r.db.ExecContext(ctx, `
		INSERT INTO vehicles (name, brand, model, year, plate, current_km)
		VALUES (?, ?, ?, ?, ?, ?)
	`, v.Name, v.Brand, v.Model, v.Year, v.Plate, v.CurrentKM)
	if err != nil {
		return err
	}
	v.ID, err = res.LastInsertId()
	return err
}

func (r *VehicleRepository) FindAll(ctx context.Context) ([]*entity.Vehicle, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, brand, model, year, plate, current_km, created_at, updated_at
		FROM vehicles
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []*entity.Vehicle
	for rows.Next() {
		v := &entity.Vehicle{}
		if err := rows.Scan(&v.ID, &v.Name, &v.Brand, &v.Model, &v.Year, &v.Plate, &v.CurrentKM, &v.CreatedAt, &v.UpdatedAt); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, v)
	}
	return vehicles, rows.Err()
}

func (r *VehicleRepository) FindByID(ctx context.Context, id int64) (*entity.Vehicle, error) {
	v := &entity.Vehicle{}
	err := r.db.QueryRowContext(ctx, `
		SELECT id, name, brand, model, year, plate, current_km, created_at, updated_at
		FROM vehicles
		WHERE id = ?
	`, id).Scan(&v.ID, &v.Name, &v.Brand, &v.Model, &v.Year, &v.Plate, &v.CurrentKM, &v.CreatedAt, &v.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return v, err
}

func (r *VehicleRepository) Update(ctx context.Context, v *entity.Vehicle) error {
	result, err := r.db.ExecContext(ctx, `
		UPDATE vehicles
		SET name=?, brand=?, model=?, year=?, plate=?, current_km=?, updated_at=datetime('now')
		WHERE id=?
	`, v.Name, v.Brand, v.Model, v.Year, v.Plate, v.CurrentKM, v.ID)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domainerrors.NewNotFoundError("Vehicle not found")
	}
	return nil
}

func (r *VehicleRepository) Delete(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM vehicles WHERE id=?`, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domainerrors.NewNotFoundError("Vehicle not found")
	}
	return nil
}

func (r *VehicleRepository) UpdateKM(ctx context.Context, id int64, km int) error {
	result, err := r.db.ExecContext(ctx, `
		UPDATE vehicles SET current_km=?, updated_at=datetime('now') WHERE id=?
	`, km, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domainerrors.NewNotFoundError("Vehicle not found")
	}
	return nil
}
