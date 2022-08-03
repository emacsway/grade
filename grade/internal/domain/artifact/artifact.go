package artifact

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/competence"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork"
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
	versioned := seedwork.NewVersionedAggregate(0)
	eventive := seedwork.NewEventiveEntity()

	return &Artifact{
		id:                 id,
		status:             status,
		name:               name,
		description:        description,
		url:                url,
		competenceIds:      competenceIds,
		authorIds:          authorIds,
		ownerId:            ownerId,
		createdAt:          createdAt,
		VersionedAggregate: versioned,
		EventiveEntity:     eventive,
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
	seedwork.VersionedAggregate
	seedwork.EventiveEntity
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
