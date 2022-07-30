package expertisearea

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewExpertiseAreaId(value uint64) (ExpertiseAreaId, error) {
	id, err := seedwork.NewIdentity[uint64](value)
	if err != nil {
		return ExpertiseAreaId{}, err
	}
	return ExpertiseAreaId{id}, nil
}

type ExpertiseAreaId struct {
	seedwork.Identity[uint64, seedwork.Identifier[uint64], seedwork.ExporterSetter[uint64]]
}
