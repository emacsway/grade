package testutils

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewTestDb() (*sql.DB, error) {
	var db_username string = getEnv("DB_USERNAME", "devel")
	var db_password string = getEnv("DB_PASSWORD", "devel")
	var db_host string = getEnv("DB_HOST", "localhost")
	var db_port string = getEnv("DB_PORT", "5432")
	var db_basename string = getEnv("DB_DATABASE", "devel_grade")

	// https://github.com/jackc/pgx/wiki/Getting-started-with-pgx-through-database-sql
	// postgres://username:password@host:port/base_name
	return sql.Open("pgx", "postgres://"+db_username+":"+db_password+"@"+db_host+":"+db_port+"/"+db_basename)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
