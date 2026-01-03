package specialist

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
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
	idFactory := memberVal.NewMemberIdFaker()
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
	Id            memberVal.MemberIdFaker
	Grade         uint8
	Commands      []interface{}
	CreatedAt     time.Time
	Repository    SpecialistRepository
	ArtifactFaker *artifact.ArtifactFaker
	EndorserFaker *endorser.EndorserFaker
	MemberFaker   *member.MemberFaker
	agg           *Specialist
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

func (f *SpecialistFaker) ReceiveEndorsement(ef *endorser.EndorserFaker) error {
	err := f.achieveGrade()
	if err != nil {
		return err
	}
	return f.receiveEndorsement(ef)
}

func (f *SpecialistFaker) achieveGrade() error {
	currentGrade, _ := grade.DefaultConstructor(0)
	targetGrade, err := grade.DefaultConstructor(f.Grade)
	if err != nil {
		return err
	}
	for currentGrade.LessThan(targetGrade) {
		err = f.EndorserFaker.Next()
		if err != nil {
			return err
		}
		ef := f.EndorserFaker
		// TODO: Remove me: ef.Id.TenantId = f.Id.TenantId
		endorserGrade, _ := currentGrade.Next()
		var gradeExporter uint8
		endorserGrade.Export(func(v uint8) { gradeExporter = v })
		ef.Grade = gradeExporter
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

func (f *SpecialistFaker) receiveEndorsement(ef *endorser.EndorserFaker) error {
	if _, err := ef.Create(); err != nil {
		return err
	}

	if err := f.ArtifactFaker.Next(); err != nil {
		return err
	}
	af := f.ArtifactFaker
	// TODO: Remove me: af.Id.TenantId = f.Id.TenantId
	af.AddAuthorId(f.Id)
	if _, err := af.Create(); err != nil {
		return err
	}

	f.Commands = append(f.Commands, ReceivedEndorsementFakeCommand{
		Endorser:  *ef,
		Artifact:  *af,
		CreatedAt: time.Now().Truncate(time.Microsecond),
	})
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
	for i := range f.Commands {
		switch cmd := f.Commands[i].(type) {
		case ReceivedEndorsementFakeCommand:
			e, err := cmd.Endorser.Create()
			if err != nil {
				return nil, err
			}
			art, err := cmd.Artifact.Create()
			if err != nil {
				return nil, err
			}
			err = e.ReserveEndorsement()
			if err != nil {
				return nil, err
			}
			err = agg.ReceiveEndorsement(*e, *art, cmd.CreatedAt)
			if err != nil {
				return nil, err
			}
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

// unidirectional flow of changes
func (f *SpecialistFaker) SetTenantId(val uint) {
	f.ArtifactFaker.SetTenantId(val)
}

func (f *SpecialistFaker) SetMemberId(val uint) {
	f.ArtifactFaker.SetMemberId(val)
}

func (f *SpecialistFaker) SetId(id memberVal.MemberIdFaker) {
	f.SetTenantId(id.TenantId)
	f.SetMemberId(id.MemberId)
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

type ReceivedEndorsementFakeCommand struct {
	Endorser  endorser.EndorserFaker
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
