package recognizer

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

var RecognizerMemberIdFakeValue = uuid.ParseSilent("f917dc0a-5b9a-45a9-958d-22885a5e16a7")

func NewRecognizerFakeFactory() RecognizerFakeFactory {
	idFactory := member.NewTenantMemberIdFakeFactory()
	idFactory.MemberId = RecognizerMemberIdFakeValue
	return RecognizerFakeFactory{
		Id:        idFactory,
		Grade:     1,
		CreatedAt: time.Now(),
	}
}

type RecognizerFakeFactory struct {
	Id        member.TenantMemberIdFakeFactory
	Grade     uint8
	CreatedAt time.Time
}

func (f RecognizerFakeFactory) Create() (*Recognizer, error) {
	id, err := f.Id.Create()
	if err != nil {
		return nil, err
	}
	r, err := NewRecognizer(id, f.CreatedAt)
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
