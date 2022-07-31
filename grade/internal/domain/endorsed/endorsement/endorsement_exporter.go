package endorsement

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

type EndorsementExporter struct {
	RecognizerId      member.TenantMemberIdExporter
	RecognizerGrade   seedwork.ExporterSetter[uint8]
	RecognizerVersion uint
	EndorsedId        member.TenantMemberIdExporter
	EndorsedGrade     seedwork.ExporterSetter[uint8]
	EndorsedVersion   uint
	ArtifactId        seedwork.ExporterSetter[uint64]
	CreatedAt         time.Time
}

func (ex *EndorsementExporter) SetState(
	recognizerGrade seedwork.ExporterSetter[uint8],
	recognizerVersion uint,
	endorsedGrade seedwork.ExporterSetter[uint8],
	endorsedVersion uint,
	artifactId seedwork.ExporterSetter[uint64],
	createdAt time.Time,
) {
	ex.RecognizerGrade = recognizerGrade
	ex.RecognizerVersion = recognizerVersion
	ex.EndorsedGrade = endorsedGrade
	ex.EndorsedVersion = endorsedVersion
	ex.ArtifactId = artifactId
	ex.CreatedAt = createdAt
}

func (ex *EndorsementExporter) SetRecognizerId(recognizerId member.TenantMemberId) {
	recognizerId.Export(&ex.RecognizerId)
}

func (ex *EndorsementExporter) SetEndorsedId(endorsedId member.TenantMemberId) {
	endorsedId.Export(&ex.EndorsedId)
}
