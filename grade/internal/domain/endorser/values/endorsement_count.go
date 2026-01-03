package values

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/seedwork/domain/specification"
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

func (c EndorsementCount) Equal(other specification.EqualOperand) bool {
	otherTyped, ok := other.(EndorsementCount)
	if !ok {
		return false
	}
	return uint(c) == uint(otherTyped)
}

func (c EndorsementCount) GreaterThanEqual(other specification.GreaterThanEqualOperand) bool {
	otherTyped, ok := other.(EndorsementCount)
	if !ok {
		return false
	}
	return uint(c) >= uint(otherTyped)
}

func (c EndorsementCount) Decrease() (EndorsementCount, error) {
	n, err := NewEndorsementCount(uint(c) - uint(1))
	if err != nil {
		return EndorsementCount(0), err
	}
	return n, nil
}

func (c EndorsementCount) Export(ex func(uint)) {
	ex(uint(c))
}
