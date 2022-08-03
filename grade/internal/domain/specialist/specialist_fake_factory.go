package specialist

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewSpecialistFakeFactory() SpecialistFakeFactory {
	idFactory := member.NewTenantMemberIdFakeFactory()
	idFactory.MemberId = 2
	return SpecialistFakeFactory{
		Id:                idFactory,
		Grade:             0,
		CreatedAt:         time.Now(),
		CurrentArtifactId: 1000,
	}
}

type SpecialistFakeFactory struct {
	Id                   member.TenantMemberIdFakeFactory
	Grade                uint8
	ReceivedEndorsements []ReceivedEndorsementFakeFactory
	CreatedAt            time.Time
	CurrentArtifactId    uint64
}

func (f *SpecialistFakeFactory) achieveGrade() error {
	currentGrade, _ := grade.DefaultConstructor(0)
	targetGrade, err := grade.DefaultConstructor(f.Grade)
	if err != nil {
		return err
	}
	for currentGrade.LessThan(targetGrade) {
		r := recognizer.NewRecognizerFakeFactory()
		rId := member.NewTenantMemberIdFakeFactory()
		rId.TenantId = f.Id.TenantId
		rId.MemberId = 1000
		r.Id = rId
		recognizerGrade, _ := currentGrade.Next()
		gradeExporter := seedwork.Uint8Exporter(0)
		recognizerGrade.Export(&gradeExporter)
		r.Grade = uint8(gradeExporter)
		var endorsementCount uint = 0
		for !currentGrade.NextGradeAchieved(endorsementCount) {
			if err := f.receiveEndorsement(r); err != nil {
				return err
			}
			endorsementCount += 2
		}
		currentGrade, err = currentGrade.Next()
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *SpecialistFakeFactory) ReceiveEndorsement(r recognizer.RecognizerFakeFactory) error {
	err := f.achieveGrade()
	if err != nil {
		return err
	}
	return f.receiveEndorsement(r)
}

func (f *SpecialistFakeFactory) receiveEndorsement(r recognizer.RecognizerFakeFactory) error {
	entf := NewReceivedEndorsementFakeFactory(r)
	entf.Artifact.Id.TenantId = f.Id.TenantId
	entf.Artifact.Id.ArtifactId = f.CurrentArtifactId
	f.CurrentArtifactId += 1
	entf.CreatedAt = time.Now()
	if err := entf.Artifact.AddAuthorId(f.Id); err != nil {
		return err
	}
	f.ReceivedEndorsements = append(f.ReceivedEndorsements, entf)
	return nil
}

func (f SpecialistFakeFactory) Create() (*Specialist, error) {
	err := f.achieveGrade()
	if err != nil {
		return nil, err
	}
	id, err := member.NewTenantMemberId(f.Id.TenantId, f.Id.MemberId)
	if err != nil {
		return nil, err
	}
	e, err := NewSpecialist(id, f.CreatedAt)
	if err != nil {
		return nil, err
	}
	for i := range f.ReceivedEndorsements {
		r, err := f.ReceivedEndorsements[i].Recognizer.Create()
		if err != nil {
			return nil, err
		}
		art, err := f.ReceivedEndorsements[i].Artifact.Create()
		if err != nil {
			return nil, err
		}
		err = r.ReserveEndorsement()
		if err != nil {
			return nil, err
		}
		err = e.ReceiveEndorsement(*r, *art, f.ReceivedEndorsements[i].CreatedAt)
		if err != nil {
			return nil, err
		}
		e.IncreaseVersion()
	}
	return e, nil
}

func NewReceivedEndorsementFakeFactory(r recognizer.RecognizerFakeFactory) ReceivedEndorsementFakeFactory {
	artifactFactory := artifact.NewArtifactFakeFactory()
	artifactFactory.Id.ArtifactId = 6
	return ReceivedEndorsementFakeFactory{
		Recognizer: r,
		Artifact:   artifactFactory,
		CreatedAt:  time.Now(),
	}
}

type ReceivedEndorsementFakeFactory struct {
	Recognizer recognizer.RecognizerFakeFactory
	Artifact   artifact.ArtifactFakeFactory
	CreatedAt  time.Time
}