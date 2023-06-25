package member

import (
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/infrastructure"
)

func NewMemberFakeFactory(
	session infrastructure.DbSession,
	opts ...member.MemberFakeFactoryOption,
) MemberFakeFactory {
	opts = append(opts, member.WithTransientId())
	return MemberFakeFactory{
		member.NewMemberFakeFactory(opts...),
		NewMemberRepository(session),
	}
}

type MemberFakeFactory struct {
	member.MemberFakeFactory
	Repository *MemberRepository
}

func (f MemberFakeFactory) Create() (*member.Member, error) {
	obj, err := f.MemberFakeFactory.Create()
	if err != nil {
		return nil, err
	}
	err = f.Repository.Insert(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
