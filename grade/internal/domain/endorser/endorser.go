package endorser

import (
	"errors"
	"time"

	"github.com/emacsway/grade/grade/internal/domain/endorser/values"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
)

var (
	ErrNoEndorsementAvailable         = errors.New("no endorsement is available")
	ErrNoEndorsementReservation       = errors.New("no endorsement is reserved")
	ErrEndorsementReservationExceeded = errors.New("endorsement reservation exceeded")
)

// FIXME: Move this constructor to tenant aggregate

func NewEndorser(
	id member.TenantMemberId,
	createdAt time.Time,
) (*Endorser, error) {
	availableCount, err := values.NewEndorsementCount(values.YearlyEndorsementCount)
	if err != nil {
		return nil, err
	}
	pendingCount, err := values.NewEndorsementCount(0)
	if err != nil {
		return nil, err
	}
	zeroGrade, _ := grade.NewGradeFactory(grade.MaxGradeValue, grade.GradeMatrix)(0)
	return &Endorser{
		id:                        id,
		grade:                     zeroGrade,
		availableEndorsementCount: availableCount,
		pendingEndorsementCount:   pendingCount,
		createdAt:                 createdAt,
	}, nil
}

// TODO: Use
// - https://martinfowler.com/eaaDev/TemporalProperty.html
// - https://martinfowler.com/eaaDev/TemporalObject.html
// to track grade by version?

type Endorser struct { // TODO: rename to Recognitory | Endorser | Originator | Sender (to Receiver)
	id                        member.TenantMemberId
	grade                     grade.Grade
	availableEndorsementCount values.EndorsementCount
	pendingEndorsementCount   values.EndorsementCount
	createdAt                 time.Time
	eventive                  aggregate.EventiveEntity
	aggregate.VersionedAggregate
}

func (e Endorser) Id() member.TenantMemberId {
	return e.id
}

func (e Endorser) Grade() grade.Grade {
	return e.grade
}

func (e *Endorser) SetGrade(val grade.Grade) error {
	e.grade = val
	return nil
}

func (e Endorser) CanReserveEndorsement() error {
	if !(e.availableEndorsementCount > e.pendingEndorsementCount) {
		return ErrEndorsementReservationExceeded
	}
	return nil
}

// TODO: Use Specification pattern instead?
// https://enterprisecraftsmanship.com/posts/specification-pattern-always-valid-domain-model/

func (e Endorser) CanCompleteEndorsement() error {
	if e.pendingEndorsementCount == 0 {
		return ErrNoEndorsementReservation
	}
	if e.availableEndorsementCount == 0 {
		return ErrNoEndorsementAvailable
	}
	if e.availableEndorsementCount < e.pendingEndorsementCount {
		return ErrEndorsementReservationExceeded
	}
	return nil
}

func (e *Endorser) ReserveEndorsement() error {
	err := e.CanReserveEndorsement()
	if err != nil {
		return err
	}
	e.pendingEndorsementCount += 1
	return nil
}

func (e *Endorser) ReleaseEndorsementReservation() error {
	if e.pendingEndorsementCount == 0 {
		return ErrNoEndorsementReservation
	}
	e.pendingEndorsementCount -= 1
	return nil
}

func (e *Endorser) CompleteEndorsement() error {
	err := e.CanCompleteEndorsement()
	if err != nil {
		return err
	}
	e.availableEndorsementCount -= 1
	e.pendingEndorsementCount -= 1
	return nil
}

func (e Endorser) PendingDomainEvents() []aggregate.DomainEvent {
	return e.eventive.PendingDomainEvents()
}

func (e *Endorser) ClearPendingDomainEvents() {
	e.eventive.ClearPendingDomainEvents()
}

func (e Endorser) Export(ex EndorserExporterSetter) {
	ex.SetId(e.id)
	ex.SetGrade(e.grade)
	ex.SetAvailableEndorsementCount(e.availableEndorsementCount)
	ex.SetPendingEndorsementCount(e.pendingEndorsementCount)
	ex.SetCreatedAt(e.createdAt)
	ex.SetVersion(e.Version())
}

type EndorserExporterSetter interface {
	SetId(member.TenantMemberId)
	SetGrade(grade.Grade)
	SetAvailableEndorsementCount(values.EndorsementCount)
	SetPendingEndorsementCount(values.EndorsementCount)
	SetCreatedAt(time.Time)
	SetVersion(uint)
}
