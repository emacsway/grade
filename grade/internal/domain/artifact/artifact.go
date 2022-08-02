package artifact

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/expertisearea"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewArtifact(
	id TenantArtifactId,
	status Status,
	name Name,
	description Description,
	url Url,
	expertiseAreaIds []expertisearea.ExpertiseAreaId,
	authorIds []member.TenantMemberId,
	createdById member.TenantMemberId,
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
		expertiseAreaIds:   expertiseAreaIds,
		authorIds:          authorIds,
		createdById:        createdById,
		createdAt:          createdAt,
		VersionedAggregate: versioned,
		EventiveEntity:     eventive,
	}
}

// Artifact is a good candidate for EventSourcing
type Artifact struct {
	id               TenantArtifactId
	status           Status
	name             Name
	description      Description
	url              Url
	expertiseAreaIds []expertisearea.ExpertiseAreaId
	authorIds        []member.TenantMemberId
	createdById      member.TenantMemberId
	createdAt        time.Time
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
