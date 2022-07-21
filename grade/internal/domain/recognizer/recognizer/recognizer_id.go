package recognizer

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

func NewRecognizerId(value uint64) (RecognizerId, error) {
	id, err := seedwork.NewIdentity[uint64](value)
	if err != nil {
		return RecognizerId{}, err
	}
	return RecognizerId{id}, nil
}

type RecognizerId struct {
	seedwork.Identity[uint64, interfaces.Identity[uint64]]
}

func (id RecognizerId) ExportTo(ex interfaces.PrimitiveExporter[uint64]) {
	ex.SetState(id.Export())
}
