package shared

func NewGrade(value uint8) (Grade, error) {
	return Grade(value), nil
}

type Grade uint8
