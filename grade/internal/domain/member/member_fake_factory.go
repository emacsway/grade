package member

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/member/values"
)

func NewMemberFakeFactory() MemberFakeFactory {
	return MemberFakeFactory{
		Id:        values.NewTenantMemberIdFakeFactory(),
		Status:    values.Active,
		FullName:  values.NewFullNameFakeFactory(),
		CreatedAt: time.Now().Truncate(time.Microsecond),
	}
}

type MemberFakeFactory struct {
	Id        values.TenantMemberIdFakeFactory
	Status    values.Status
	FullName  values.FullNameFakeFactory
	CreatedAt time.Time
}

func (f MemberFakeFactory) Create() (*Member, error) {
	id, err := f.Id.Create()
	if err != nil {
		return nil, err
	}
	fullName, err := f.FullName.Create()
	if err != nil {
		return nil, err
	}
	return NewMember(
		id, f.Status, fullName, f.CreatedAt,
	), nil
}
