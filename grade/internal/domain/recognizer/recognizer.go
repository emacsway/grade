package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewRecognizer(
	id recognizer.RecognizerId,
	userId external.UserId,
	grade shared.Grade,
	availableEndorsementCount recognizer.AvailableEndorsementCount,
	version uint,
) (*Recognizer, error) {
	return &Recognizer{
		Id:                        id,
		UserId:                    userId,
		Grade:                     grade,
		AvailableEndorsementCount: availableEndorsementCount,
		Version:                   version,
	}, nil
}

type Recognizer struct {
	Id                        recognizer.RecognizerId
	UserId                    external.UserId
	Grade                     shared.Grade
	AvailableEndorsementCount recognizer.AvailableEndorsementCount
	Version                   uint
}
