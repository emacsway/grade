package session

/*
 * Some parts of the code based on
 * https://github.com/mongodb/mongo-go-driver/blob/master/mongo/session.go
 * https://github.com/mongodb/mongo-go-driver/blob/master/internal/background_context.go
 */

import (
	"context"
	"strings"

	"database/sql"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/emacsway/grade/grade/internal/application"
)

type sessionKey struct{}

func SessionFromContext(ctx context.Context) DbSession {
	val := ctx.Value(sessionKey{})
	if val == nil {
		return nil
	}

	sess, ok := val.(DbSession)
	if !ok {
		return nil
	}

	return sess
}

func NewSessionContext(ctx context.Context, db *sql.DB) *SessionContext {
	sess := NewPgxSession(db)
	return &SessionContext{
		context.WithValue(ctx, sessionKey{}, sess),
		sess,
	}
}

type SessionContext struct {
	context.Context
	DbSession
}

func (s *SessionContext) Atomic(callback application.SessionContextCallback) error {
	callbackUnclothed := func(dbSession application.Session) error {
		sessionContext := &SessionContext{
			NewBackgroundContext(s.Context),
			dbSession.(DbSession),
		}
		return callback(sessionContext)
	}
	return s.DbSession.Atomic(callbackUnclothed)
}

func NewPgxSession(db *sql.DB) *PgxSession {
	return &PgxSession{
		db:         db,
		dbExecutor: db,
	}
}

type PgxSession struct {
	db         *sql.DB
	dbExecutor DbExecutor
}

func (s *PgxSession) Atomic(callback application.SessionCallback) error {
	// TODO: Add support for SavePoint:
	// https://github.com/golang/go/issues/7898#issuecomment-580080390
	if s.db == nil {
		return errors.New("savePoint is not currently supported")
	}
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "unable to start transaction")
	}
	newSession := &PgxSession{
		dbExecutor: tx,
	}
	err = callback(newSession)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return multierror.Append(err, txErr)
		}
		return err
	}
	if txErr := tx.Commit(); txErr != nil {
		return errors.Wrap(err, "failed to commit tx")
	}
	return nil
}

func (s *PgxSession) Exec(query string, args ...any) (Result, error) {
	if IsAutoincrementInsertQuery(query) {
		return s.insert(query, args...)
	}
	return s.dbExecutor.Exec(query, args...)
}

func (s *PgxSession) insert(query string, args ...any) (Result, error) {
	var id int
	err := s.dbExecutor.QueryRow(query, args...).Scan(&id)
	return LastInsertId(id), err
}

func (s *PgxSession) Query(query string, args ...any) (Rows, error) {
	return s.dbExecutor.Query(query, args...)
}

func (s *PgxSession) QueryRow(query string, args ...any) Row {
	return s.dbExecutor.QueryRow(query, args...)
}

type DbExecutor interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

type LastInsertId int64

func (i LastInsertId) LastInsertId() (int64, error) {
	return int64(i), nil
}

func (i LastInsertId) RowsAffected() (int64, error) {
	return 0, errors.New("RowsAffected is not supported by INSERT command")
}

type RowsAffected int64

func (RowsAffected) LastInsertId() (int64, error) {
	return 0, errors.New("LastInsertId is not supported by this driver")
}

func (v RowsAffected) RowsAffected() (int64, error) {
	return int64(v), nil
}

func IsInsertQuery(query string) bool {
	return strings.TrimSpace(query)[:6] == "INSERT" && !strings.Contains(query, "RETURNING")
}

func IsAutoincrementInsertQuery(query string) bool {
	return strings.TrimSpace(query)[:6] == "INSERT" && strings.Contains(query, "RETURNING")
}

type backgroundContext struct {
	context.Context
	childValuesCtx context.Context
}

// NewBackgroundContext creates a new Context whose behavior matches that of context.Background(), but Value calls are
// forwarded to the provided ctx parameter. If ctx is nil, context.Background() is returned.
func NewBackgroundContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}

	return &backgroundContext{
		Context:        context.Background(),
		childValuesCtx: ctx,
	}
}

func (b *backgroundContext) Value(key interface{}) interface{} {
	return b.childValuesCtx.Value(key)
}
