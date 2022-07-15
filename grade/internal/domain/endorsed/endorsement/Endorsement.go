package endorsement

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"time"
)

func NewEndorsement(
	recognizerId recognizer.RecognizerId,
	recognizerGrade domain.Grade,
	recognizerVersion uint,
	endorsedId endorsed.EndorsedId,
	endorsedGrade domain.Grade,
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
	RecognizerGrade   domain.Grade
	RecognizerVersion uint
	EndorsedId        endorsed.EndorsedId
	EndorsedGrade     domain.Grade
	EndorsedVersion   uint
	ArtifactId        artifact.ArtifactId
	CreatedAt         time.Time
}
