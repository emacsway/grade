package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/interfaces"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

func NewRecognizerFakeFactory() *RecognizerFakeFactory {
	return &RecognizerFakeFactory{
		1, 2, 1, 20, 1, time.Now(),
	}
}

type RecognizerFakeFactory struct {
	Id                        uint64
	MemberId                  uint64
	Grade                     uint8
	AvailableEndorsementCount uint8
	Version                   uint
	CreatedAt                 time.Time
}

func (f RecognizerFakeFactory) Create() (*Recognizer, error) {
	id, _ := recognizer.NewRecognizerId(f.Id)
	memberId, _ := external.NewMemberId(f.MemberId)
	grade, _ := shared.NewGrade(f.Grade)
	count, _ := recognizer.NewEndorsementCount(f.AvailableEndorsementCount)
	return NewRecognizer(id, memberId, grade, count, f.Version, f.CreatedAt)
}

func (f RecognizerFakeFactory) Export() RecognizerState {
	return RecognizerState{
		f.Id, f.MemberId, f.Grade, f.AvailableEndorsementCount, f.Version, f.CreatedAt,
	}
}

func (f RecognizerFakeFactory) ExportTo(ex interfaces.RecognizerExporter) {
	var id, memberId seedwork.Uint64Exporter
	var grade, availableEndorsementCount seedwork.Uint8Exporter

	id.SetState(f.Id)
	memberId.SetState(f.MemberId)
	grade.SetState(f.Grade)
	availableEndorsementCount.SetState(f.AvailableEndorsementCount)
	ex.SetState(
		&id, &memberId, &grade, &availableEndorsementCount, f.Version, f.CreatedAt,
	)
}
