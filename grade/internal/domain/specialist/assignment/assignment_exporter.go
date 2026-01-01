package assignment

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/specialist/assignment/values"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
)

type AssignmentExporter struct {
	SpecialistId      member.MemberIdExporter
	SpecialistVersion uint
	AssignedGrade     exporters.Uint8Exporter
	Reason            exporters.StringExporter
	CreatedAt         time.Time
}

func (ex *AssignmentExporter) SetSpecialistId(val member.MemberId) {
	val.Export(&ex.SpecialistId)
}

func (ex *AssignmentExporter) SetSpecialistVersion(val uint) {
	ex.SpecialistVersion = val
}

func (ex *AssignmentExporter) SetAssignedGrade(val grade.Grade) {
	val.Export(&ex.AssignedGrade)
}

func (ex *AssignmentExporter) SetReason(val values.Reason) {
	val.Export(&ex.Reason)
}

func (ex *AssignmentExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
