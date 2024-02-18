package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/artifact/events"
	artifactVal "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/artifact/queries"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/repository"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

func NewArtifactRepository(currentSession session.DbSession) *ArtifactRepository {
	return &ArtifactRepository{
		session:    currentSession,
		eventStore: repository.NewEventStore(currentSession, "Artifact", eventQuery),
	}
}

type ArtifactRepository struct {
	session    session.DbSession
	eventStore *repository.EventStore
}

func (r *ArtifactRepository) Insert(agg *artifact.Artifact, eventMeta aggregate.EventMeta) error {
	return r.eventStore.Save(agg, eventMeta)
}

func (r *ArtifactRepository) NextId(tenantId tenantVal.TenantId) (artifactVal.ArtifactId, error) {
	q := queries.ArtifactNextIdGetQuery{TenantId: tenantId}
	return q.Get(r.session)
}

func eventQuery(iEvent aggregate.PersistentDomainEvent) (q session.EventSourcedQueryEvaluator) {
	switch event := iEvent.(type) {
	case *events.ArtifactProposed:
		q = queries.NewArtifactProposedQuery(event)
	}
	return q
}
