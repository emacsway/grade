package endorser

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
)

var EndorserMemberIdFakeValue = uint(1004)

func NewEndorserFakeFactory() EndorserFakeFactory {
	idFactory := member.NewTenantMemberIdFakeFactory()
	idFactory.MemberId = EndorserMemberIdFakeValue
	return EndorserFakeFactory{
		Id:        idFactory,
		Grade:     1,
		CreatedAt: time.Now().Truncate(time.Microsecond),
	}
}

type EndorserFakeFactory struct {
	Id        member.TenantMemberIdFakeFactory
	Grade     uint8
	CreatedAt time.Time
}

func (f EndorserFakeFactory) Create() (*Endorser, error) {
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
