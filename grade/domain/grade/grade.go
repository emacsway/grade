package grade

import (
	"errors"
	"time"
)

const Begginer = 0

type Recommendation struct {
	recommended  Member
	recommending Member
}

type Testimonial struct {
	supporter Member
	promoted  Member
}

type Member struct {
	grade Grade
}

func (m Member) Recommend(other Member, timeNow time.Time) (Testimonial, error) {
	if m.grade.LT(other.grade) {
		return Testimonial{}, errors.New("cannot support elders")
	}

	return Testimonial{
		supporter: m,
		promoted:  other,
	}, nil
}

type Grade int

func (g Grade) LT(other Grade) bool {
	return g < other
}
