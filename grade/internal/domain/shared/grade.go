package shared

import (
	"errors"
	"fmt"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

const MaxGradeValue = uint8(5)

var (
	ErrInvalidGrade = errors.New(fmt.Sprintf("grade should be between 0 and %d", MaxGradeValue))
)

var DefaultConstructor = NewGradeFactory(MaxGradeValue, GradeMatrix)

var GradeMatrix = map[uint8]uint{
	0: 6,
	1: 10,
	2: 14,
	3: 20,
	4: 40,
}

type GradeConstructor func(uint8) (Grade, error)

type Grade struct {
	value                     uint8
	nextGradeEndorsementCount uint
	constructor               func(uint8) (Grade, error)
}

func NewGradeFactory(maxGradeValue uint8, matrix map[uint8]uint) GradeConstructor {
	var constructor GradeConstructor
	constructor = func(value uint8) (Grade, error) {
		if value > maxGradeValue {
			return Grade{value: 0}, ErrInvalidGrade
		}
		g := Grade{
			value:                     value,
			nextGradeEndorsementCount: matrix[value],
		}
		g.constructor = constructor
		return g, nil

	}

	return constructor
}

func (g Grade) NextGradeAchieved(endorsementCount uint) bool {
	return endorsementCount >= g.nextGradeEndorsementCount
}

func (g Grade) Next() (Grade, error) {
	nextGrade, err := g.constructor(g.value + 1)
	if err != nil {
		return g, err
	}
	return nextGrade, nil
}

func (g Grade) Previous() (Grade, error) {
	previousGrade, err := g.constructor(g.value - 1)
	if err != nil {
		return g, err
	}
	return previousGrade, nil
}

func (g Grade) LessThan(other Grade) bool {
	return g.value < other.value
}

func (g Grade) GreaterThan(other Grade) bool {
	return g.value > other.value
}

func (g Grade) Equal(other Grade) bool {
	return g.value == other.value
}

func (g Grade) Export(ex seedwork.ExporterSetter[uint8]) {
	ex.SetState(g.value)
}
