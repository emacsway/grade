package recognizer

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func NewDb() (*sql.DB, error) {
	return sql.Open("pgx", "postgres://devel:devel@localhost:5432/devel")
}
