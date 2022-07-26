package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

func NewRecognizerFakeFactory() (*RecognizerFakeFactory, error) {
	return &RecognizerFakeFactory{
		Id:        1,
		Grade:     1,
		CreatedAt: time.Now(),
	}, nil
}

type RecognizerFakeFactory struct {
	Id        uint64
	Grade     uint8
	CreatedAt time.Time
}

func (f RecognizerFakeFactory) Create() (*Recognizer, error) {
	id, err := member.NewMemberId(f.Id)
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
