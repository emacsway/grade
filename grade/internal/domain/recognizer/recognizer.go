package recognizer

import (
	"errors"
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

var (
	ErrNoEndorsementAvailable         = errors.New("no endorsement is available")
	ErrNoEndorsementReservation       = errors.New("no endorsement is reserved")
	ErrEndorsementReservationExceeded = errors.New("endorsement reservation exceeded")
)

// FIXME: Move this constructor to tenant aggregate
func NewRecognizer(
	id member.TenantMemberId,
	createdAt time.Time,
) (*Recognizer, error) {
	availableCount, err := NewEndorsementCount(YearlyEndorsementCount)
	if err != nil {
		return nil, err
	}
	pendingCount, err := NewEndorsementCount(0)
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
	zeroGrade, _ := grade.NewGradeFactory(grade.MaxGradeValue, grade.GradeMatrix)(0)
	return &Recognizer{
		id:                        id,
		grade:                     zeroGrade,
		availableEndorsementCount: availableCount,
		pendingEndorsementCount:   pendingCount,
		createdAt:                 createdAt,
		VersionedAggregate:        versioned,
		EventiveEntity:            eventive,
	}, nil
}

type Recognizer struct {
	id                        member.TenantMemberId
	grade                     grade.Grade
	availableEndorsementCount EndorsementCount
	pendingEndorsementCount   EndorsementCount
	createdAt                 time.Time
	seedwork.VersionedAggregate
	seedwork.EventiveEntity
}

func (r Recognizer) GetId() member.TenantMemberId {
	return r.id
}

func (r Recognizer) GetGrade() grade.Grade {
	return r.grade
}

func (r *Recognizer) SetGrade(val grade.Grade) error {
	r.grade = val
	return nil
}

func (r Recognizer) CanReserveEndorsement() error {
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
	err := r.CanReserveEndorsement()
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

func (r Recognizer) Export(ex RecognizerExporterSetter) {
	ex.SetId(r.id)
	ex.SetGrade(r.grade)
	ex.SetAvailableEndorsementCount(r.availableEndorsementCount)
	ex.SetPendingEndorsementCount(r.pendingEndorsementCount)
	ex.SetVersion(r.GetVersion())
	ex.SetCreatedAt(r.createdAt)
}

type RecognizerExporterSetter interface {
	SetId(member.TenantMemberId)
	SetGrade(grade.Grade)
	SetAvailableEndorsementCount(EndorsementCount)
	SetPendingEndorsementCount(EndorsementCount)
	SetVersion(uint)
	SetCreatedAt(time.Time)
}
