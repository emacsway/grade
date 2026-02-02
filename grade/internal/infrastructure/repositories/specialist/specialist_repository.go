package specialist

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/specialist"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/specialist/queries"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

func NewSpecialistRepository(currentSession session.DbSession) *SpecialistRepository {
	return &SpecialistRepository{
		session: currentSession,
	}
}

type SpecialistRepository struct {
	session session.DbSession
}

func (r SpecialistRepository) Insert(agg *specialist.Specialist) error {
	q := queries.SpecialistInsertQuery{}
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

/*
func (r *SpecialistRepository) Get(id memberVal.MemberId) (*specialist.Specialist, error) {
	q := queries.SpecialistGetQuery{Id: id}
	return q.Get(r.session)
}
*/
