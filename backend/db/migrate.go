package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/pkg/errors"
	"path/filepath"
)

var migrationsSource string = "./migrations"

func Migrate(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return errors.Wrap(err, "error creating sqlite3 driver")
	}

	path, err := filepath.Abs(migrationsSource)
	if err != nil {
		return errors.Wrap(err, "Error creating path to schema migration")
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", path),
		"sqlite3",
		driver,
	)
	if err != nil {
		return errors.Wrap(err, "Error initiating database migration")
	}

	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "Failed to run schema migration(s)")
	}
	return nil
}
