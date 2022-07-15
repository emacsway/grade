package domain

func NewGrade(value uint8) Grade {
	return Grade(value)
}

type Grade uint8
