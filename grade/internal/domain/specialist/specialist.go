package specialist

import (
	"errors"
	"time"

	"github.com/hashicorp/go-multierror"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/aggregate"
	"github.com/emacsway/grade/grade/internal/domain/specialist/assignment"
	"github.com/emacsway/grade/grade/internal/domain/specialist/endorsement"
	"github.com/emacsway/grade/grade/internal/domain/specialist/events"
)

var (
	ErrCrossTenantEndorsement = errors.New(
		"endorser can't endorse cross-tenant members",
	)
	ErrCrossTenantArtifact = errors.New(
		"endorser can't endorse for cross-tenant artifact",
	)
	ErrNotAuthor = errors.New(
		"only author of the artifact can be endorsed",
	)
	ErrEndorsementOneself = errors.New(
		"endorser can't endorse himself",
	)
	ErrLowerGradeEndorses = errors.New(
		"it is allowed to receive endorsements only from members with equal or higher grade",
	)
	ErrAlreadyEndorsed = errors.New(
		"this artifact has already been endorsed by the endorser",
	)
)

// FIXME: Move this constructor to tenant aggregate
func NewSpecialist(
	id member.TenantMemberId,
	createdAt time.Time,
) (*Specialist, error) {
	zeroGrade, _ := grade.DefaultConstructor(0)
	return &Specialist{
		id:        id,
		grade:     zeroGrade,
		createdAt: createdAt,
	}, nil
}

type Specialist struct {
	id                   member.TenantMemberId
	grade                grade.Grade
	receivedEndorsements []endorsement.Endorsement
	assignments          []assignment.Assignment
	createdAt            time.Time
	eventive             aggregate.EventiveEntity
	aggregate.VersionedAggregate
}

func (s *Specialist) ReceiveEndorsement(e endorser.Endorser, a artifact.Artifact, t time.Time) error {
	err := s.canReceiveEndorsement(e, a)
	if err != nil {
		return err
	}
	ent, err := endorsement.NewEndorsement(s.id, s.grade, s.Version(), a.Id(), e.Id(), e.Grade(), e.Version(), t)
	if err != nil {
		return err
	}
	s.receivedEndorsements = append(s.receivedEndorsements, ent)
	s.eventive.AddDomainEvent(events.NewEndorsementReceived(
		s.id, s.grade, s.Version(), e.Id(), e.Grade(), s.Version(), a.Id(), t,
	))
	err = s.actualizeGrade(t)
	if err != nil {
		return err
	}
	return nil
}

func (s Specialist) canReceiveEndorsement(e endorser.Endorser, a artifact.Artifact) error {
	err := e.CanCompleteEndorsement()
	if err != nil {
		return err
	}
	return s.canBeEndorsed(e, a)
}

func (s Specialist) canBeEndorsed(e endorser.Endorser, a artifact.Artifact) error {
	var errs error
	if !e.Id().TenantId().Equal(s.id.TenantId()) {
		errs = multierror.Append(errs, ErrCrossTenantEndorsement)
	}
	if !a.Id().TenantId().Equal(s.id.TenantId()) {
		errs = multierror.Append(errs, ErrCrossTenantArtifact)
	}
	if !a.HasAuthor(s.id) {
		errs = multierror.Append(errs, ErrNotAuthor)
	}
	if e.Id().Equal(s.id) {
		errs = multierror.Append(errs, ErrEndorsementOneself)
	}
	if e.Grade().LessThan(s.grade) {
		errs = multierror.Append(errs, ErrLowerGradeEndorses)
	}
	for i := range s.receivedEndorsements {
		if s.receivedEndorsements[i].IsEndorsedBy(e.Id(), a.Id()) {
			errs = multierror.Append(errs, ErrAlreadyEndorsed)
			break
		}
	}
	return errs
}

func (s Specialist) CanBeginEndorsement(e endorser.Endorser, a artifact.Artifact) error {
	err := e.CanReserveEndorsement()
	if err != nil {
		return err
	}
	return s.canBeEndorsed(e, a)
}

func (s *Specialist) actualizeGrade(t time.Time) error {
	if s.grade.NextGradeAchieved(s.getReceivedEndorsementCount()) {
		assignedGrade, err := s.grade.Next()
		if err != nil {
			return err
		}
		reason, err := assignment.NewReason("Achieved")
		if err != nil {
			return err
		}
		s.eventive.AddDomainEvent(events.NewGradeAssigned(s.id, s.Version(), assignedGrade, reason, t))
		return s.setGrade(assignedGrade, reason, t)
	}
	return nil
}
func (s Specialist) getReceivedEndorsementCount() uint {
	var counter uint
	for i := range s.receivedEndorsements {
		if s.receivedEndorsements[i].SpecialistGrade().Equal(s.grade) {
			counter += uint(s.receivedEndorsements[i].Weight())
		}
	}
	return counter
}

func (s *Specialist) setGrade(g grade.Grade, reason assignment.Reason, t time.Time) error {
	a, err := assignment.NewAssignment(
		s.id, s.Version(), g, reason, t,
	)
	if err != nil {
		return err
	}
	s.assignments = append(s.assignments, a)
	s.grade = g
	return nil
}

func (s *Specialist) DecreaseGrade(reason assignment.Reason, t time.Time) error {
	previousGrade, err := s.grade.Next()
	if err != nil {
		return err
	}
	return s.setGrade(previousGrade, reason, t)
}

func (s Specialist) PendingDomainEvents() []aggregate.DomainEvent {
	return s.eventive.PendingDomainEvents()
}

func (s *Specialist) ClearPendingDomainEvents() {
	s.eventive.ClearPendingDomainEvents()
}

func (s Specialist) Export(ex SpecialistExporterSetter) {
	ex.SetId(s.id)
	ex.SetGrade(s.grade)
	ex.SetVersion(s.Version())
	ex.SetCreatedAt(s.createdAt)

	for i := range s.receivedEndorsements {
		ex.AddEndorsement(s.receivedEndorsements[i])
	}
	for i := range s.assignments {
		ex.AddAssignment(s.assignments[i])
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
