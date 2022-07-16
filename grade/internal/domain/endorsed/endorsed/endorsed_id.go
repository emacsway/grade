package endorsed

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewEndorsedId(value uint64) (EndorsedId, error) {
	id, err := seedwork.NewIdentity[uint64](value)
	if err != nil {
		return EndorsedId{}, err
	}
	return EndorsedId{id}, nil
}

type EndorsedId struct {
	seedwork.Identity[uint64, interfaces.Identity[uint64]]
}
