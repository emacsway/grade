package endorsed

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/assignment"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

type EndorsedExporter struct {
	Id                   member.TenantMemberIdExporter
	Grade                seedwork.Uint8Exporter
	ReceivedEndorsements []endorsement.EndorsementExporter
	Assignments          []assignment.AssignmentExporter
	Version              uint
	CreatedAt            time.Time
}

func (ex *EndorsedExporter) SetId(val member.TenantMemberId) {
	val.Export(&ex.Id)
}

func (ex *EndorsedExporter) SetGrade(val grade.Grade) {
	val.Export(&ex.Grade)
}

func (ex *EndorsedExporter) AddEndorsement(val endorsement.Endorsement) {
	var endorsementExporter endorsement.EndorsementExporter
	val.Export(&endorsementExporter)
	ex.ReceivedEndorsements = append(ex.ReceivedEndorsements, endorsementExporter)
}

func (ex *EndorsedExporter) AddAssignment(val assignment.Assignment) {
	var assignmentExporter assignment.AssignmentExporter
	val.Export(&assignmentExporter)
	ex.Assignments = append(ex.Assignments, assignmentExporter)
}

func (ex *EndorsedExporter) SetVersion(val uint) {
	ex.Version = val
}

func (ex *EndorsedExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
