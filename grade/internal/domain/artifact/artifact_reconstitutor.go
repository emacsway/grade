package artifact

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competence "github.com/emacsway/grade/grade/internal/domain/competence/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

// TODO: Remove me and use only anArtifact.eventSourced.LoadFrom(pastEvents []PersistentDomainEvent)

type ArtifactReconstitutor struct {
	Id            values.TenantArtifactIdReconstitutor
	Status        uint8
	Name          string
	Description   string
	Url           string
	CompetenceIds []competence.TenantCompetenceIdReconstitutor
	AuthorIds     []member.TenantMemberIdReconstitutor
	OwnerId       member.TenantMemberIdReconstitutor
	CreatedAt     time.Time
	Version       uint
}

func (r ArtifactReconstitutor) Reconstitute() (*Artifact, error) {
	id, err := r.Id.Reconstitute()
	if err != nil {
		return nil, err
	}
	status, err := values.NewStatus(r.Status)
	if err != nil {
		return nil, err
	}
	name, err := values.NewName(r.Name)
	if err != nil {
		return nil, err
	}
	description, err := values.NewDescription(r.Description)
	if err != nil {
		return nil, err
	}
	url, err := values.NewUrl(r.Url)
	if err != nil {
		return nil, err
	}
	competenceIds := []competence.TenantCompetenceId{}
	for i := range r.CompetenceIds {
		competenceId, err := r.CompetenceIds[i].Reconstitute()
		if err != nil {
			return nil, err
		}
		competenceIds = append(competenceIds, competenceId)
	}
	authorIds := []member.TenantMemberId{}
	for i := range r.AuthorIds {
		authorId, err := r.AuthorIds[i].Reconstitute()
		if err != nil {
			return nil, err
		}
		authorIds = append(authorIds, authorId)
	}
	ownerId, err := r.OwnerId.Reconstitute()
	if err != nil {
		return nil, err
	}
	return &Artifact{
		id:            id,
		status:        status,
		name:          name,
		description:   description,
		url:           url,
		competenceIds: competenceIds,
		authorIds:     authorIds,
		ownerId:       ownerId,
		createdAt:     r.CreatedAt,
		eventSourced:  aggregate.NewEventSourcedAggregate(r.Version),
	}, nil
}
