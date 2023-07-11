package db

import (
	"database/sql"
	"github.com/pkg/errors"
)

func NewDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize DB")
	}

	err = Migrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
