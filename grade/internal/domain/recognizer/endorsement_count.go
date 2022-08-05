package recognizer

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/seedwork/identity"
)

const YearlyEndorsementCount = uint(20)

var (
	ErrInvalidEndorsementCount = fmt.Errorf(
		"endorsement count should be between 0 and %d", YearlyEndorsementCount,
	)
)

func NewEndorsementCount(value uint) (EndorsementCount, error) {
	if value > YearlyEndorsementCount {
		return EndorsementCount(0), ErrInvalidEndorsementCount
	}
	return EndorsementCount(value), nil
}

type EndorsementCount uint

func (c EndorsementCount) HasAvailable() bool {
	return uint(c) > uint(0)
}

func (c EndorsementCount) Decrease() (EndorsementCount, error) {
	n, err := NewEndorsementCount(uint(c) - uint(1))
	if err != nil {
		return EndorsementCount(0), err
	}
	return n, nil
}

func (c EndorsementCount) Export(ex identity.ExporterSetter[uint]) {
	ex.SetState(uint(c))
}
