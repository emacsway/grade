package endorser

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
)

var EndorserMemberIdFakeValue = uint(1004)

type EndorserFakerOption func(*EndorserFaker)

func WithTenantId(tenantId uint) EndorserFakerOption {
	return func(f *EndorserFaker) {
		f.Id.TenantId = tenantId
	}
}

func WithMemberId(memberId uint) EndorserFakerOption {
	return func(f *EndorserFaker) {
		f.Id.MemberId = memberId
	}
}

func WithRepository(repo EndorserRepository) EndorserFakerOption {
	return func(f *EndorserFaker) {
		f.Repository = repo
	}
}

func WithMemberFaker(memberFaker *member.MemberFaker) EndorserFakerOption {
	return func(f *EndorserFaker) {
		f.MemberFaker = memberFaker
	}
}

func NewEndorserFaker(opts ...EndorserFakerOption) *EndorserFaker {
	f := &EndorserFaker{
		Id:          memberVal.NewMemberIdFaker(),
		Repository:  EndorserDummyRepository{},
		MemberFaker: member.NewMemberFaker(),
	}
	f.Id.MemberId = EndorserMemberIdFakeValue
	f.fake()
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type EndorserFaker struct {
	Id          memberVal.MemberIdFaker
	Grade       uint8
	CreatedAt   time.Time
	Repository  EndorserRepository
	MemberFaker *member.MemberFaker
	agg         *Endorser
}

func (f *EndorserFaker) fake() {
	f.Grade = 1
	f.CreatedAt = time.Now().Truncate(time.Microsecond)
}

func (f *EndorserFaker) Next() error {
	f.fake()
	f.MemberFaker.Next()
	err := f.BuildDependencies()
	if err != nil {
		return err
	}
	f.agg = nil
	return nil
}

func (f *EndorserFaker) Create() (*Endorser, error) {
	if f.agg != nil {
		return f.agg, nil
	}

	id, err := f.Id.Create()
	if err != nil {
		return nil, err
	}
	agg, err := NewEndorser(id, f.CreatedAt)
	if err != nil {
		return nil, err
	}
	g, err := grade.DefaultConstructor(f.Grade)
	if err != nil {
		return nil, err
	}
	err = agg.SetGrade(g)
	if err != nil {
		return nil, err
	}
	err = f.Repository.Insert(agg)
	if err != nil {
		return nil, err
	}
	f.agg = agg
	return agg, nil
}

// unidirectional flow of changes
func (f *EndorserFaker) SetTenantId(val uint) {
	f.MemberFaker.SetTenantId(val)
}

func (f *EndorserFaker) SetMemberId(val uint) {
	f.MemberFaker.SetMemberId(val)
}

func (f *EndorserFaker) SetId(id memberVal.MemberIdFaker) {
	f.SetTenantId(id.TenantId)
	f.SetMemberId(id.MemberId)
}

func (f *EndorserFaker) BuildDependencies() (err error) {
	err = f.MemberFaker.BuildDependencies()
	if err != nil {
		return err
	}
	_, err = f.MemberFaker.Create() // Use repo if it is needed to get the instance.
	if err != nil {
		return err
	}
	f.Id = f.MemberFaker.Id
	return err
}

type EndorserRepository interface {
	Insert(*Endorser) error
}

type EndorserDummyRepository struct {
}

func (r EndorserDummyRepository) Insert(agg *Endorser) error {
	return nil
}
