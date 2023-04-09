package recognizer

import (
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
