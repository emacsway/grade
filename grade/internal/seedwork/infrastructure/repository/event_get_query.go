package repository

import (
	"github.com/emacsway/grade/grade/internal/seedwork/domain/aggregate"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
)

type EventReconstitutor func(
	streamId StreamId,
	streamPosition uint,
	eventType string,
	eventVersion uint,
	payload []byte,
	metadata []byte,
) (aggregate.PersistentDomainEvent, error)

type EventGetQuery struct {
	StreamId           StreamId
	SincePosition      uint
	EventReconstitutor EventReconstitutor
}

func (q EventGetQuery) sql() string {
	return `
		SELECT
		    stream_position, event_type, event_version, payload, metadata
		FROM
			event_log
		WHERE
			tenant_id=$1 AND stream_type=$2 AND stream_id=$3 AND stream_position > $4
		ORDER BY
			tenant_id, stream_type, stream_id, stream_position`
}
func (q EventGetQuery) params() []any {
	return []any{q.StreamId.TenantId(), q.StreamId.StreamType(), q.StreamId.StreamId(), q.SincePosition}
}
func (q *EventGetQuery) Stream(s session.DbSessionQuerier) ([]aggregate.PersistentDomainEvent, error) {
	stream := []aggregate.PersistentDomainEvent{}
	rows, err := s.Query(q.sql(), q.params()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var streamPosition uint
		var eventType string
		var eventVersion uint
		var payload []byte
		var metadata []byte
		err := rows.Scan(&streamPosition, &eventType, &eventVersion, &payload, &metadata)
		if err != nil {
			return nil, err
		}
		event, err := q.EventReconstitutor(q.StreamId, streamPosition, eventType, eventVersion, payload, metadata)
		if err != nil {
			return nil, err
		}
		stream = append(stream, event)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return stream, nil
}
