package assignment

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

type AssignmentExporter struct {
	EndorsedId      member.TenantMemberIdExporter
	EndorsedVersion uint
	AssignedGrade   seedwork.Uint8Exporter
	Reason          seedwork.StringExporter
	CreatedAt       time.Time
}

func (ex *AssignmentExporter) SetEndorsedId(val member.TenantMemberId) {
	val.Export(&ex.EndorsedId)
}

func (ex *AssignmentExporter) SetEndorsedVersion(val uint) {
	ex.EndorsedVersion = val
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
