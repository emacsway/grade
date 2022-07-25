package endorsed

import (
	"errors"
	interfaces2 "github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry/interfaces"
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

var (
	ErrAlreadyEndorsed = errors.New(
		"this artifact has already been endorsed by the recogniser",
	)
)

func NewEndorsed(
	id endorsed.EndorsedId,
	memberId external.MemberId,
	grade shared.Grade,
	endorsements []endorsement.Endorsement,
	gradeLogEntries []gradelogentry.GradeLogEntry,
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
		memberId:             memberId,
		grade:                grade,
		receivedEndorsements: endorsements,
		gradeLogEntries:      gradeLogEntries,
		VersionedAggregate:   versioned,
		EventiveEntity:       eventive,
		createdAt:            createdAt,
	}, nil
}

type Endorsed struct {
	id                   endorsed.EndorsedId
	memberId             external.MemberId
	grade                shared.Grade
	receivedEndorsements []endorsement.Endorsement
	gradeLogEntries      []gradelogentry.GradeLogEntry
	createdAt            time.Time
	seedwork.VersionedAggregate
	seedwork.EventiveEntity
}

func (e *Endorsed) ReceiveEndorsement(r recognizer.Recognizer, aId artifact.ArtifactId, t time.Time) error {
	err := e.canReceiveEndorsement(r, aId)
	if err != nil {
		return err
	}
	ent, err := endorsement.NewEndorsement(
		r.GetId(), r.GetGrade(), r.GetVersion(),
		e.id, e.grade, e.GetVersion(),
		aId, t,
	)
	if err != nil {
		return err
	}
	e.receivedEndorsements = append(e.receivedEndorsements, ent)
	err = e.actualizeGrade(t)
	if err != nil {
		return err
	}
	return nil
}

func (e Endorsed) canReceiveEndorsement(r recognizer.Recognizer, aId artifact.ArtifactId) error {
	err := r.CanCompleteEndorsement()
	if err != nil {
		return err
	}
	for _, ent := range e.receivedEndorsements {
		if ent.IsEndorsedBy(r.GetId(), aId) {
			return ErrAlreadyEndorsed
		}
	}
	return nil
}

func (e *Endorsed) actualizeGrade(t time.Time) error {
	var err error
	if e.grade == shared.WithoutGrade && e.getReceivedEndorsementCount() >= 6 {
		err = e.setGrade(shared.Grade3, t)
	} else if e.grade == shared.Grade3 && e.getReceivedEndorsementCount() >= 10 {
		err = e.setGrade(shared.Grade2, t)
	} else if e.grade == shared.Grade2 && e.getReceivedEndorsementCount() >= 14 {
		err = e.setGrade(shared.Grade1, t)
	} else if e.grade == shared.Grade1 && e.getReceivedEndorsementCount() >= 20 {
		err = e.setGrade(shared.Candidate, t)
	} else if e.grade == shared.Candidate && e.getReceivedEndorsementCount() >= 10 {
		err = e.setGrade(shared.Expert, t)
	}
	return err
}
func (e Endorsed) getReceivedEndorsementCount() uint {
	var counter uint
	for _, v := range e.receivedEndorsements {
		if v.GetEndorsedGrade() == e.grade {
			counter += uint(v.GetWeight())
		}
	}
	return counter
}

func (e *Endorsed) setGrade(g shared.Grade, t time.Time) error {
	gle, err := gradelogentry.NewGradeLogEntry(
		e.id, e.GetVersion(), g, t,
	)
	if err != nil {
		return err
	}
	e.gradeLogEntries = append(e.gradeLogEntries, gle)
	e.grade = g
	return nil
}

func (e Endorsed) ExportTo(ex EndorsedExporter) {
	var id, memberId seedwork.Uint64Exporter
	var grade seedwork.Uint8Exporter
	var receivedEndorsements []interfaces.EndorsementExporter
	var gradeLogEntries []interfaces2.GradeLogEntryExporter

	for _, v := range e.receivedEndorsements {
		re := &endorsement.EndorsementExporter{}
		v.ExportTo(re)
		receivedEndorsements = append(receivedEndorsements, re)
	}

	for _, v := range e.gradeLogEntries {
		gle := &gradelogentry.GradeLogEntryExporter{}
		v.ExportTo(gle)
		gradeLogEntries = append(gradeLogEntries, gle)
	}

	e.id.ExportTo(&id)
	e.memberId.ExportTo(&memberId)
	e.grade.ExportTo(&grade)
	ex.SetState(
		&id, &memberId, &grade, receivedEndorsements, gradeLogEntries, e.GetVersion(), e.createdAt,
	)
}

func (e Endorsed) Export() EndorsedState {
	var receivedEndorsements []endorsement.EndorsementState
	var gradeLogEntries []gradelogentry.GradeLogEntryState
	for _, v := range e.receivedEndorsements {
		receivedEndorsements = append(receivedEndorsements, v.Export())
	}
	for _, v := range e.gradeLogEntries {
		gradeLogEntries = append(gradeLogEntries, v.Export())
	}
	return EndorsedState{
		e.id.Export(), e.memberId.Export(), e.grade.Export(),
		receivedEndorsements, gradeLogEntries,
		e.GetVersion(), e.createdAt,
	}
}