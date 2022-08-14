package artifact

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence"
	"github.com/emacsway/grade/grade/internal/domain/member"
)

type ArtifactProposed struct {
	id               TenantArtifactId
	status           Status
	name             Name
	description      Description
	url              Url
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
	SetId(id TenantArtifactId)
	SetStatus(Status)
	SetName(Name)
	SetDescription(Description)
	SetUrl(Url)
	AddCompetenceId(competence.TenantCompetenceId)
	AddAuthorId(member.TenantMemberId)
	SetOwnerId(member.TenantMemberId)
	SetEventType(string)
	SetAggregateVersion(uint)
	SetCreatedAt(time.Time)
}
