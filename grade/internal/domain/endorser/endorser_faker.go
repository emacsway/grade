package endorser

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
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

func NewEndorserFaker(opts ...EndorserFakerOption) *EndorserFaker {
	idFactory := member.NewTenantMemberIdFaker()
	idFactory.MemberId = EndorserMemberIdFakeValue
	f := &EndorserFaker{
		Id:         idFactory,
		Grade:      1,
		CreatedAt:  time.Now().Truncate(time.Microsecond),
		Repository: EndorserDummyRepository{},
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type EndorserFaker struct {
	Id         member.TenantMemberIdFaker
	Grade      uint8
	CreatedAt  time.Time
	Repository EndorserRepository
}

func (f EndorserFaker) Create() (*Endorser, error) {
	id, err := f.Id.Create()
	if err != nil {
		return nil, err
	}
	r, err := NewEndorser(id, f.CreatedAt)
	if err != nil {
		return nil, err
	}
	g, err := grade.DefaultConstructor(f.Grade)
	if err != nil {
		return nil, err
	}
	err = r.SetGrade(g)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type EndorserRepository interface {
	Insert(*Endorser) error
}

type EndorserDummyRepository struct{}

func (r EndorserDummyRepository) Insert(agg *Endorser) error {
	return nil
}
