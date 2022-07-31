package expertisearea

import (
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func NewExpertiseAreaId(value uint64) (ExpertiseAreaId, error) {
	id, err := seedwork.NewUint64Identity(value)
	if err != nil {
		return ExpertiseAreaId{}, err
	}
	return ExpertiseAreaId{id}, nil
}

type ExpertiseAreaId struct {
	seedwork.Uint64Identity
}
