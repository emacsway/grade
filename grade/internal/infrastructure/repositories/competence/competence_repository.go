package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/competence"
	"github.com/emacsway/grade/grade/internal/domain/competence/events"
	competenceVal "github.com/emacsway/grade/grade/internal/domain/competence/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/competence/queries"
)

func NewCompetenceRepository(session infrastructure.DbSession) *CompetenceRepository {
	return &CompetenceRepository{
		session: session,
	}
}

type CompetenceRepository struct {
	session infrastructure.DbSession
}

func (r *CompetenceRepository) Insert(agg *competence.Competence) error {
	return r.save(agg)
}

func (r *CompetenceRepository) Update(agg *competence.Competence) error {
	q := &queries.OptimisticOfflineLockLockQuery{}
	agg.Export(q)
	q.SetInitialVersion(agg.Version() - uint(len(agg.PendingDomainEvents())))
	_, err := q.Evaluate(r.session)
	if err != nil {
		return err
	}
	return r.save(agg)
}

func (r *CompetenceRepository) save(agg *competence.Competence) error {
	pendingEvents := agg.PendingDomainEvents()
	for i := range pendingEvents {
		var q infrastructure.QueryEvaluator

		switch event := pendingEvents[i].(type) {
		case events.CompetenceCreated:
			q = &queries.CompetenceCreatedQuery{}
			qt := q.(events.CompetenceCreatedExporterSetter)
			event.Export(qt)
		case events.NameUpdated:
			q = &queries.NameUpdatedQuery{}
			qt := q.(events.NameUpdatedExporterSetter)
			event.Export(qt)
		}
		_, err := q.Evaluate(r.session)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *CompetenceRepository) Get(id competenceVal.TenantCompetenceId) (*competence.Competence, error) {
	q := queries.CompetenceGetQuery{Id: id}
	return q.Get(r.session)
}
