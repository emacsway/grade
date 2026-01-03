package assignment

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/specialist/assignment/values"
)

type AssignmentExporter struct {
	SpecialistId      member.MemberIdExporter
	SpecialistVersion uint
	AssignedGrade     uint8
	Reason            string
	CreatedAt         time.Time
}

func (ex *AssignmentExporter) SetSpecialistId(val member.MemberId) {
	val.Export(&ex.SpecialistId)
}

func (ex *AssignmentExporter) SetSpecialistVersion(val uint) {
	ex.SpecialistVersion = val
}

func (ex *AssignmentExporter) SetAssignedGrade(val grade.Grade) {
	val.Export(func(v uint8) { ex.AssignedGrade = v })
}

func (ex *AssignmentExporter) SetReason(val values.Reason) {
	val.Export(func(v string) { ex.Reason = v })
}

func (ex *AssignmentExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
