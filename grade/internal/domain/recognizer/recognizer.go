package recognizer

import (
	"errors"
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

var (
	ErrNoEndorsementAvailable     = errors.New("no endorsement is available")
	ErrNoEndorsementCanBeReserved = errors.New("no endorsement can be reserved")
	ErrNoEndorsementReservation   = errors.New("there is no endorsement reservation")
	ErrRecognizerUnableToComplete = errors.New("recognizer is not able to complete endorsement")
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

func (r Recognizer) canReserveEndorsement() error {
	if !(r.availableEndorsementCount > r.pendingEndorsementCount) {
		return ErrNoEndorsementCanBeReserved
	}
	return nil
}

func (r Recognizer) CanCompleteEndorsement() error {
	if !(r.pendingEndorsementCount > 0 && r.availableEndorsementCount >= r.pendingEndorsementCount) {
		return ErrRecognizerUnableToComplete
	}
	return nil
}

func (r *Recognizer) ReserveEndorsement() error {
	err := r.canReserveEndorsement()
	if err != nil {
		return err
	}
	r.pendingEndorsementCount += 1
	return nil
}

func (r *Recognizer) ReleaseEndorsementReservation() {
	r.pendingEndorsementCount -= 1
}

func (r *Recognizer) CompleteEndorsement() error {
	if r.availableEndorsementCount == 0 {
		return ErrNoEndorsementAvailable
	}
	if r.pendingEndorsementCount == 0 {
		return ErrNoEndorsementReservation
	}
	r.availableEndorsementCount -= 1
	r.pendingEndorsementCount -= 1
	return nil
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

func (r Recognizer) Export() RecognizerState {
	return RecognizerState{
		r.id.Export(), r.memberId.Export(), r.grade.Export(),
		r.availableEndorsementCount.Export(),
		r.pendingEndorsementCount.Export(), r.GetVersion(), r.createdAt,
	}
}
