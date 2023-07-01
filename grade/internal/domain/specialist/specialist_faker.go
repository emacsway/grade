package specialist

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

var SpecialistMemberIdFakeValue = memberVal.MemberIdFakeValue

type SpecialistFakerOption func(*SpecialistFaker)

func WithTenantId(tenantId uint) SpecialistFakerOption {
	return func(f *SpecialistFaker) {
		f.Id.TenantId = tenantId
	}
}

func WithMemberId(memberId uint) SpecialistFakerOption {
	return func(f *SpecialistFaker) {
		f.Id.MemberId = memberId
	}
}

func WithRepository(repo SpecialistRepository) SpecialistFakerOption {
	return func(f *SpecialistFaker) {
		f.Repository = repo
	}
}

func NewSpecialistFaker(opts ...SpecialistFakerOption) *SpecialistFaker {
	idFactory := memberVal.NewTenantMemberIdFaker()
	idFactory.MemberId = SpecialistMemberIdFakeValue
	f := &SpecialistFaker{
		Id:         idFactory,
		Grade:      0,
		CreatedAt:  time.Now().Truncate(time.Microsecond),
		Dependency: SpecialistDependencyFakerImp{},
		Repository: SpecialistDummyRepository{},
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type SpecialistFaker struct {
	Id                   memberVal.TenantMemberIdFaker
	Grade                uint8
	ReceivedEndorsements []ReceivedEndorsementFakeItem
	CreatedAt            time.Time
	Dependency           SpecialistDependencyFaker
	Repository           SpecialistRepository
}

func (f *SpecialistFaker) achieveGrade() error {
	currentGrade, _ := grade.DefaultConstructor(0)
	targetGrade, err := grade.DefaultConstructor(f.Grade)
	if err != nil {
		return err
	}
	for currentGrade.LessThan(targetGrade) {
		rf := f.Dependency.MakeEndorserFaker()
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

func (f *SpecialistFaker) ReceiveEndorsement(e *endorser.EndorserFaker) error {
	err := f.achieveGrade()
	if err != nil {
		return err
	}
	return f.receiveEndorsement(e)
}

func (f *SpecialistFaker) receiveEndorsement(e *endorser.EndorserFaker) error {
	entf := f.Dependency.MakeReceivedEndorsementFakeItem(e)
	entf.Artifact.Id.TenantId = f.Id.TenantId
	if len(f.ReceivedEndorsements) > 0 {
		entf.Artifact.Id = f.ReceivedEndorsements[len(f.ReceivedEndorsements)-1].Artifact.Id
	}
	err := entf.Artifact.Next()
	if err != nil {
		return err
	}
	entf.CreatedAt = time.Now().Truncate(time.Microsecond)
	if err := entf.Artifact.AddAuthorId(f.Id); err != nil {
		return err
	}
	f.ReceivedEndorsements = append(f.ReceivedEndorsements, entf)
	return nil
}

func (f SpecialistFaker) Create() (*Specialist, error) {
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

type SpecialistDependencyFaker interface {
	MakeReceivedEndorsementFakeItem(*endorser.EndorserFaker) ReceivedEndorsementFakeItem
	MakeArtifactFaker() *artifact.ArtifactFaker
	MakeEndorserFaker() *endorser.EndorserFaker
}

type SpecialistDependencyFakerImp struct{}

func (d SpecialistDependencyFakerImp) MakeReceivedEndorsementFakeItem(
	e *endorser.EndorserFaker,
) ReceivedEndorsementFakeItem {
	return ReceivedEndorsementFakeItem{
		Endorser:  e,
		Artifact:  d.MakeArtifactFaker(),
		CreatedAt: time.Now().Truncate(time.Microsecond),
	}
}

func (d SpecialistDependencyFakerImp) MakeArtifactFaker() *artifact.ArtifactFaker {
	return artifact.NewArtifactFaker()
}

func (d SpecialistDependencyFakerImp) MakeEndorserFaker() *endorser.EndorserFaker {
	return endorser.NewEndorserFaker()
}

type ReceivedEndorsementFakeItem struct {
	Endorser  *endorser.EndorserFaker
	Artifact  *artifact.ArtifactFaker
	CreatedAt time.Time
}

type SpecialistRepository interface {
	Insert(*Specialist) error
}

type SpecialistDummyRepository struct{}

func (r SpecialistDummyRepository) Insert(agg *Specialist) error {
	return nil
}
