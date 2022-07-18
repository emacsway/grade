package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewRecognizer(
	id recognizer.RecognizerId,
	userId external.UserId,
	grade shared.Grade,
	availableEndorsementCount recognizer.AvailableEndorsementCount,
	version uint,
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
		Id:                        id,
		UserId:                    userId,
		Grade:                     grade,
		AvailableEndorsementCount: availableEndorsementCount,
		VersionedAggregate:        versioned,
		EventiveEntity:            eventive,
	}, nil
}

type Recognizer struct {
	Id                        recognizer.RecognizerId
	UserId                    external.UserId
	Grade                     shared.Grade
	AvailableEndorsementCount recognizer.AvailableEndorsementCount
	seedwork.VersionedAggregate
	seedwork.EventiveEntity
}
