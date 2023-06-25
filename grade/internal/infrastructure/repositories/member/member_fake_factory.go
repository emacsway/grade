package member

import (
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewMemberFakeFactory(
	session infrastructure.DbSession,
	opts ...member.MemberFakeFactoryOption,
) *MemberFakeFactory {
	opts = append(opts, member.WithTransientId())
	f := &MemberFakeFactory{
		member.NewMemberFakeFactory(opts...),
		NewMemberRepository(session),
	}
	return f
}

type MemberFakeFactory struct {
	member.MemberFakeFactory
	// Repo and dependecies should be at Aggregate-level FakeFactory, not at TenantMemberIdFakeFactory
	Repository *MemberRepository
}

func (f *MemberFakeFactory) Create() (*member.Member, error) {
	var aggExp member.MemberExporter
	agg, err := f.MemberFakeFactory.Create()
	if err != nil {
		return nil, err
	}
	err = f.Repository.Insert(agg)
	if err != nil {
		return nil, err
	}
	agg.Export(&aggExp)
	f.Id.MemberId = uint(aggExp.Id.MemberId)
	return agg, nil
}
