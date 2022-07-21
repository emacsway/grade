package endorsement

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"time"
)

type EndorsementExporter struct {
	RecognizerId      seedwork.Uint64Exporter
	RecognizerGrade   seedwork.Uint8Exporter
	RecognizerVersion uint
	EndorsedId        seedwork.Uint64Exporter
	EndorsedGrade     seedwork.Uint8Exporter
	EndorsedVersion   uint
	ArtifactId        seedwork.Uint64Exporter
	CreatedAt         time.Time
}

func (ex *EndorsementExporter) SetState(
	recognizerId seedwork.Uint64Exporter,
	recognizerGrade seedwork.Uint8Exporter,
	recognizerVersion uint,
	endorsedId seedwork.Uint64Exporter,
	endorsedGrade seedwork.Uint8Exporter,
	endorsedVersion uint,
	artifactId seedwork.Uint64Exporter,
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
