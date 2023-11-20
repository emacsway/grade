package events

import (
	"github.com/emacsway/grade/grade/internal/domain/competence/values"
)

func NewNameUpdated(
	id values.CompetenceId,
	name values.Name,
) *NameUpdated {
	return &NameUpdated{
		id:   id,
		name: name,
	}
}

type NameUpdated struct {
	id               values.CompetenceId
	name             values.Name
	aggregateVersion uint
}

func (e NameUpdated) EventType() string {
	return "NameUpdated"
}

func (e NameUpdated) EventVersion() uint8 {
	return 1
}

func (e NameUpdated) AggregateVersion() uint {
	return e.aggregateVersion
}

func (e *NameUpdated) SetAggregateVersion(val uint) {
	e.aggregateVersion = val
}

func (e NameUpdated) Export(ex NameUpdatedExporterSetter) {
	ex.SetId(e.id)
	ex.SetName(e.name)
	ex.SetEventType(e.EventType())
	ex.SetAggregateVersion(e.AggregateVersion())
}

type NameUpdatedExporterSetter interface {
	SetId(id values.CompetenceId)
	SetName(values.Name)
	SetEventType(string)
	SetAggregateVersion(uint)
}
