package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

func NewEndorsed(
	id endorsed.EndorsedId,
	userId external.UserId,
	grade shared.Grade,
	endorsements []endorsement.Endorsement,
	version uint,
	createdAt time.Time,
) (*Endorsed, error) {
	versioned, err := seedwork.NewVersionedAggregate(version)
	if err != nil {
		return nil, err
	}
	eventive, err := seedwork.NewEventiveEntity()
	if err != nil {
		return nil, err
	}
	return &Endorsed{
		id:                   id,
		userId:               userId,
		grade:                grade,
		receivedEndorsements: endorsements,
		VersionedAggregate:   versioned,
		EventiveEntity:       eventive,
		createdAt:            createdAt,
	}, nil
}

type Endorsed struct {
	id                   endorsed.EndorsedId
	userId               external.UserId
	grade                shared.Grade
	receivedEndorsements []endorsement.Endorsement
	createdAt            time.Time
	seedwork.VersionedAggregate
	seedwork.EventiveEntity
}

func (e Endorsed) CreateMemento() EndorsedMemento {
	var receivedEndorsements []endorsement.EndorsementMemento
	for i, v := range e.receivedEndorsements {
		receivedEndorsements[i] = v.CreateMemento()
	}
	return EndorsedMemento{
		e.id.CreateMemento(),
		e.userId.CreateMemento(),
		e.grade.CreateMemento(),
		receivedEndorsements,
		e.GetVersion(),
		e.createdAt,
	}
}

type EndorsedMemento struct {
	Id                   uint64
	UserId               uint64
	Grade                uint8
	ReceivedEndorsements []endorsement.EndorsementMemento
	Version              uint
	CreatedAt            time.Time
}
