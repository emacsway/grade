package endorsement

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
	"time"
)

type EndorsementExporter struct {
	RecognizerId      interfaces.Exporter[uint64]
	RecognizerGrade   interfaces.Exporter[uint8]
	RecognizerVersion uint
	EndorsedId        interfaces.Exporter[uint64]
	EndorsedGrade     interfaces.Exporter[uint8]
	EndorsedVersion   uint
	ArtifactId        interfaces.Exporter[uint64]
	CreatedAt         time.Time
}

func (ex *EndorsementExporter) SetState(
	recognizerId interfaces.Exporter[uint64],
	recognizerGrade interfaces.Exporter[uint8],
	recognizerVersion uint,
	endorsedId interfaces.Exporter[uint64],
	endorsedGrade interfaces.Exporter[uint8],
	endorsedVersion uint,
	artifactId interfaces.Exporter[uint64],
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
	RecognizerId      uint64
	RecognizerGrade   uint8
	RecognizerVersion uint
	EndorsedId        uint64
	EndorsedGrade     uint8
	EndorsedVersion   uint
	ArtifactId        uint64
	CreatedAt         time.Time
}
