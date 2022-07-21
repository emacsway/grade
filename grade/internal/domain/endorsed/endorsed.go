package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement/interfaces"
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
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

func (e Endorsed) Export() EndorsedState {
	var receivedEndorsements []endorsement.EndorsementState
	for _, v := range e.receivedEndorsements {
		receivedEndorsements = append(receivedEndorsements, v.Export())
	}
	return EndorsedState{
		e.id.Export(),
		e.userId.Export(),
		e.grade.Export(),
		receivedEndorsements,
		e.GetVersion(),
		e.createdAt,
	}
}

func (e Endorsed) ExportTo(ex EndorsedExporter) {
	var id, userId seedwork.Uint64Exporter
	var grade seedwork.Uint8Exporter
	var receivedEndorsements []interfaces.EndorsementExporter

	for _, v := range e.receivedEndorsements {
		re := &endorsement.EndorsementExporter{}
		v.ExportTo(re)
		receivedEndorsements = append(receivedEndorsements, re)
	}

	e.id.ExportTo(&id)
	e.userId.ExportTo(&userId)
	e.grade.ExportTo(&grade)
	ex.SetState(
		&id, &userId, &grade, receivedEndorsements, e.GetVersion(), e.createdAt,
	)
}

type EndorsedState struct {
	Id                   uint64
	UserId               uint64
	Grade                uint8
	ReceivedEndorsements []endorsement.EndorsementState
	Version              uint
	CreatedAt            time.Time
}
