package endorsement

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

func NewEndorsement(
	recognizerId recognizer.RecognizerId,
	recognizerGrade shared.Grade,
	recognizerVersion uint,
	endorsedId endorsed.EndorsedId,
	endorsedGrade shared.Grade,
	endorsedVersion uint,
	artifactId artifact.ArtifactId,
	createdAt time.Time,
) (Endorsement, error) {
	return Endorsement{
		RecognizerId:      recognizerId,
		RecognizerGrade:   recognizerGrade,
		RecognizerVersion: recognizerVersion,
		EndorsedId:        endorsedId,
		EndorsedGrade:     endorsedGrade,
		EndorsedVersion:   endorsedVersion,
		ArtifactId:        artifactId,
		CreatedAt:         createdAt,
	}, nil
}

type Endorsement struct {
	RecognizerId      recognizer.RecognizerId
	RecognizerGrade   shared.Grade
	RecognizerVersion uint
	EndorsedId        endorsed.EndorsedId
	EndorsedGrade     shared.Grade
	EndorsedVersion   uint
	ArtifactId        artifact.ArtifactId
	CreatedAt         time.Time
}
