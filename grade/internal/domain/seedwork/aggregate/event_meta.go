package aggregate

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

// The source of this data is application layer.

func NewEventMeta(
	eventId uuid.Uuid,
	correlationId uuid.Uuid,
	causationId uuid.Uuid,
	causalDependencies []CausalDependency,
	// occurredAt is the time of taking a slice of the state of the streams,
	// i.e. the time of obtaining the vector clock.
	// Therefore, it is the same for all aggregate events at the time of saving.
	occurredAt time.Time,
) EventMeta {
	return EventMeta{
		eventId:            eventId,
		correlationId:      correlationId,
		causationId:        causationId,
		causalDependencies: causalDependencies,
		occurredAt:         occurredAt,
	}
}

type EventMeta struct {
	eventId            uuid.Uuid
	correlationId      uuid.Uuid
	causationId        uuid.Uuid
	causalDependencies []CausalDependency
	occurredAt         time.Time
}

func (m EventMeta) EventId() uuid.Uuid {
	return m.eventId
}

func (m EventMeta) CorrelationId() uuid.Uuid {
	return m.correlationId
}

func (m EventMeta) CausationId() uuid.Uuid {
	return m.causationId
}

func (m EventMeta) CausalDependencies() []CausalDependency {
	return m.causalDependencies
}

func (m EventMeta) OccurredAt() time.Time {
	return m.occurredAt
}

func (m EventMeta) Spawn(eventId uuid.Uuid) EventMeta {
	n := m
	n.eventId = eventId
	return n
}

func (m EventMeta) Export(ex EventMetaExporterSetter) {
	ex.SetEventId(m.eventId)
	ex.SetCorrelationId(m.correlationId)
	ex.SetCausationId(m.causationId)
	for i := range m.causalDependencies {
		ex.AddCausalDependency(m.causalDependencies[i])
	}
	ex.SetOccurredAt(m.occurredAt)
}

type EventMetaExporterSetter interface {
	SetEventId(uuid.Uuid)
	SetCorrelationId(uuid.Uuid)
	SetCausationId(uuid.Uuid)
	AddCausalDependency(CausalDependency)
	SetOccurredAt(time.Time)
}

type EventMetaExporter struct {
	EventId            uuid.Uuid
	CorrelationId      uuid.Uuid
	CausationId        uuid.Uuid
	CausalDependencies []CausalDependencyExporter
	OccurredAt         time.Time
}

func (ex *EventMetaExporter) SetEventId(val uuid.Uuid) {
	ex.EventId = val
}

func (ex *EventMetaExporter) SetCorrelationId(val uuid.Uuid) {
	ex.CorrelationId = val
}

func (ex *EventMetaExporter) SetCausationId(val uuid.Uuid) {
	ex.CausationId = val
}

func (ex *EventMetaExporter) AddCausalDependency(val CausalDependency) {
	var causalDependencyExp CausalDependencyExporter
	val.Export(&causalDependencyExp)
	ex.CausalDependencies = append(ex.CausalDependencies, causalDependencyExp)
}

func (ex *EventMetaExporter) SetOccurredAt(val time.Time) {
	ex.OccurredAt = val
}
