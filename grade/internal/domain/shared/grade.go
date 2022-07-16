package shared

import "errors"

var (
	ErrInvalidGrade = errors.New("grade should be between 0 and 5")
)

func NewGrade(value uint8) (Grade, error) {
	if value < 0 || value > 5 {
		return Grade(0), ErrInvalidGrade
	}
	return Grade(value), nil
}

type Grade uint8

func (g Grade) HasNext() bool {
	return g < 5
}

func (g Grade) Next() (Grade, error) {
	nextGrade, err := NewGrade(uint8(g) + 1)
	if err != nil {
		return Grade(0), err
	}
	return nextGrade, nil
}
