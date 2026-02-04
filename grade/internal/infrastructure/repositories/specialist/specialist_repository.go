package specialist

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/specialist"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/specialist/queries"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

func NewSpecialistRepository() *SpecialistRepository {
	return &SpecialistRepository{}
}

type SpecialistRepository struct{}

func (r *SpecialistRepository) Insert(s session.Session, agg *specialist.Specialist) error {
	q := queries.SpecialistInsertQuery{}
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

/*
func (r *SpecialistRepository) Get(s session.Session, id memberVal.MemberId) (*specialist.Specialist, error) {
	q := queries.SpecialistGetQuery{Id: id}
	return q.Get(s)
}
*/
