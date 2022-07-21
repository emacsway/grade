package endorsement

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

func NewEndorsementFakeFactory() *EndorsementFakeFactory {
	return &EndorsementFakeFactory{
		1, 2, 3, 4, 1, 5, 6, time.Now(),
	}
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
	recognizerId, _ := recognizer.NewRecognizerId(f.RecognizerId)
	recognizerGrade, _ := shared.NewGrade(f.RecognizerGrade)
	endorsedId, _ := endorsed.NewEndorsedId(f.EndorsedId)
	endorsedGrade, _ := shared.NewGrade(f.EndorsedGrade)
	artifactId, _ := artifact.NewArtifactId(f.ArtifactId)
	return NewEndorsement(
		recognizerId, recognizerGrade, f.RecognizerVersion,
		endorsedId, endorsedGrade, f.EndorsedVersion,
		artifactId, f.CreatedAt,
	)
}

func (f EndorsementFakeFactory) CreateMemento() EndorsementMemento {
	return EndorsementMemento{
		f.RecognizerId, f.RecognizerGrade, f.RecognizerVersion,
		f.EndorsedId, f.EndorsedGrade, f.EndorsedVersion,
		f.ArtifactId, f.CreatedAt,
	}
}
