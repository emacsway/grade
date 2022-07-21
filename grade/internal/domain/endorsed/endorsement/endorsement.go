package endorsement

import (
	"errors"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

var ErrHigherGradeEndorsed = errors.New("it is allowed to endorse only members with equal or lower grade")

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
	if recognizerGrade < endorsedGrade {
		return Endorsement{}, ErrHigherGradeEndorsed
	}
	return Endorsement{
		recognizerId:      recognizerId,
		recognizerGrade:   recognizerGrade,
		recognizerVersion: recognizerVersion,
		endorsedId:        endorsedId,
		endorsedGrade:     endorsedGrade,
		endorsedVersion:   endorsedVersion,
		artifactId:        artifactId,
		createdAt:         createdAt,
	}, nil
}

type Endorsement struct {
	recognizerId      recognizer.RecognizerId
	recognizerGrade   shared.Grade
	recognizerVersion uint
	endorsedId        endorsed.EndorsedId
	endorsedGrade     shared.Grade
	endorsedVersion   uint
	artifactId        artifact.ArtifactId
	createdAt         time.Time
}

func (e Endorsement) GetRecognizerId() recognizer.RecognizerId {
	return e.recognizerId
}

func (e Endorsement) GetArtifactId() artifact.ArtifactId {
	return e.artifactId
}

func (e Endorsement) Export() EndorsementState {
	return EndorsementState{
		e.recognizerId.Export(), e.recognizerGrade.Export(), e.recognizerVersion,
		e.endorsedId.Export(), e.endorsedGrade.Export(), e.endorsedVersion,
		e.artifactId.Export(), e.createdAt,
	}
}

func (e Endorsement) ExportTo(ex interfaces.EndorsementExporter) {
	var recognizerId, endorsedId, artifactId seedwork.Uint64Exporter
	var recognizerGrade, endorsedGrade seedwork.Uint8Exporter

	e.recognizerId.ExportTo(&recognizerId)
	e.recognizerGrade.ExportTo(&recognizerGrade)
	e.endorsedId.ExportTo(&endorsedId)
	e.endorsedGrade.ExportTo(&endorsedGrade)
	e.artifactId.ExportTo(&artifactId)
	ex.SetState(
		&recognizerId, &recognizerGrade, e.recognizerVersion,
		&endorsedId, &endorsedGrade, e.endorsedVersion,
		&artifactId, e.createdAt,
	)
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
