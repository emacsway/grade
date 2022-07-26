package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

func NewRecognizerFakeFactory() (*RecognizerFakeFactory, error) {
	return &RecognizerFakeFactory{
		Id:        1,
		MemberId:  1,
		Grade:     1,
		CreatedAt: time.Now(),
	}, nil
}

type RecognizerFakeFactory struct {
	Id        uint64
	MemberId  uint64
	Grade     uint8
	CreatedAt time.Time
}

func (f RecognizerFakeFactory) Create() (*Recognizer, error) {
	id, err := recognizer.NewRecognizerId(f.Id)
	if err != nil {
		return nil, err
	}
	memberId, err := external.NewMemberId(f.MemberId)
	if err != nil {
		return nil, err
	}
	grade, err := shared.NewGrade(f.Grade)
	if err != nil {
		return nil, err
	}
	return NewRecognizer(id, memberId, grade, f.CreatedAt)
}
