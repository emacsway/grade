package endorsement

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork"
)

type EndorsementExporter struct {
	RecognizerId      member.TenantMemberIdExporter
	RecognizerGrade   seedwork.Uint8Exporter
	RecognizerVersion uint
	SpecialistId      member.TenantMemberIdExporter
	SpecialistGrade   seedwork.Uint8Exporter
	SpecialistVersion uint
	ArtifactId        artifact.TenantArtifactIdExporter
	CreatedAt         time.Time
}

func (ex *EndorsementExporter) SetRecognizerId(val member.TenantMemberId) {
	val.Export(&ex.RecognizerId)
}

func (ex *EndorsementExporter) SetRecognizerGrade(val grade.Grade) {
	val.Export(&ex.RecognizerGrade)
}

func (ex *EndorsementExporter) SetRecognizerVersion(val uint) {
	ex.RecognizerVersion = val
}

func (ex *EndorsementExporter) SetSpecialistId(val member.TenantMemberId) {
	val.Export(&ex.SpecialistId)
}

func (ex *EndorsementExporter) SetSpecialistGrade(val grade.Grade) {
	val.Export(&ex.SpecialistGrade)
}

func (ex *EndorsementExporter) SetSpecialistVersion(val uint) {
	ex.SpecialistVersion = val
}

func (ex *EndorsementExporter) SetArtifactId(val artifact.TenantArtifactId) {
	val.Export(&ex.ArtifactId)
}

func (ex *EndorsementExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
