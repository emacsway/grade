package recognizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecognizerCanCompleteEndorsementSpecification(t *testing.T) {
	cases := []struct {
		Prepare        func(*Recognizer) error
		ExpectedResult bool
	}{
		{func(r *Recognizer) error {
			return r.ReserveEndorsement()
		}, true},
		{func(r *Recognizer) error {
			return nil
		}, false},
		{func(r *Recognizer) error {
			for i := uint(0); i < YearlyEndorsementCount; i++ {
				err := r.ReserveEndorsement()
				if err != nil {
					return err
				}
				err = r.CompleteEndorsement()
				if err != nil {
					return err
				}
			}
			return nil
		}, false},
		{func(r *Recognizer) error {
			err := r.ReserveEndorsement()
			if err != nil {
				return err
			}
			err = r.ReleaseEndorsementReservation()
			if err != nil {
				return err
			}
			return nil
		}, false},
	}
	f := NewRecognizerFakeFactory()
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			r, err := f.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = c.Prepare(r)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			sp := RecognizerCanCompleteEndorsementSpecification{}
			result, err := sp.IsSatisfiedBy(*r)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			assert.Equal(t, c.ExpectedResult, result)
		})
	}
}
