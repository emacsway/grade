package testutils

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewTestDb() (*sql.DB, error) {
	// https://github.com/jackc/pgx/wiki/Getting-started-with-pgx-through-database-sql
	return sql.Open("pgx", getEnv("DATABASE_URL", "postgres://devel:devel@localhost:5432/devel_grade"))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
