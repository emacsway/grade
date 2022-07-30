package recognizer

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewRecognizerFakeFactory() *RecognizerFakeFactory {
	idFactory := member.NewTenantMemberIdFakeFactory()
	idFactory.MemberId = 1
	return &RecognizerFakeFactory{
		Id:        idFactory,
		Grade:     1,
		CreatedAt: time.Now(),
	}
}

type RecognizerFakeFactory struct {
	Id        *member.TenantMemberIdFakeFactory
	Grade     uint8
	CreatedAt time.Time
}

func (f RecognizerFakeFactory) Create() (*Recognizer, error) {
	id, err := member.NewTenantMemberId(f.Id.TenantId, f.Id.MemberId)
	if err != nil {
		return nil, err
	}
	r, err := NewRecognizer(id, f.CreatedAt)
	if err != nil {
		return nil, err
	}
	grade, err := shared.NewGrade(f.Grade)
	if err != nil {
		return nil, err
	}
	err = r.SetGrade(grade)
	if err != nil {
		return nil, err
	}
	return r, nil
}
