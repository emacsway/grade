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
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/domain/aggregate"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/infrastructure/repository"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

func NewArtifactRepository() *ArtifactRepository {
	return &ArtifactRepository{
		eventStore: repository.NewEventStore(nil, "Artifact", eventToQuery),
	}
}

type ArtifactRepository struct {
	eventStore *repository.EventStore
}

func (r *ArtifactRepository) Insert(s session.Session, agg *artifact.Artifact, eventMeta aggregate.EventMeta) error {
	return r.eventStore.Save(s, agg, eventMeta)
}

func (r *ArtifactRepository) NextId(s session.Session, tenantId tenantVal.TenantId) (artifactVal.ArtifactId, error) {
	q := queries.ArtifactNextIdGetQuery{TenantId: tenantId}
	return q.Get(s)
}

func (r *ArtifactRepository) Get(s session.Session, id artifactVal.ArtifactId) (*artifact.Artifact, error) {
	idExporter := &artifactVal.ArtifactIdExporter{}
	id.Export(idExporter)
	streamId, err := r.eventStore.NewStreamId(uint(idExporter.TenantId), strconv.FormatUint(uint64(idExporter.ArtifactId), 10))
	if err != nil {
		return nil, err
	}
	q := repository.EventGetQuery{
		StreamId:           streamId,
		EventReconstitutor: rowToEvent,
	}
	stream, err := q.Stream(r.eventStore.MakeReadCodecFactory(), s)
	if err != nil {
		return nil, err
	}
	rec := &artifact.ArtifactReconstitutor{
		PastEvents: stream,
	}
	return rec.Reconstitute()
}

func eventToQuery(iEvent aggregate.PersistentDomainEvent) (q repository.EventSourcedQueryEvaluator) {
	switch event := iEvent.(type) {
	case *events.ArtifactProposed:
		q = queries.NewArtifactProposedQuery(event)
	}
	return q
}

func rowToEvent(
	codec repository.Codec,
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
		err := codec.Decode(payload, &rec)
		if err != nil {
			return nil, err
		}
		rec.AggregateId.TenantId = streamId.TenantId().(uint)
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
