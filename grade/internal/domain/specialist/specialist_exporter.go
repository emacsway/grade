package specialist

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/specialist/assignment"
	"github.com/emacsway/grade/grade/internal/domain/specialist/endorsement"
)

type SpecialistExporter struct {
	Id                   member.MemberIdExporter
	Grade                uint8
	ReceivedEndorsements []endorsement.EndorsementExporter
	Assignments          []assignment.AssignmentExporter
	CreatedAt            time.Time
	Version              uint
}

func (ex *SpecialistExporter) SetId(val member.MemberId) {
	val.Export(&ex.Id)
}

func (ex *SpecialistExporter) SetGrade(val grade.Grade) {
	val.Export(func(v uint8) { ex.Grade = v })
}

func (ex *SpecialistExporter) AddEndorsement(val endorsement.Endorsement) {
	var endorsementExporter endorsement.EndorsementExporter
	val.Export(&endorsementExporter)
	ex.ReceivedEndorsements = append(ex.ReceivedEndorsements, endorsementExporter)
}

func (ex *SpecialistExporter) AddAssignment(val assignment.Assignment) {
	var assignmentExporter assignment.AssignmentExporter
	val.Export(&assignmentExporter)
	ex.Assignments = append(ex.Assignments, assignmentExporter)
}

func (ex *SpecialistExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}

func (ex *SpecialistExporter) SetVersion(val uint) {
	ex.Version = val
}
