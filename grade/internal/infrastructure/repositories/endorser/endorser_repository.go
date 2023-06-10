package recognizer

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/recognizer"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewRecognizerRepository(session infrastructure.DbSession) RecognizerRepository {
	return RecognizerRepository{
		session: session,
	}
}

type RecognizerRepository struct {
	session infrastructure.DbSession
}

func (r RecognizerRepository) Insert(obj recognizer.Recognizer) error {
	q := RecognizerInsertQuery{}
	obj.Export(&q)
	result, err := q.Execute(r.session)
	if err != nil {
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows != 0 {
		return fmt.Errorf("wrong rows affected: %d", affectedRows)
	}
	return nil
}
