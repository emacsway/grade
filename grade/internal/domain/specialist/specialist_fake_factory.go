package specialist

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

var SpecialistMemberIdFakeValue = member.MemberIdFakeValue

func NewSpecialistFakeFactory() SpecialistFakeFactory {
	idFactory := member.NewTenantMemberIdFakeFactory()
	idFactory.MemberId = SpecialistMemberIdFakeValue
	return SpecialistFakeFactory{
		Id:        idFactory,
		Grade:     0,
		CreatedAt: time.Now(),
	}
}

type SpecialistFakeFactory struct {
	Id                   member.TenantMemberIdFakeFactory
	Grade                uint8
	ReceivedEndorsements []ReceivedEndorsementFakeFactory
	CreatedAt            time.Time
}

func (f *SpecialistFakeFactory) achieveGrade() error {
	currentGrade, _ := grade.DefaultConstructor(0)
	targetGrade, err := grade.DefaultConstructor(f.Grade)
	if err != nil {
		return err
	}
	for currentGrade.LessThan(targetGrade) {
		rf := endorser.NewEndorserFakeFactory()
		rf.Id.TenantId = f.Id.TenantId
		endorserGrade, _ := currentGrade.Next()
		gradeExporter := exporters.Uint8Exporter(0)
		endorserGrade.Export(&gradeExporter)
		rf.Grade = uint8(gradeExporter)
		var endorsementCount uint = 0
		for !currentGrade.NextGradeAchieved(endorsementCount) {
			if err := f.receiveEndorsement(rf); err != nil {
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

func (f *SpecialistFakeFactory) ReceiveEndorsement(e endorser.EndorserFakeFactory) error {
	err := f.achieveGrade()
	if err != nil {
		return err
	}
	return f.receiveEndorsement(e)
}

func (f *SpecialistFakeFactory) receiveEndorsement(e endorser.EndorserFakeFactory) error {
	entf := NewReceivedEndorsementFakeFactory(e)
	entf.Artifact.Id.TenantId = f.Id.TenantId
	if len(f.ReceivedEndorsements) > 0 {
		entf.Artifact.Id = f.ReceivedEndorsements[len(f.ReceivedEndorsements)-1].Artifact.Id
	}
	entf.Artifact.Id.NextArtifactId()
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
	id, err := f.Id.Create()
	if err != nil {
		return nil, err
	}
	s, err := NewSpecialist(id, f.CreatedAt)
	if err != nil {
		return nil, err
	}
	for i := range f.ReceivedEndorsements {
		r, err := f.ReceivedEndorsements[i].Endorser.Create()
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
		err = s.ReceiveEndorsement(*r, *art, f.ReceivedEndorsements[i].CreatedAt)
		if err != nil {
			return nil, err
		}
		s.SetVersion(s.Version() + 1)
	}
	return s, nil
}

func NewReceivedEndorsementFakeFactory(e endorser.EndorserFakeFactory) ReceivedEndorsementFakeFactory {
	artifactFactory := artifact.NewArtifactFakeFactory()
	artifactFactory.Id.NextArtifactId()
	return ReceivedEndorsementFakeFactory{
		Endorser:  e,
		Artifact:  artifactFactory,
		CreatedAt: time.Now(),
	}
}

type ReceivedEndorsementFakeFactory struct {
	Endorser  endorser.EndorserFakeFactory
	Artifact  artifact.ArtifactFakeFactory
	CreatedAt time.Time
}
