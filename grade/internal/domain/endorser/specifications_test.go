package endorser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndorserCanCompleteEndorsementSpecification(t *testing.T) {
	cases := []struct {
		Prepare        func(*Endorser) error
		ExpectedResult bool
	}{
		{func(e *Endorser) error {
			return e.ReserveEndorsement()
		}, true},
		{func(e *Endorser) error {
			return nil
		}, false},
		{func(e *Endorser) error {
			for i := uint(0); i < YearlyEndorsementCount; i++ {
				err := e.ReserveEndorsement()
				if err != nil {
					return err
				}
				err = e.CompleteEndorsement()
				if err != nil {
					return err
				}
			}
			return nil
		}, false},
		{func(e *Endorser) error {
			err := e.ReserveEndorsement()
			if err != nil {
				return err
			}
			err = e.ReleaseEndorsementReservation()
			if err != nil {
				return err
			}
			return nil
		}, false},
	}
	f := NewEndorserFakeFactory()
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
			sp := EndorserCanCompleteEndorsementSpecification{}
			result, err := sp.IsSatisfiedBy(*r)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			assert.Equal(t, c.ExpectedResult, result)
		})
	}
}
