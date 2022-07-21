package external

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewUserId(value uint64) (UserId, error) {
	id, err := seedwork.NewIdentity[uint64](value)
	if err != nil {
		return UserId{}, err
	}
	return UserId{id}, nil
}

type UserId struct {
	seedwork.Identity[uint64, interfaces.Identity[uint64]]
}

func (id UserId) ExportTo(ex interfaces.PrimitiveExporter[uint64]) {
	ex.SetState(id.Export())
}
