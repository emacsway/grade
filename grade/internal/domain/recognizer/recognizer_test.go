package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRecognizerCreateMemento(t *testing.T) {
	f := NewRecognizerTestFactory()
	agg, _ := f.Create()
	assert.Equal(t, f.CreateMemento(), agg.CreateMemento())
}

func NewRecognizerTestFactory() *RecognizerTestFactory {
	return &RecognizerTestFactory{
		1, 2, 0, 20, 1, time.Now(),
	}
}

type RecognizerTestFactory struct {
	Id                        uint64
	UserId                    uint64
	Grade                     uint8
	AvailableEndorsementCount uint8
	Version                   uint
	CreatedAt                 time.Time
}

func (f RecognizerTestFactory) Create() (*Recognizer, error) {
	id, _ := recognizer.NewRecognizerId(f.Id)
	userId, _ := external.NewUserId(f.UserId)
	grade, _ := shared.NewGrade(f.Grade)
	count, _ := recognizer.NewAvailableEndorsementCount(f.AvailableEndorsementCount)
	return NewRecognizer(id, userId, grade, count, f.Version, f.CreatedAt)
}

func (f RecognizerTestFactory) CreateMemento() RecognizerMemento {
	return RecognizerMemento{
		f.Id, f.UserId, f.Grade, f.AvailableEndorsementCount, f.Version, f.CreatedAt,
	}
}
