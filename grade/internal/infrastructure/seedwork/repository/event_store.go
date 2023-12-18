package repository

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type EventQueryFactory func(aggregate.PersistentDomainEvent) infrastructure.EventSourcedQueryEvaluator

func NewEventStore(session infrastructure.DbSession, streamType string, eventQuery EventQueryFactory) *EventStore {
	return &EventStore{
		session:    session,
		streamType: streamType,
		eventQuery: eventQuery,
	}
}

type EventStore struct {
	session    infrastructure.DbSession
	streamType string
	eventQuery EventQueryFactory
}

func (r *EventStore) Save(
	agg aggregate.DomainEventAccessor[aggregate.PersistentDomainEvent],
	eventMeta aggregate.EventMeta,
) error {
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
