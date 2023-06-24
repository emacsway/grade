package member

import (
	"time"
)

func NewMemberFakeFactory() MemberFakeFactory {
	return MemberFakeFactory{
		Id:        NewTenantMemberIdFakeFactory(),
		Status:    Active,
		FullName:  NewFullNameFakeFactory(),
		CreatedAt: time.Now().Truncate(time.Microsecond),
	}
}

type MemberFakeFactory struct {
	Id        TenantMemberIdFakeFactory
	Status    Status
	FullName  FullNameFakeFactory
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
