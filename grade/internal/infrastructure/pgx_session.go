package infrastructure

import (
	"github.com/jackc/pgx/v4"

	"github.com/emacsway/grade/grade/internal/domain/seedwork"
)

func NewPgxSession(db pgx.Conn) *PgxSession {
	return &PgxSession{
		db: db,
	}
}

type PgxSession struct {
	db pgx.Conn
}

func (s *PgxSession) Atomic(func(session seedwork.Session) error) error {
	return nil
}
