package endorser

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/endorser"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/endorser/queries"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

func NewEndorserRepository(currentSession session.DbSession) *EndorserRepository {
	return &EndorserRepository{
		session: currentSession,
	}
}

type EndorserRepository struct {
	session session.DbSession
}

func (r EndorserRepository) Insert(agg *endorser.Endorser) error {
	q := queries.EndorserInsertQuery{}
	agg.Export(&q)
	result, err := q.Evaluate(r.session)
	if err != nil {
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows != 1 {
		return fmt.Errorf("wrong rows affected: %d", affectedRows)
	}
	return nil
}

func (r *EndorserRepository) Get(id memberVal.MemberId) (*endorser.Endorser, error) {
	q := queries.EndorserGetQuery{Id: id}
	return q.Get(r.session)
}
