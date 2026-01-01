package artifact

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/artifact/events"
	artifactVal "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/artifact/queries"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/aggregate"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/repository"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
)

func NewArtifactRepository(currentSession session.DbSession) *ArtifactRepository {
	return &ArtifactRepository{
		session:    currentSession,
		eventStore: repository.NewEventStore(currentSession, "Artifact", eventToQuery),
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

func (r *ArtifactRepository) Get(id artifactVal.ArtifactId) (*artifact.Artifact, error) {
	idExporter := &artifactVal.ArtifactIdExporter{}
	id.Export(idExporter)
	streamId, err := r.eventStore.NewStreamId(uint(idExporter.TenantId), idExporter.ArtifactId.String())
	if err != nil {
		return nil, err
	}
	q := repository.EventGetQuery{
		StreamId:           streamId,
		EventReconstitutor: rowToEvent,
	}
	stream, err := q.Stream(r.session)
	if err != nil {
		return nil, err
	}
	rec := &artifact.ArtifactReconstitutor{
		PastEvents: stream,
	}
	return rec.Reconstitute()
}

func eventToQuery(iEvent aggregate.PersistentDomainEvent) (q session.EventSourcedQueryEvaluator) {
	switch event := iEvent.(type) {
	case *events.ArtifactProposed:
		q = queries.NewArtifactProposedQuery(event)
	}
	return q
}

func rowToEvent(
	streamId repository.StreamId,
	streamPosition uint,
	eventType string,
	eventVersion uint,
	payload []byte,
	metadata []byte,
) (aggregate.PersistentDomainEvent, error) {
	metaRec := aggregate.EventMetaReconstitutor{}
	err := json.Unmarshal(metadata, &metaRec)
	if err != nil {
		return nil, err
	}
	expectedCase := c{eventType, eventVersion}
	switch expectedCase {
	case c{events.ArtifactProposed{}.EventType(), 1}:
		rec := events.ArtifactProposedReconstitutor{}
		err := json.Unmarshal([]byte(payload), &rec)
		if err != nil {
			return nil, err
		}
		rec.AggregateId.TenantId = streamId.TenantId()
		artifactId, err := strconv.ParseUint(streamId.StreamId(), 10, 0)
		if err != nil {
			return nil, err
		}
		rec.AggregateId.ArtifactId = uint(artifactId)
		rec.AggregateVersion = streamPosition
		rec.EventMeta = metaRec
		return rec.Reconstitute()
	}
	return nil, fmt.Errorf("unknown eventType %s", eventType)
}

type c struct {
	EventType    string
	EventVersion uint
}
