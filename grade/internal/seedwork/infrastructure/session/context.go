package session

/*
 * Some parts of the code based on
 * https://github.com/mongodb/mongo-go-driver/blob/master/mongo/session.go
 * https://github.com/mongodb/mongo-go-driver/blob/master/internal/background_context.go
 */

import (
	"context"

	"database/sql"

	"github.com/emacsway/grade/grade/internal/seedwork/application/session"
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

func (s *SessionContext) Atomic(callback session.SessionContextCallback) error {
	callbackUnclothed := func(dbSession session.Session) error {
		sessionContext := &SessionContext{
			NewBackgroundContext(s.Context),
			dbSession.(DbSession),
		}
		return callback(sessionContext)
	}
	return s.DbSession.Atomic(callbackUnclothed)
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
