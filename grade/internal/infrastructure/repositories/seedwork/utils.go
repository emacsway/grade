package seedwork

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var reQuestionPlaceholder = regexp.MustCompile(`'(?:[^']|'')*'|"(?:[^"]|"")*"|(\?)`)
var reDollarPlaceholder = regexp.MustCompile(`'(?:[^']|'')*'|"(?:[^"]|"")*"|(\$\d+)`)

func Rebind(query string) string {
	// See also another example of implementation
	// https://github.com/jmoiron/sqlx/blob/1abdd3dc2a5d5257e3714cee6e8e98835cdb1b2e/bind.go#L60
	var offset = 0
	return reQuestionPlaceholder.ReplaceAllStringFunc(query, func(s string) string {
		if s == "?" {
			offset += 1
			return fmt.Sprintf("$%d", offset)
		}
		return s
	})
}

func RebindReverse(query string) string {
	return reDollarPlaceholder.ReplaceAllStringFunc(query, func(s string) string {
		if s[:1] == "$" {
			return "?"
		}
		return s
	})
}

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