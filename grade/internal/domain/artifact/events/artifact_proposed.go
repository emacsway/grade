package events

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/competence"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

type ArtifactProposed struct {
	id               values.TenantArtifactId
	status           values.Status
	name             values.Name
	description      values.Description
	url              values.Url
	competenceIds    []competence.TenantCompetenceId
	authorIds        []member.TenantMemberId
	ownerId          member.TenantMemberId
	createdAt        time.Time
	aggregateVersion uint
}

// EventType should be used instead of Invoke(Aggregate) approach
func (e ArtifactProposed) EventType() string {
	return "ArtifactProposed"
}

func (e ArtifactProposed) EventVersion() uint8 {
	return 1
}

func (e ArtifactProposed) AggregateVersion() uint {
	return e.aggregateVersion
}

func (e *ArtifactProposed) SetAggregateVersion(val uint) {
	e.aggregateVersion = val
}

func (e ArtifactProposed) Export(ex ArtifactProposedExporterSetter) {
	ex.SetId(e.id)
	ex.SetStatus(e.status)
	ex.SetName(e.name)
	ex.SetDescription(e.description)
	ex.SetUrl(e.url)
	for i := range e.competenceIds {
		ex.AddCompetenceId(e.competenceIds[i])
	}
	for i := range e.authorIds {
		ex.AddAuthorId(e.authorIds[i])
	}
	ex.SetDescription(e.description)
	ex.SetOwnerId(e.ownerId)
	ex.SetCreatedAt(e.createdAt)
	ex.SetEventType(e.EventType())
	ex.SetAggregateVersion(e.AggregateVersion())
}

type ArtifactProposedExporterSetter interface {
	SetId(id values.TenantArtifactId)
	SetStatus(values.Status)
	SetName(values.Name)
	SetDescription(values.Description)
	SetUrl(values.Url)
	AddCompetenceId(competence.TenantCompetenceId)
	AddAuthorId(member.TenantMemberId)
	SetOwnerId(member.TenantMemberId)
	SetEventType(string)
	SetAggregateVersion(uint)
	SetCreatedAt(time.Time)
}
