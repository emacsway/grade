package repository

import (
	"encoding/json"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

type EventInsertQuery struct {
	params  [8]any
	payload any
	meta    aggregate.EventMetaExporter
}

func (q EventInsertQuery) sql() string {
	return `
		INSERT INTO event_log
		(tenant_id, stream_type, stream_id, stream_position, event_type, event_version, payload, metadata)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)`
}
func (q *EventInsertQuery) SetTenantId(val tenantVal.TenantId) {
	var v exporters.UintExporter
	val.Export(&v)
	q.params[0] = v
}
func (q *EventInsertQuery) SetStreamType(val string) {
	q.params[1] = val
}
func (q *EventInsertQuery) SetStreamId(val string) {
	q.params[2] = val
}
func (q *EventInsertQuery) SetAggregateVersion(val uint) {
	q.params[3] = val
}
func (q *EventInsertQuery) SetEventType(val string) {
	q.params[4] = val
}
func (q *EventInsertQuery) SetEventMeta(val aggregate.EventMeta) {
	val.Export(&q.meta)
}
func (q *EventInsertQuery) SetPayload(val any) {
	q.payload = val
}
func (q *EventInsertQuery) SetEventVersion(val uint8) {
	q.params[5] = val
}

func (q *EventInsertQuery) Evaluate(s infrastructure.DbSession) (infrastructure.Result, error) {
	payload, err := json.Marshal(q.payload)
	if err != nil {
		return nil, err
	}
	q.params[6] = payload
	meta, err := json.Marshal(q.meta)
	if err != nil {
		return nil, err
	}
	q.params[7] = meta
	return s.Exec(q.sql(), q.params[:]...)
}
