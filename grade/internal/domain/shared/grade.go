package shared

import (
	"errors"
	"fmt"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"
)

const MaxGradeValue = uint8(5)

var (
	ErrInvalidGrade = errors.New(fmt.Sprintf("grade should be between 0 and %d", MaxGradeValue))
)

const (
	Expert       = Grade(5)
	Candidate    = Grade(4)
	Grade1       = Grade(3)
	Grade2       = Grade(2)
	Grade3       = Grade(1)
	WithoutGrade = Grade(0)
)

func NewGrade(value uint8) (Grade, error) {
	if value > MaxGradeValue {
		return Grade(0), ErrInvalidGrade
	}
	return Grade(value), nil
}

type Grade uint8

func (g Grade) HasNext() bool {
	return uint8(g) < MaxGradeValue
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
