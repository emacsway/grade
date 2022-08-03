package assignment

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork"
)

type AssignmentExporter struct {
	SpecialistId      member.TenantMemberIdExporter
	SpecialistVersion uint
	AssignedGrade     seedwork.Uint8Exporter
	Reason            seedwork.StringExporter
	CreatedAt         time.Time
}

func (ex *AssignmentExporter) SetSpecialistId(val member.TenantMemberId) {
	val.Export(&ex.SpecialistId)
}

func (ex *AssignmentExporter) SetSpecialistVersion(val uint) {
	ex.SpecialistVersion = val
}

func (ex *AssignmentExporter) SetAssignedGrade(val grade.Grade) {
	val.Export(&ex.AssignedGrade)
}

func (ex *AssignmentExporter) SetReason(val Reason) {
	val.Export(&ex.Reason)
}

func (ex *AssignmentExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
