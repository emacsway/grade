package events

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competenceVal "github.com/emacsway/grade/grade/internal/domain/competence/values"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

type ArtifactProposedReconstitutor struct {
	AggregateId      values.ArtifactIdReconstitutor
	Status           uint8
	Name             string
	Description      string
	Url              string
	CompetenceIds    []competenceVal.CompetenceIdReconstitutor
	AuthorIds        []memberVal.MemberIdReconstitutor
	OwnerId          memberVal.MemberIdReconstitutor
	CreatedAt        time.Time
	AggregateVersion uint
	EventMeta        aggregate.EventMetaReconstitutor
}

func (r ArtifactProposedReconstitutor) Reconstitute() (*ArtifactProposed, error) {
	id, err := r.AggregateId.Reconstitute()
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
	competenceIds := []competenceVal.CompetenceId{}
	for i := range r.CompetenceIds {
		competenceId, err := r.CompetenceIds[i].Reconstitute()
		if err != nil {
			return nil, err
		}
		competenceIds = append(competenceIds, competenceId)
	}
	authorIds := []memberVal.MemberId{}
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
	eventMeta, err := r.EventMeta.Reconstitute()
	if err != nil {
		return nil, err
	}
	return &ArtifactProposed{
		aggregateId:      id,
		status:           status,
		name:             name,
		description:      description,
		url:              url,
		competenceIds:    competenceIds,
		authorIds:        authorIds,
		ownerId:          ownerId,
		createdAt:        r.CreatedAt,
		aggregateVersion: r.AggregateVersion,
		eventMeta:        eventMeta,
	}, nil
}
