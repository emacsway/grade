package shared

import (
	"errors"
	"fmt"
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
