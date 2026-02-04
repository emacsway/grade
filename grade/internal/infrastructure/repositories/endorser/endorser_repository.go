package endorser

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/endorser"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/endorser/queries"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

func NewEndorserRepository() *EndorserRepository {
	return &EndorserRepository{}
}

type EndorserRepository struct{}

func (r *EndorserRepository) Insert(s session.Session, agg *endorser.Endorser) error {
	q := queries.EndorserInsertQuery{}
	agg.Export(&q)
	result, err := q.Evaluate(s)
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

func (r *EndorserRepository) Get(s session.Session, id memberVal.MemberId) (*endorser.Endorser, error) {
	q := queries.EndorserGetQuery{Id: id}
	return q.Get(s)
}
