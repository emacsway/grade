package artifact

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact/values"
	competenceVal "github.com/emacsway/grade/grade/internal/domain/competence/values"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

type ArtifactReconstitutor struct {
	Snapshot   *ArtifactSnapshotReconstitutor
	PastEvents []aggregate.PersistentDomainEvent
}

func (r ArtifactReconstitutor) Reconstitute() (agg *Artifact, err error) {
	if r.Snapshot != nil {
		agg, err = r.Snapshot.Reconstitute()
		if err != nil {
			return nil, err
		}
	} else {
		agg, err = EmptyAggregate()
		if err != nil {
			return nil, err
		}
	}
	agg.eventSourced.LoadFrom(r.PastEvents)
	return agg, nil
}

// It will be used to load snapshot only.
// To load events use anArtifact.eventSourced.LoadFrom(pastEvents []PersistentDomainEvent)

type ArtifactSnapshotReconstitutor struct {
	Id            values.ArtifactIdReconstitutor
	Status        uint8
	Name          string
	Description   string
	Url           string
	CompetenceIds []competenceVal.CompetenceIdReconstitutor
	AuthorIds     []memberVal.MemberIdReconstitutor
	OwnerId       memberVal.MemberIdReconstitutor
	CreatedAt     time.Time
	Version       uint
}

func (r ArtifactSnapshotReconstitutor) Reconstitute() (*Artifact, error) {
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
	agg, err := NewArtifact(
		id,
		status,
		name,
		description,
		url,
		competenceIds,
		authorIds,
		ownerId,
		r.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	agg.SetVersion(r.Version)
	return agg, nil
}
