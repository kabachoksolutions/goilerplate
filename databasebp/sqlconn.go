package databasebp

import (
	"database/sql"
	"fmt"
)

const (
	PostgresDriver = "postgres"
)

func createSQLConn(uri, env, driver string) (*sql.DB, error) {
	if env == "development" {
		uri = uri + "?sslmode=disable"
	}

	db, err := sql.Open(driver, uri)
	if err != nil {
		return nil, fmt.Errorf("databasebp: failed to open sql conn: %w", err)
	}

	return db, nil
}
