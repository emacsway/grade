package endorsement

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewEndorsementFakeFactory() (*EndorsementFakeFactory, error) {
	return &EndorsementFakeFactory{
		RecognizerId:      1,
		RecognizerGrade:   2,
		RecognizerVersion: 3,
		EndorsedId:        4,
		EndorsedGrade:     1,
		EndorsedVersion:   5,
		ArtifactId:        6,
		CreatedAt:         time.Now(),
	}, nil
}

type EndorsementFakeFactory struct {
	RecognizerId      uint64
	RecognizerGrade   uint8
	RecognizerVersion uint
	EndorsedId        uint64
	EndorsedGrade     uint8
	EndorsedVersion   uint
	ArtifactId        uint64
	CreatedAt         time.Time
}

func (f EndorsementFakeFactory) Create() (Endorsement, error) {
	recognizerId, _ := external.NewMemberId(f.RecognizerId)
	recognizerGrade, _ := shared.NewGrade(f.RecognizerGrade)
	endorsedId, _ := external.NewMemberId(f.EndorsedId)
	endorsedGrade, _ := shared.NewGrade(f.EndorsedGrade)
	artifactId, _ := artifact.NewArtifactId(f.ArtifactId)
	return NewEndorsement(
		recognizerId, recognizerGrade, f.RecognizerVersion,
		endorsedId, endorsedGrade, f.EndorsedVersion,
		artifactId, f.CreatedAt,
	)
}
