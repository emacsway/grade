package endorser

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewEndorserRepository(session infrastructure.DbSession) EndorserRepository {
	return EndorserRepository{
		session: session,
	}
}

type EndorserRepository struct {
	session infrastructure.DbSession
}

func (r EndorserRepository) Insert(obj endorser.Endorser) error {
	q := EndorserInsertQuery{}
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
