package repository

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

type EventQueryFactory func(aggregate.PersistentDomainEvent) session.EventSourcedQueryEvaluator

func NewEventStore(currentSession session.DbSession, streamType string, eventQuery EventQueryFactory) *EventStore {
	return &EventStore{
		session:    currentSession,
		streamType: streamType,
		eventQuery: eventQuery,
	}
}

type EventStore struct {
	session    session.DbSession
	streamType string
	eventQuery EventQueryFactory
}

func (r EventStore) NewStreamId(
	tenantId int,
	streamId string,
) (StreamId, error) {
	return NewStreamId(tenantId, r.streamType, streamId)
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
