package endorser

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

var EndorserMemberIdFakeValue = uint(1004)

func NewEndorserFaker() *EndorserFaker {
	idFactory := member.NewTenantMemberIdFaker()
	idFactory.MemberId = EndorserMemberIdFakeValue
	return &EndorserFaker{
		Id:        idFactory,
		Grade:     1,
		CreatedAt: time.Now().Truncate(time.Microsecond),
	}
}

type EndorserFaker struct {
	Id        member.TenantMemberIdFaker
	Grade     uint8
	CreatedAt time.Time
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
