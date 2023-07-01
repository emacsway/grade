package endorser

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/endorser/queries"
)

func NewEndorserRepository(session infrastructure.DbSession) *EndorserRepository {
	return &EndorserRepository{
		session: session,
	}
}

type EndorserRepository struct {
	session infrastructure.DbSession
}

func (r EndorserRepository) Insert(obj *endorser.Endorser) error {
	q := queries.EndorserInsertQuery{}
	obj.Export(&q)
	result, err := q.Evaluate(r.session)
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
