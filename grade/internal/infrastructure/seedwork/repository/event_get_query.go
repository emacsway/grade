package repository

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
)

type EventGetQuery struct {
	StreamId      StreamId
	SincePosition uint
}

func (q EventGetQuery) sql() string {
	return `
		SELECT
		    tenant_id, stream_type, stream_id, stream_position, event_type, event_version, payload, metadata
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
		err := rows.Scan() // TODO: implement me
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return stream, nil
}
