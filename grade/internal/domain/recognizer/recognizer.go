package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
	"time"
)

func NewRecognizer(
	id recognizer.RecognizerId,
	userId external.UserId,
	grade shared.Grade,
	availableEndorsementCount recognizer.AvailableEndorsementCount,
	version uint,
	createdAt time.Time,
) (*Recognizer, error) {
	versioned, err := seedwork.NewVersionedAggregate(version)
	if err != nil {
		return nil, err
	}
	eventive, err := seedwork.NewEventiveEntity()
	if err != nil {
		return nil, err
	}
	return &Recognizer{
		id:                        id,
		userId:                    userId,
		grade:                     grade,
		availableEndorsementCount: availableEndorsementCount,
		VersionedAggregate:        versioned,
		EventiveEntity:            eventive,
		CreatedAt:                 createdAt,
	}, nil
}

type Recognizer struct {
	id                        recognizer.RecognizerId
	userId                    external.UserId
	grade                     shared.Grade
	availableEndorsementCount recognizer.AvailableEndorsementCount
	CreatedAt                 time.Time
	seedwork.VersionedAggregate
	seedwork.EventiveEntity
}

func (r Recognizer) GetId() recognizer.RecognizerId {
	return r.id
}

func (r Recognizer) GetGrade() shared.Grade {
	return r.grade
}

func (r Recognizer) CreateMemento() RecognizerMemento {
	return RecognizerMemento{
		r.id.CreateMemento(),
		r.userId.CreateMemento(),
		r.grade.CreateMemento(),
		r.availableEndorsementCount.CreateMemento(),
		r.GetVersion(),
		r.CreatedAt,
	}
}

type RecognizerMemento struct {
	Id                        uint64
	UserId                    uint64
	Grade                     uint8
	AvailableEndorsementCount uint8
	Version                   uint
	CreatedAt                 time.Time
}
