package endorsed

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewEndorsedFakeFactory() *EndorsedFakeFactory {
	idFactory := member.NewTenantMemberIdFakeFactory()
	idFactory.MemberId = 2
	return &EndorsedFakeFactory{
		Id:                idFactory,
		Grade:             0,
		CreatedAt:         time.Now(),
		CurrentArtifactId: 1000,
	}
}

type EndorsedFakeFactory struct {
	Id                   *member.TenantMemberIdFakeFactory
	Grade                uint8
	ReceivedEndorsements []*ReceivedEndorsementFakeFactory
	CreatedAt            time.Time
	CurrentArtifactId    uint64
}

func (f *EndorsedFakeFactory) achieveGrade() error {
	currentGrade, _ := shared.DefaultConstructor(0)
	targetGrade, err := shared.DefaultConstructor(f.Grade)
	if err != nil {
		return err
	}
	for currentGrade.LessThan(targetGrade) {
		r := recognizer.NewRecognizerFakeFactory()
		rId := member.NewTenantMemberIdFakeFactory()
		rId.MemberId = 1000
		r.Id = rId
		recognizerGrade, _ := currentGrade.Next()
		gradeExporter := seedwork.Uint8Exporter(0)
		recognizerGrade.Export(&gradeExporter)
		r.Grade = uint8(gradeExporter)
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
	e := NewReceivedEndorsementFakeFactory(r)
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
	id, err := member.NewTenantMemberId(f.Id.TenantId, f.Id.MemberId)
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

func NewReceivedEndorsementFakeFactory(r *recognizer.RecognizerFakeFactory) *ReceivedEndorsementFakeFactory {
	return &ReceivedEndorsementFakeFactory{
		Recognizer: r,
		ArtifactId: 6,
		CreatedAt:  time.Now(),
	}
}

type ReceivedEndorsementFakeFactory struct {
	Recognizer *recognizer.RecognizerFakeFactory
	ArtifactId uint64
	CreatedAt  time.Time
}
