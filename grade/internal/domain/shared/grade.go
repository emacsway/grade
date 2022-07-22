package shared

import (
	"errors"
	"fmt"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

const maxGradeValue = uint8(5)

var (
	ErrInvalidGrade = errors.New(fmt.Sprintf("grade should be between 0 and %d", maxGradeValue))
)

func NewGrade(value uint8) (Grade, error) {
	if value > maxGradeValue {
		return Grade(0), ErrInvalidGrade
	}
	return Grade(value), nil
}

type Grade uint8

func (g Grade) HasNext() bool {
	return uint8(g) < maxGradeValue
}

func (g Grade) Next() (Grade, error) {
	nextGrade, err := NewGrade(uint8(g) + 1)
	if err != nil {
		return Grade(0), err
	}
	return nextGrade, nil
}

func (g Grade) Export() uint8 {
	return uint8(g)
}

func (g *Grade) Import(value uint8) {
	*g = Grade(value)
}

func (g Grade) ExportTo(ex interfaces.Exporter[uint8]) {
	ex.SetState(uint8(g))
}
