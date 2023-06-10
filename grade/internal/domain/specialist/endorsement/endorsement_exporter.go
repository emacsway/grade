package endorsement

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

type EndorsementExporter struct {
	EndorserId        member.TenantMemberIdExporter
	EndorserGrade     exporters.Uint8Exporter
	EndorserVersion   uint
	SpecialistId      member.TenantMemberIdExporter
	SpecialistGrade   exporters.Uint8Exporter
	SpecialistVersion uint
	ArtifactId        artifact.TenantArtifactIdExporter
	CreatedAt         time.Time
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

func (ex *EndorsementExporter) SetEndorserId(val member.TenantMemberId) {
	val.Export(&ex.EndorserId)
}

func (ex *EndorsementExporter) SetEndorserGrade(val grade.Grade) {
	val.Export(&ex.EndorserGrade)
}

func (ex *EndorsementExporter) SetEndorserVersion(val uint) {
	ex.EndorserVersion = val
}

func (ex *EndorsementExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
