package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsed"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/external"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

func NewEndorsed(
	id endorsed.EndorsedId,
	userId external.UserId,
	grade shared.Grade,
	endorsements []endorsement.Endorsement,
	version uint,
) (*Endorsed, error) {
	return &Endorsed{
		Id:                   id,
		UserId:               userId,
		Grade:                grade,
		ReceivedEndorsements: endorsements,
		Version:              version,
	}, nil
}

type Endorsed struct {
	Id                   endorsed.EndorsedId
	UserId               external.UserId
	Grade                shared.Grade
	ReceivedEndorsements []endorsement.Endorsement
	Version              uint
}
