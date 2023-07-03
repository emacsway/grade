package events

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence/values"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
)

type CompetenceCreated struct {
	id               values.TenantCompetenceId
	name             values.Name
	ownerId          memberVal.TenantMemberId
	createdAt        time.Time
	aggregateVersion uint
}

func (e CompetenceCreated) EventType() string {
	return "CompetenceAvailable"
}

func (e CompetenceCreated) EventVersion() uint8 {
	return 1
}

func (e CompetenceCreated) AggregateVersion() uint {
	return e.aggregateVersion
}

func (e *CompetenceCreated) SetAggregateVersion(val uint) {
	e.aggregateVersion = val
}

func (e CompetenceCreated) Export(ex ArtifactProposedExporterSetter) {
	ex.SetId(e.id)
	ex.SetName(e.name)
	ex.SetOwnerId(e.ownerId)
	ex.SetCreatedAt(e.createdAt)
	ex.SetEventType(e.EventType())
	ex.SetAggregateVersion(e.AggregateVersion())
}

type ArtifactProposedExporterSetter interface {
	SetId(id values.TenantCompetenceId)
	SetName(values.Name)
	SetOwnerId(memberVal.TenantMemberId)
	SetCreatedAt(time.Time)
	SetEventType(string)
	SetAggregateVersion(uint)
}
