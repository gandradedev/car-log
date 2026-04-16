package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite"
)

type Config struct {
	Path              string
	RunStartupScripts bool
}

func New(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite", cfg.Path)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		db.Close()
		return nil, fmt.Errorf("enabling foreign keys: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	if err := migrate(db); err != nil {
		db.Close()
		return nil, err
	}

	if cfg.RunStartupScripts {
		if err := runStartupScripts(db); err != nil {
			db.Close()
			return nil, err
		}
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS vehicles (
			id          INTEGER PRIMARY KEY AUTOINCREMENT,
			name        TEXT    NOT NULL,
			brand       TEXT    NOT NULL,
			model       TEXT    NOT NULL,
			year        INTEGER NOT NULL,
			plate       TEXT    NOT NULL UNIQUE,
			current_km  INTEGER NOT NULL DEFAULT 0,
			created_at  DATETIME NOT NULL DEFAULT (datetime('now')),
			updated_at  DATETIME NOT NULL DEFAULT (datetime('now'))
		)`,

		`CREATE TABLE IF NOT EXISTS maintenance_types (
			id               INTEGER PRIMARY KEY AUTOINCREMENT,
			name             TEXT    NOT NULL UNIQUE,
			interval_km      INTEGER,
			interval_months  INTEGER,
			created_at       DATETIME NOT NULL DEFAULT (datetime('now')),
			updated_at       DATETIME NOT NULL DEFAULT (datetime('now'))
		)`,

		`CREATE TABLE IF NOT EXISTS maintenance_records (
			id                  INTEGER PRIMARY KEY AUTOINCREMENT,
			vehicle_id          INTEGER NOT NULL REFERENCES vehicles(id) ON DELETE CASCADE,
			maintenance_type_id INTEGER NOT NULL REFERENCES maintenance_types(id),
			date                DATE    NOT NULL,
			km_at_maintenance   INTEGER NOT NULL,
			cost                REAL    NOT NULL DEFAULT 0,
			workshop            TEXT,
			notes               TEXT,
			created_at          DATETIME NOT NULL DEFAULT (datetime('now')),
			updated_at          DATETIME NOT NULL DEFAULT (datetime('now'))
		)`,

		`CREATE TABLE IF NOT EXISTS maintenance_schedules (
			id                  INTEGER PRIMARY KEY AUTOINCREMENT,
			vehicle_id          INTEGER NOT NULL REFERENCES vehicles(id) ON DELETE CASCADE,
			maintenance_type_id INTEGER NOT NULL REFERENCES maintenance_types(id),
			due_date            DATE,
			due_km              INTEGER,
			status              TEXT NOT NULL DEFAULT 'pending' CHECK(status IN ('pending', 'completed', 'overdue')),
			created_at          DATETIME NOT NULL DEFAULT (datetime('now')),
			updated_at          DATETIME NOT NULL DEFAULT (datetime('now'))
		)`,
	}

	for _, stmt := range statements {
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("executing migration: %w", err)
		}
	}

	return nil
}

func runStartupScripts(db *sql.DB) error {
	entries, err := os.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("reading migrations directory: %w", err)
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".sql" {
			continue
		}
		content, err := os.ReadFile(filepath.Join("migrations", entry.Name()))
		if err != nil {
			return fmt.Errorf("reading migration file %s: %w", entry.Name(), err)
		}
		if err := execSQL(db, string(content)); err != nil {
			return fmt.Errorf("executing migration file %s: %w", entry.Name(), err)
		}
	}
	return nil
}

func execSQL(db *sql.DB, sql string) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	defer tx.Rollback()

	for _, stmt := range strings.Split(sql, ";") {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		if _, err := tx.Exec(stmt); err != nil {
			return fmt.Errorf("executing statement: %w", err)
		}
	}

	return tx.Commit()
}
