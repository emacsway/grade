package competence

import (
	"github.com/emacsway/grade/grade/internal/domain/seedwork"
)

func NewCompetenceId(value uint64) (CompetenceId, error) {
	id, err := seedwork.NewUint64Identity(value)
	if err != nil {
		return CompetenceId{}, err
	}
	return CompetenceId{id}, nil
}

type CompetenceId struct {
	seedwork.Uint64Identity
}
