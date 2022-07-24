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
	memberId external.MemberId,
	grade shared.Grade,
	availableEndorsementCount recognizer.EndorsementCount,
	pendingEndorsementCount recognizer.EndorsementCount,
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
		memberId:                  memberId,
		grade:                     grade,
		availableEndorsementCount: availableEndorsementCount,
		pendingEndorsementCount:   pendingEndorsementCount,
		createdAt:                 createdAt,
		VersionedAggregate:        versioned,
		EventiveEntity:            eventive,
	}, nil
}

type Recognizer struct {
	id                        recognizer.RecognizerId
	memberId                  external.MemberId
	grade                     shared.Grade
	availableEndorsementCount recognizer.EndorsementCount
	pendingEndorsementCount   recognizer.EndorsementCount
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
		r.id.Export(), r.memberId.Export(), r.grade.Export(),
		r.availableEndorsementCount.Export(),
		r.pendingEndorsementCount.Export(), r.GetVersion(), r.createdAt,
	}
}

func (r Recognizer) ExportTo(ex interfaces.RecognizerExporter) {
	var id, memberId seedwork.Uint64Exporter
	var grade, availableEndorsementCount, pendingEndorsementCount seedwork.Uint8Exporter

	r.id.ExportTo(&id)
	r.memberId.ExportTo(&memberId)
	r.grade.ExportTo(&grade)
	r.availableEndorsementCount.ExportTo(&availableEndorsementCount)
	r.pendingEndorsementCount.ExportTo(&pendingEndorsementCount)
	ex.SetState(
		&id, &memberId, &grade, &availableEndorsementCount, &pendingEndorsementCount, r.GetVersion(), r.createdAt,
	)
}

type RecognizerState struct {
	Id                        uint64
	MemberId                  uint64
	Grade                     uint8
	AvailableEndorsementCount uint8
	PendingEndorsementCount   uint8
	Version                   uint
	CreatedAt                 time.Time
}
