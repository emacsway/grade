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
	ErrNoEndorsementAvailable         = errors.New("no endorsement is available")
	ErrNoEndorsementReservation       = errors.New("there is no endorsement reservation")
	ErrEndorsementReservationExceeded = errors.New("endorsement reservation exceeded")
)

func NewRecognizer(
	id recognizer.RecognizerId,
	memberId external.MemberId,
	grade shared.Grade,
	createdAt time.Time,
) (*Recognizer, error) {
	availableCount, err := recognizer.NewEndorsementCount(recognizer.YearlyEndorsementCount)
	if err != nil {
		return nil, err
	}
	pendingCount, err := recognizer.NewEndorsementCount(0)
	if err != nil {
		return nil, err
	}
	versioned, err := seedwork.NewVersionedAggregate(0)
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
		availableEndorsementCount: availableCount,
		pendingEndorsementCount:   pendingCount,
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
		return ErrEndorsementReservationExceeded
	}
	return nil
}

func (r Recognizer) CanCompleteEndorsement() error {
	if r.pendingEndorsementCount == 0 {
		return ErrNoEndorsementReservation
	}
	if r.availableEndorsementCount == 0 {
		return ErrNoEndorsementAvailable
	}
	if r.availableEndorsementCount < r.pendingEndorsementCount {
		return ErrEndorsementReservationExceeded
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

func (r *Recognizer) ReleaseEndorsementReservation() error {
	if r.pendingEndorsementCount == 0 {
		return ErrNoEndorsementReservation
	}
	r.pendingEndorsementCount -= 1
	return nil
}

func (r *Recognizer) CompleteEndorsement() error {
	err := r.CanCompleteEndorsement()
	if err != nil {
		return err
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
