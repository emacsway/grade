package endorsement

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
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
	recognizerId, err := member.NewMemberId(f.RecognizerId)
	if err != nil {
		return Endorsement{}, err
	}
	recognizerGrade, err := shared.NewGrade(f.RecognizerGrade)
	if err != nil {
		return Endorsement{}, err
	}
	endorsedId, err := member.NewMemberId(f.EndorsedId)
	if err != nil {
		return Endorsement{}, err
	}
	endorsedGrade, err := shared.NewGrade(f.EndorsedGrade)
	if err != nil {
		return Endorsement{}, err
	}
	artifactId, err := artifact.NewArtifactId(f.ArtifactId)
	if err != nil {
		return Endorsement{}, err
	}
	return NewEndorsement(
		recognizerId, recognizerGrade, f.RecognizerVersion,
		endorsedId, endorsedGrade, f.EndorsedVersion,
		artifactId, f.CreatedAt,
	)
}
