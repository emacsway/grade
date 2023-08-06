package artifact

import (
	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/artifact/events"
	artifactVal "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/artifact/queries"
)

func NewArtifactRepository(session infrastructure.DbSession) *ArtifactRepository {
	return &ArtifactRepository{
		session:    session,
		streamType: "Artifact",
	}
}

type ArtifactRepository struct {
	session    infrastructure.DbSession
	streamType string
}

func (r *ArtifactRepository) Insert(agg *artifact.Artifact, eventMeta aggregate.EventMeta) error {
	return r.save(agg, eventMeta)
}

func (r *ArtifactRepository) save(agg *artifact.Artifact, eventMeta aggregate.EventMeta) error {
	pendingEvents := agg.PendingDomainEvents()
	for _, iEvent := range pendingEvents {
		iEvent.SetEventMeta(eventMeta)
		q := r.eventQuery(iEvent)
		q.SetStreamType(r.streamType)
		_, err := q.Evaluate(r.session)
		if err != nil {
			return err
		}
	}
	agg.ClearPendingDomainEvents()
	return nil
}
func (r ArtifactRepository) eventQuery(iEvent aggregate.PersistentDomainEvent) (q infrastructure.EventSourcedQueryEvaluator) {
	switch event := iEvent.(type) {
	case *events.ArtifactProposed:
		q = &queries.ArtifactProposedQuery{}
		qt := q.(events.ArtifactProposedExporterSetter)
		event.Export(qt)
	}
	return q
}

func (r *ArtifactRepository) NextId(tenantId tenantVal.TenantId) (artifactVal.TenantArtifactId, error) {
	q := queries.ArtifactNextIdGetQuery{TenantId: tenantId}
	return q.Get(r.session)
}
