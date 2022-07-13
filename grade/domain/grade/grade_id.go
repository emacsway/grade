package grade

import "github.com/emacsway/qualifying-grade/grade/domain/seedwork/interfaces"

func NewGradeId(value int) GradeId {
	return GradeId{value: value}
}

type GradeId struct {
	value int
}

func (gradeId GradeId) Equals(other interfaces.Identity[int]) bool {
	return gradeId.value == other.GetValue()
}

func (gradeId GradeId) GetValue() int {
	return gradeId.value
}
