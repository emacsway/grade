package specialist

import (
	"errors"
	"time"

	"github.com/hashicorp/go-multierror"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/specialist/assignment"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/specialist/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/specialist/events"
)

var (
	ErrCrossTenantEndorsement = errors.New(
		"recognizer can't endorse cross-tenant members",
	)
	ErrCrossTenantArtifact = errors.New(
		"recognizer can't endorse for cross-tenant artifact",
	)
	ErrNotAuthor = errors.New(
		"only author of the artifact can be endorsed",
	)
	ErrEndorsementOneself = errors.New(
		"recognizer can't endorse himself",
	)
	ErrLowerGradeEndorses = errors.New(
		"it is allowed to receive endorsements only from members with equal or higher grade",
	)
	ErrAlreadyEndorsed = errors.New(
		"this artifact has already been endorsed by the recognizer",
	)
)

// FIXME: Move this constructor to tenant aggregate
func NewSpecialist(
	id member.TenantMemberId,
	createdAt time.Time,
) (*Specialist, error) {
	versioned := seedwork.NewVersionedAggregate(0)
	eventive := seedwork.NewEventiveEntity()
	zeroGrade, _ := grade.DefaultConstructor(0)
	return &Specialist{
		id:                 id,
		grade:              zeroGrade,
		createdAt:          createdAt,
		VersionedAggregate: versioned,
		EventiveEntity:     eventive,
	}, nil
}

type Specialist struct {
	id                   member.TenantMemberId
	grade                grade.Grade
	receivedEndorsements []endorsement.Endorsement
	assignments          []assignment.Assignment
	createdAt            time.Time
	seedwork.VersionedAggregate
	seedwork.EventiveEntity
}

func (e *Specialist) ReceiveEndorsement(r recognizer.Recognizer, a artifact.Artifact, t time.Time) error {
	err := e.canReceiveEndorsement(r, a)
	if err != nil {
		return err
	}
	ent, err := endorsement.NewEndorsement(
		r.Id(), r.Grade(), r.Version(),
		e.id, e.grade, e.Version(),
		a.Id(), t,
	)
	if err != nil {
		return err
	}
	e.receivedEndorsements = append(e.receivedEndorsements, ent)
	e.AddDomainEvent(events.NewEndorsementReceived(
		e.id, e.grade, e.Version(), r.Id(), r.Grade(), e.Version(), a.Id(), t,
	))
	err = e.actualizeGrade(t)
	if err != nil {
		return err
	}
	return nil
}

func (e Specialist) canReceiveEndorsement(r recognizer.Recognizer, a artifact.Artifact) error {
	err := r.CanCompleteEndorsement()
	if err != nil {
		return err
	}
	return e.canBeEndorsed(r, a)
}

func (e Specialist) canBeEndorsed(r recognizer.Recognizer, a artifact.Artifact) error {
	var errs error
	if !r.Id().TenantId().Equal(e.id.TenantId()) {
		errs = multierror.Append(errs, ErrCrossTenantEndorsement)
	}
	if !a.Id().TenantId().Equal(e.id.TenantId()) {
		errs = multierror.Append(errs, ErrCrossTenantArtifact)
	}
	if !a.HasAuthor(e.id) {
		errs = multierror.Append(errs, ErrNotAuthor)
	}
	if r.Id().Equal(e.id) {
		errs = multierror.Append(errs, ErrEndorsementOneself)
	}
	if r.Grade().LessThan(e.grade) {
		errs = multierror.Append(errs, ErrLowerGradeEndorses)
	}
	for _, ent := range e.receivedEndorsements {
		if ent.IsEndorsedBy(r.Id(), a.Id()) {
			errs = multierror.Append(errs, ErrAlreadyEndorsed)
			break
		}
	}
	return errs
}

func (e Specialist) CanBeginEndorsement(r recognizer.Recognizer, a artifact.Artifact) error {
	err := r.CanReserveEndorsement()
	if err != nil {
		return err
	}
	return e.canBeEndorsed(r, a)
}

func (e *Specialist) actualizeGrade(t time.Time) error {
	if e.grade.NextGradeAchieved(e.getReceivedEndorsementCount()) {
		assignedGrade, err := e.grade.Next()
		if err != nil {
			return err
		}
		reason, err := assignment.NewReason("Achieved")
		if err != nil {
			return err
		}
		e.AddDomainEvent(events.NewGradeAssigned(e.id, e.Version(), assignedGrade, reason, t))
		return e.setGrade(assignedGrade, reason, t)
	}
	return nil
}
func (e Specialist) getReceivedEndorsementCount() uint {
	var counter uint
	for _, v := range e.receivedEndorsements {
		if v.SpecialistGrade().Equal(e.grade) {
			counter += uint(v.Weight())
		}
	}
	return counter
}

func (e *Specialist) setGrade(g grade.Grade, reason assignment.Reason, t time.Time) error {
	a, err := assignment.NewAssignment(
		e.id, e.Version(), g, reason, t,
	)
	if err != nil {
		return err
	}
	e.assignments = append(e.assignments, a)
	e.grade = g
	return nil
}

func (e *Specialist) DecreaseGrade(reason assignment.Reason, t time.Time) error {
	previousGrade, err := e.grade.Next()
	if err != nil {
		return err
	}
	return e.setGrade(previousGrade, reason, t)
}

func (e Specialist) Export(ex SpecialistExporterSetter) {
	ex.SetId(e.id)
	ex.SetGrade(e.grade)
	ex.SetVersion(e.Version())
	ex.SetCreatedAt(e.createdAt)

	for i := range e.receivedEndorsements {
		ex.AddEndorsement(e.receivedEndorsements[i])
	}
	for i := range e.assignments {
		ex.AddAssignment(e.assignments[i])
	}
}

type SpecialistExporterSetter interface {
	SetId(member.TenantMemberId)
	SetGrade(grade.Grade)
	AddEndorsement(endorsement.Endorsement)
	AddAssignment(assignment.Assignment)
	SetVersion(uint)
	SetCreatedAt(time.Time)
}
