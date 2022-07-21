package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

func NewEndorsedFakeFactory() *EndorsedFakeFactory {
	return &EndorsedFakeFactory{
		1, 2, 0, []*endorsement.EndorsementFakeFactory{}, 1, time.Now(), 5,
	}
}

type EndorsedFakeFactory struct {
	Id                   uint64
	UserId               uint64
	Grade                uint8
	ReceivedEndorsements []*endorsement.EndorsementFakeFactory
	Version              uint
	CreatedAt            time.Time
	CurrentArtifactId    uint64
}

func (f *EndorsedFakeFactory) AddReceivedEndorsement(r *recognizer.RecognizerFakeFactory) {
	e := endorsement.NewEndorsementFakeFactory()
	e.EndorsedId = f.Id
	e.EndorsedGrade = f.Grade
	e.EndorsedVersion = f.Version
	e.RecognizerId = r.Id
	e.RecognizerGrade = r.Grade
	e.RecognizerVersion = r.Version
	e.ArtifactId = f.CurrentArtifactId
	e.CreatedAt = time.Now()
	f.CurrentArtifactId += 1
	f.ReceivedEndorsements = append(f.ReceivedEndorsements, e)
}

func (f EndorsedFakeFactory) Create() (*Endorsed, error) {
	var receivedEndorsements []endorsement.Endorsement
	for _, v := range f.ReceivedEndorsements {
		e, err := v.Create()
		if err != nil {
			return nil, err
		}
		receivedEndorsements = append(receivedEndorsements, e)
	}
	id, _ := endorsed.NewEndorsedId(f.Id)
	userId, _ := external.NewUserId(f.UserId)
	grade, _ := shared.NewGrade(f.Grade)
	return NewEndorsed(id, userId, grade, receivedEndorsements, f.Version, f.CreatedAt)
}

func (f EndorsedFakeFactory) CreateMemento() EndorsedMemento {
	var receivedEndorsements []endorsement.EndorsementMemento
	for _, v := range f.ReceivedEndorsements {
		receivedEndorsements = append(receivedEndorsements, v.CreateMemento())
	}
	return EndorsedMemento{
		f.Id, f.UserId, f.Grade, receivedEndorsements, f.Version, f.CreatedAt,
	}
}
