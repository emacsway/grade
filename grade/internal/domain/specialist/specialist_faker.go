package specialist

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
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

func WithArtifactFaker(artifactFaker *artifact.ArtifactFaker) SpecialistFakerOption {
	return func(f *SpecialistFaker) {
		f.ArtifactFaker = artifactFaker
	}
}

func WithEndorserFaker(endorserFaker *endorser.EndorserFaker) SpecialistFakerOption {
	return func(f *SpecialistFaker) {
		f.EndorserFaker = endorserFaker
	}
}

func WithMemberFaker(memberFaker *member.MemberFaker) SpecialistFakerOption {
	return func(f *SpecialistFaker) {
		f.MemberFaker = memberFaker
	}
}

func NewSpecialistFaker(opts ...SpecialistFakerOption) *SpecialistFaker {
	idFactory := memberVal.NewTenantMemberIdFaker()
	idFactory.MemberId = SpecialistMemberIdFakeValue
	f := &SpecialistFaker{
		Id:            idFactory,
		Repository:    SpecialistDummyRepository{},
		ArtifactFaker: artifact.NewArtifactFaker(),
		EndorserFaker: endorser.NewEndorserFaker(),
		MemberFaker:   member.NewMemberFaker(),
	}
	f.fake()
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
	Repository           SpecialistRepository
	ArtifactFaker        *artifact.ArtifactFaker
	EndorserFaker        *endorser.EndorserFaker
	MemberFaker          *member.MemberFaker
	agg                  *Specialist
}

func (f *SpecialistFaker) fake() {
	f.Grade = 0
	f.CreatedAt = time.Now().Truncate(time.Microsecond)
}

func (f *SpecialistFaker) Next() error {
	f.fake()
	f.MemberFaker.Next()
	err := f.BuildDependencies()
	if err != nil {
		return err
	}
	f.agg = nil
	return nil
}

func (f *SpecialistFaker) achieveGrade() error {
	currentGrade, _ := grade.DefaultConstructor(0)
	targetGrade, err := grade.DefaultConstructor(f.Grade)
	if err != nil {
		return err
	}
	for currentGrade.LessThan(targetGrade) {
		ef := f.MakeEndorserFaker()
		ef.Id.TenantId = f.Id.TenantId
		endorserGrade, _ := currentGrade.Next()
		gradeExporter := exporters.Uint8Exporter(0)
		endorserGrade.Export(&gradeExporter)
		ef.Grade = uint8(gradeExporter)
		var endorsementCount uint = 0
		for !currentGrade.NextGradeAchieved(endorsementCount) {
			if err := f.receiveEndorsement(ef); err != nil {
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
	entf, err := f.MakeReceivedEndorsementFakeItem(e)
	if err != nil {
		return err
	}
	entf.Artifact.Id.TenantId = f.Id.TenantId
	entf.CreatedAt = time.Now().Truncate(time.Microsecond)
	entf.Artifact.AddAuthorId(f.Id)
	_, err = entf.Artifact.Create()
	if err != nil {
		return err
	}
	f.ReceivedEndorsements = append(f.ReceivedEndorsements, entf)
	return nil
}

func (f SpecialistFaker) Create() (*Specialist, error) {
	if f.agg != nil {
		return f.agg, nil
	}
	err := f.achieveGrade()
	if err != nil {
		return nil, err
	}
	id, err := f.Id.Create()
	if err != nil {
		return nil, err
	}
	agg, err := NewSpecialist(id, f.CreatedAt)
	if err != nil {
		return nil, err
	}
	for i := range f.ReceivedEndorsements {
		e, err := f.ReceivedEndorsements[i].Endorser.Create()
		if err != nil {
			return nil, err
		}
		art, err := f.ReceivedEndorsements[i].Artifact.Create()
		if err != nil {
			return nil, err
		}
		err = e.ReserveEndorsement()
		if err != nil {
			return nil, err
		}
		err = agg.ReceiveEndorsement(*e, *art, f.ReceivedEndorsements[i].CreatedAt)
		if err != nil {
			return nil, err
		}
		agg.SetVersion(agg.Version() + 1)
	}
	err = f.Repository.Insert(agg)
	if err != nil {
		return nil, err
	}
	f.agg = agg
	return agg, nil
}

func (f *SpecialistFaker) BuildDependencies() (err error) {
	err = f.ArtifactFaker.BuildDependencies()
	if err != nil {
		return err
	}
	_, err = f.ArtifactFaker.Create()
	if err != nil {
		return err
	}
	*f.EndorserFaker.MemberFaker = *f.ArtifactFaker.MemberFaker
	*f.MemberFaker = *f.ArtifactFaker.MemberFaker
	f.Id = f.MemberFaker.Id
	return err
}

func (f *SpecialistFaker) MakeReceivedEndorsementFakeItem(
	e *endorser.EndorserFaker,
) (ReceivedEndorsementFakeItem, error) {
	err := f.ArtifactFaker.Next()
	if err != nil {
		return ReceivedEndorsementFakeItem{}, err
	}
	return ReceivedEndorsementFakeItem{
		Endorser:  e,
		Artifact:  *f.ArtifactFaker,
		CreatedAt: time.Now().Truncate(time.Microsecond),
	}, nil
}

func (f SpecialistFaker) MakeEndorserFaker() *endorser.EndorserFaker {
	return endorser.NewEndorserFaker()
}

type ReceivedEndorsementFakeItem struct {
	Endorser  *endorser.EndorserFaker
	Artifact  artifact.ArtifactFaker
	CreatedAt time.Time
}

type SpecialistRepository interface {
	Insert(*Specialist) error
}

type SpecialistDummyRepository struct{}

func (r SpecialistDummyRepository) Insert(agg *Specialist) error {
	return nil
}
