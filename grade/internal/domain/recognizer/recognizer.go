package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

func NewRecognizer(
	id recognizer.RecognizerId,
	userId external.UserId,
	grade shared.Grade,
	availableEndorsementCount recognizer.AvailableEndorsementCount,
	version uint,
	createdAt time.Time,
) (*Recognizer, error) {
	versioned, err := seedwork.NewVersionedAggregate(version)
	if err != nil {
		return nil, err
	}
	eventive, err := seedwork.NewEventiveEntity()
	if err != nil {
		return nil, err
	}
	return &Recognizer{
		id:                        id,
		userId:                    userId,
		grade:                     grade,
		availableEndorsementCount: availableEndorsementCount,
		createdAt:                 createdAt,
		VersionedAggregate:        versioned,
		EventiveEntity:            eventive,
	}, nil
}

type Recognizer struct {
	id                        recognizer.RecognizerId
	userId                    external.UserId
	grade                     shared.Grade
	availableEndorsementCount recognizer.AvailableEndorsementCount
	createdAt                 time.Time
	seedwork.VersionedAggregate
	seedwork.EventiveEntity
}

func (r Recognizer) GetId() recognizer.RecognizerId {
	return r.id
}

func (r Recognizer) GetGrade() shared.Grade {
	return r.grade
}

func (r Recognizer) Export() RecognizerState {
	return RecognizerState{
		r.id.Export(),
		r.userId.Export(),
		r.grade.Export(),
		r.availableEndorsementCount.Export(),
		r.GetVersion(),
		r.createdAt,
	}
}

func (r Recognizer) ExportTo(ex interfaces.RecognizerExporter) {
	var id, userId seedwork.Uint64Exporter
	var grade, availableEndorsementCount seedwork.Uint8Exporter

	id.SetState(r.id.Export())
	userId.SetState(r.userId.Export())
	grade.SetState(r.grade.Export())
	availableEndorsementCount.SetState(r.availableEndorsementCount.Export())
	ex.SetState(
		id,
		userId,
		grade,
		availableEndorsementCount,
		r.GetVersion(),
		r.createdAt,
	)
}

type RecognizerState struct {
	Id                        uint64
	UserId                    uint64
	Grade                     uint8
	AvailableEndorsementCount uint8
	Version                   uint
	CreatedAt                 time.Time
}
