package endorsement

import (
	"time"

	artifact "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

type EndorsementExporter struct {
	EndorserId        member.MemberIdExporter
	EndorserGrade     uint8
	EndorserVersion   uint
	SpecialistId      member.MemberIdExporter
	SpecialistGrade   uint8
	SpecialistVersion uint
	ArtifactId        artifact.ArtifactIdExporter
	CreatedAt         time.Time
}

func (ex *EndorsementExporter) SetSpecialistId(val member.MemberId) {
	val.Export(&ex.SpecialistId)
}

func (ex *EndorsementExporter) SetSpecialistGrade(val grade.Grade) {
	val.Export(func(v uint8) { ex.SpecialistGrade = v })
}

func (ex *EndorsementExporter) SetSpecialistVersion(val uint) {
	ex.SpecialistVersion = val
}

func (ex *EndorsementExporter) SetArtifactId(val artifact.ArtifactId) {
	val.Export(&ex.ArtifactId)
}

func (ex *EndorsementExporter) SetEndorserId(val member.MemberId) {
	val.Export(&ex.EndorserId)
}

func (ex *EndorsementExporter) SetEndorserGrade(val grade.Grade) {
	val.Export(func(v uint8) { ex.EndorserGrade = v })
}

func (ex *EndorsementExporter) SetEndorserVersion(val uint) {
	ex.EndorserVersion = val
}

func (ex *EndorsementExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
