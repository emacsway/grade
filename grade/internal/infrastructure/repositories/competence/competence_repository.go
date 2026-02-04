package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/competence"
	"github.com/emacsway/grade/grade/internal/domain/competence/events"
	competenceVal "github.com/emacsway/grade/grade/internal/domain/competence/values"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/competence/queries"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/domain/aggregate"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

func NewCompetenceRepository() *CompetenceRepository {
	return &CompetenceRepository{}
}

type CompetenceRepository struct{}

func (r *CompetenceRepository) Insert(s session.Session, agg *competence.Competence) error {
	return r.save(s, agg)
}

func (r *CompetenceRepository) Update(s session.Session, agg *competence.Competence) error {
	q := &queries.OptimisticOfflineLockLockQuery{}
	agg.Export(q)
	q.SetInitialVersion(agg.Version() - uint(len(agg.PendingDomainEvents())))
	_, err := q.Evaluate(s)
	if err != nil {
		return err
	}
	return r.save(s, agg)
}

func (r *CompetenceRepository) save(s session.Session, agg *competence.Competence) error {
	pendingEvents := agg.PendingDomainEvents()
	for _, iEvent := range pendingEvents {
		q := r.eventQuery(iEvent)
		_, err := q.Evaluate(s)
		if err != nil {
			return err
		}
	}
	agg.ClearPendingDomainEvents()
	return nil
}

func (r CompetenceRepository) eventQuery(iEvent aggregate.DomainEvent) (q session.QueryEvaluator) {
	switch event := iEvent.(type) {
	case *events.CompetenceCreated:
		q = &queries.CompetenceCreatedQuery{}
		qt := q.(events.CompetenceCreatedExporterSetter)
		event.Export(qt)
	case *events.NameUpdated:
		q = &queries.NameUpdatedQuery{}
		qt := q.(events.NameUpdatedExporterSetter)
		event.Export(qt)
	}
	return q
}

func (r *CompetenceRepository) Get(s session.Session, id competenceVal.CompetenceId) (*competence.Competence, error) {
	q := queries.CompetenceGetQuery{Id: id}
	return q.Get(s)
}
