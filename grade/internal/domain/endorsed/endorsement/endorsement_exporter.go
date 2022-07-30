package endorsement

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

type EndorsementExporter struct {
	RecognizerId      member.TenantMemberIdExporterSetter
	RecognizerGrade   seedwork.ExporterSetter[uint8]
	RecognizerVersion uint
	EndorsedId        member.TenantMemberIdExporterSetter
	EndorsedGrade     seedwork.ExporterSetter[uint8]
	EndorsedVersion   uint
	ArtifactId        seedwork.ExporterSetter[uint64]
	CreatedAt         time.Time
}

func (ex *EndorsementExporter) SetState(
	recognizerId member.TenantMemberIdExporterSetter,
	recognizerGrade seedwork.ExporterSetter[uint8],
	recognizerVersion uint,
	endorsedId member.TenantMemberIdExporterSetter,
	endorsedGrade seedwork.ExporterSetter[uint8],
	endorsedVersion uint,
	artifactId seedwork.ExporterSetter[uint64],
	createdAt time.Time,
) {
	ex.RecognizerId = recognizerId
	ex.RecognizerGrade = recognizerGrade
	ex.RecognizerVersion = recognizerVersion
	ex.EndorsedId = endorsedId
	ex.EndorsedGrade = endorsedGrade
	ex.EndorsedVersion = endorsedVersion
	ex.ArtifactId = artifactId
	ex.CreatedAt = createdAt
}

type EndorsementState struct {
	RecognizerId      member.TenantMemberIdState
	RecognizerGrade   uint8
	RecognizerVersion uint
	EndorsedId        member.TenantMemberIdState
	EndorsedGrade     uint8
	EndorsedVersion   uint
	ArtifactId        uint64
	CreatedAt         time.Time
}
