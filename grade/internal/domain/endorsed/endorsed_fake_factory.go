package endorsed

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewEndorsedFakeFactory() (*EndorsedFakeFactory, error) {
	return &EndorsedFakeFactory{
		Id:                2,
		MemberId:          2,
		Grade:             0,
		CreatedAt:         time.Now(),
		CurrentArtifactId: 1000,
	}, nil
}

type EndorsedFakeFactory struct {
	Id                   uint64
	MemberId             uint64
	Grade                uint8
	ReceivedEndorsements []*EndorsementFakeFactory
	CreatedAt            time.Time
	CurrentArtifactId    uint64
}

func (f *EndorsedFakeFactory) achieveGrade() error {
	currentGrade := shared.WithoutGrade
	targetGrade, err := shared.NewGrade(f.Grade)
	if err != nil {
		return err
	}
	for currentGrade < targetGrade {
		r, err := recognizer.NewRecognizerFakeFactory()
		if err != nil {
			return err
		}
		r.Id = 1000
		recognizerGrade, _ := currentGrade.Next()
		r.Grade = recognizerGrade.Export()
		var endorsementCount uint = 0
		for !currentGrade.NextGradeAchieved(endorsementCount) {
			f.receiveEndorsement(r)
			endorsementCount += 2
		}
		currentGrade, err = currentGrade.Next()
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *EndorsedFakeFactory) ReceiveEndorsement(r *recognizer.RecognizerFakeFactory) error {
	err := f.achieveGrade()
	if err != nil {
		return err
	}
	f.receiveEndorsement(r)
	return nil
}

func (f *EndorsedFakeFactory) receiveEndorsement(r *recognizer.RecognizerFakeFactory) {
	e := NewEndorsementFakeFactory(r)
	e.ArtifactId = f.CurrentArtifactId
	f.CurrentArtifactId += 1
	e.CreatedAt = time.Now()
	f.ReceivedEndorsements = append(f.ReceivedEndorsements, e)
}

func (f EndorsedFakeFactory) Create() (*Endorsed, error) {
	err := f.achieveGrade()
	if err != nil {
		return nil, err
	}
	id, err := external.NewMemberId(f.Id)
	if err != nil {
		return nil, err
	}
	e, err := NewEndorsed(id, f.CreatedAt)
	if err != nil {
		return nil, err
	}
	for _, entf := range f.ReceivedEndorsements {
		r, err := entf.Recognizer.Create()
		if err != nil {
			return nil, err
		}
		artifactId, err := artifact.NewArtifactId(entf.ArtifactId)
		if err != nil {
			return nil, err
		}
		err = r.ReserveEndorsement()
		if err != nil {
			return nil, err
		}
		err = e.ReceiveEndorsement(*r, artifactId, entf.CreatedAt)
		if err != nil {
			return nil, err
		}
		e.IncreaseVersion()
	}
	return e, nil
}

func NewEndorsementFakeFactory(r *recognizer.RecognizerFakeFactory) *EndorsementFakeFactory {
	return &EndorsementFakeFactory{
		Recognizer: r,
		ArtifactId: 6,
		CreatedAt:  time.Now(),
	}
}

type EndorsementFakeFactory struct {
	Recognizer *recognizer.RecognizerFakeFactory
	ArtifactId uint64
	CreatedAt  time.Time
}
