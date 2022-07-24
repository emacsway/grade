package recognizer

import (
	"errors"
	"fmt"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

const yearlyEndorsementCount = uint8(20)

var (
	ErrInvalidEndorsementCount = errors.New(fmt.Sprintf(
		"endorsement count should be between 0 and %d", yearlyEndorsementCount,
	))
)

func NewEndorsementCount(value uint8) (EndorsementCount, error) {
	if value > yearlyEndorsementCount {
		return EndorsementCount(0), ErrInvalidEndorsementCount
	}
	return EndorsementCount(value), nil
}

type EndorsementCount uint8

func (c EndorsementCount) HasAvailable() bool {
	return uint8(c) > uint8(0)
}

func (c EndorsementCount) Decrease() (EndorsementCount, error) {
	n, err := NewEndorsementCount(uint8(c) - uint8(1))
	if err != nil {
		return EndorsementCount(0), err
	}
	return n, nil
}

func (c EndorsementCount) Export() uint8 {
	return uint8(c)
}

func (c EndorsementCount) ExportTo(ex interfaces.Exporter[uint8]) {
	ex.SetState(uint8(c))
}
