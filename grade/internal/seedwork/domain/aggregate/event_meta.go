package aggregate

import (
	"time"

	"github.com/emacsway/grade/grade/internal/seedwork/domain/uuid"
)

// The source of this data is application layer.

func NewEventMeta(
	eventId uuid.Uuid,
	causationId uuid.Uuid,
	correlationId uuid.Uuid,
	causalDependencies []CausalDependency,
	// occurredAt is the time of taking a slice of the state of the streams,
	// i.e. the time of obtaining the vector clock.
	// Therefore, it is the same for all aggregate events at the time of saving.
	occurredAt time.Time,
	reason string,
) (EventMeta, error) {
	return EventMeta{
		eventId:            eventId,
		causationId:        causationId,
		correlationId:      correlationId,
		causalDependencies: causalDependencies,
		occurredAt:         occurredAt,
		reason:             reason,
	}, nil
}

type EventMeta struct {
	eventId            uuid.Uuid
	causationId        uuid.Uuid // CommandId or causation EventId
	correlationId      uuid.Uuid // CommandId
	causalDependencies []CausalDependency
	occurredAt         time.Time
	reason             string
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

func (m EventMeta) Reason() string {
	return m.reason
}

func (m EventMeta) Spawn(eventId uuid.Uuid) EventMeta {
	n := m
	n.causationId = n.eventId
	n.eventId = eventId
	return n
}

func (m EventMeta) Export(ex EventMetaExporterSetter) {
	ex.SetEventId(m.eventId)
	ex.SetCausationId(m.causationId)
	ex.SetCorrelationId(m.correlationId)
	for i := range m.causalDependencies {
		ex.AddCausalDependency(m.causalDependencies[i])
	}
	ex.SetOccurredAt(m.occurredAt)
}

type EventMetaExporterSetter interface {
	SetEventId(uuid.Uuid)
	SetCausationId(uuid.Uuid)
	SetCorrelationId(uuid.Uuid)
	AddCausalDependency(CausalDependency)
	SetOccurredAt(time.Time)
	SetReason(string)
}

type EventMetaExporter struct {
	EventId            uuid.Uuid
	CausationId        uuid.Uuid
	CorrelationId      uuid.Uuid
	CausalDependencies []CausalDependencyExporter
	OccurredAt         time.Time
	Reason             string
}

func (ex *EventMetaExporter) SetEventId(val uuid.Uuid) {
	ex.EventId = val
}

func (ex *EventMetaExporter) SetCausationId(val uuid.Uuid) {
	ex.CausationId = val
}

func (ex *EventMetaExporter) SetCorrelationId(val uuid.Uuid) {
	ex.CorrelationId = val
}

func (ex *EventMetaExporter) AddCausalDependency(val CausalDependency) {
	var causalDependencyExp CausalDependencyExporter
	val.Export(&causalDependencyExp)
	ex.CausalDependencies = append(ex.CausalDependencies, causalDependencyExp)
}

func (ex *EventMetaExporter) SetOccurredAt(val time.Time) {
	ex.OccurredAt = val
}

func (ex *EventMetaExporter) SetReason(val string) {
	ex.Reason = val
}

type EventMetaReconstitutor struct {
	EventId            uuid.Uuid
	CausationId        uuid.Uuid
	CorrelationId      uuid.Uuid
	CausalDependencies []CausalDependencyReconstitutor
	OccurredAt         time.Time
	Reason             string
}

func (r EventMetaReconstitutor) Reconstitute() (EventMeta, error) {
	causalDependencies := []CausalDependency{}
	for i := range r.CausalDependencies {
		causalDependency, err := r.CausalDependencies[i].Reconstitute()
		if err != nil {
			return EventMeta{}, err
		}
		causalDependencies = append(causalDependencies, causalDependency)
	}
	return NewEventMeta(r.EventId, r.CausationId, r.CorrelationId, causalDependencies, r.OccurredAt, r.Reason)
}
