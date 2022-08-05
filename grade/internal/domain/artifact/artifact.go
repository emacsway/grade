package artifact

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

func NewArtifact(
	id TenantArtifactId,
	status Status,
	name Name,
	description Description,
	url Url,
	competenceIds []competence.TenantCompetenceId,
	authorIds []member.TenantMemberId,
	ownerId member.TenantMemberId,
	createdAt time.Time,
) *Artifact {
	return &Artifact{
		id:            id,
		status:        status,
		name:          name,
		description:   description,
		url:           url,
		competenceIds: competenceIds,
		authorIds:     authorIds,
		ownerId:       ownerId,
		createdAt:     createdAt,
	}
}

// Artifact is a good candidate for EventSourcing
type Artifact struct {
	id            TenantArtifactId
	status        Status
	name          Name
	description   Description
	url           Url
	competenceIds []competence.TenantCompetenceId
	authorIds     []member.TenantMemberId
	ownerId       member.TenantMemberId
	createdAt     time.Time
	eventSourced  aggregate.EventSourcedAggregate
}

func (a Artifact) Id() TenantArtifactId {
	return a.id
}

func (a Artifact) HasAuthor(authorId member.TenantMemberId) bool {
	for i := range a.authorIds {
		if a.authorIds[i].Equal(authorId) {
			return true
		}
	}
	return false
}

func (a Artifact) PendingDomainEvents() []aggregate.DomainEvent {
	return a.eventSourced.PendingDomainEvents()
}

func (a *Artifact) ClearPendingDomainEvents() {
	a.eventSourced.ClearPendingDomainEvents()
}

func (a Artifact) Version() uint {
	return a.eventSourced.Version()
}

func (a *Artifact) SetVersion(val uint) {
	a.eventSourced.SetVersion(val)
}
