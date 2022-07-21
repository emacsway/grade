package endorsement

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndorsementConstructor(t *testing.T) {
	cases := []struct {
		RecognizerGrade uint8
		EndorsedGrade   uint8
		ExpectedError   error
	}{
		{0, 0, nil},
		{1, 0, nil},
		{0, 1, ErrHigherGradeEndorsed},
	}
	f := NewEndorsementFakeFactory()
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			f.RecognizerGrade = c.RecognizerGrade
			f.EndorsedGrade = c.EndorsedGrade
			e, err := f.Create()
			assert.Equal(t, f.RecognizerGrade, c.RecognizerGrade)
			assert.Equal(t, c.ExpectedError, err)
			if err == nil {
				assert.Equal(t, f.CreateMemento(), e.CreateMemento())
			}
		})
	}
}

func TestEndorsementCreateMemento(t *testing.T) {
	f := NewEndorsementFakeFactory()
	e, _ := f.Create()
	assert.Equal(t, f.CreateMemento(), e.CreateMemento())
}
