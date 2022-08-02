package specialist

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/specialist/assignment"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/specialist/endorsement"
)

type SpecialistExporter struct {
	Id                   member.TenantMemberIdExporter
	Grade                seedwork.Uint8Exporter
	ReceivedEndorsements []endorsement.EndorsementExporter
	Assignments          []assignment.AssignmentExporter
	Version              uint
	CreatedAt            time.Time
}

func (ex *SpecialistExporter) SetId(val member.TenantMemberId) {
	val.Export(&ex.Id)
}

func (ex *SpecialistExporter) SetGrade(val grade.Grade) {
	val.Export(&ex.Grade)
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

func (ex *SpecialistExporter) SetVersion(val uint) {
	ex.Version = val
}

func (ex *SpecialistExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
