package endorser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/endorser/values"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

func TestEndorserCanCompleteEndorsement(t *testing.T) {
	cases := []struct {
		Prepare       func(*Endorser) error
		ExpectedError error
	}{
		{func(e *Endorser) error {
			return e.ReserveEndorsement()
		}, nil},
		{func(e *Endorser) error {
			return nil
		}, ErrNoEndorsementReservation},
		{func(e *Endorser) error {
			for i := uint(0); i < values.YearlyEndorsementCount; i++ {
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
		}, ErrNoEndorsementReservation},
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
		}, ErrNoEndorsementReservation},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			f := NewEndorserFaker()
			e, err := f.Create(nil)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = c.Prepare(e)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = e.CanCompleteEndorsement()
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}

func TestEndorserExport(t *testing.T) {
	var actualExporter EndorserExporter
	f := NewEndorserFaker()
	agg, err := f.Create(nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg.Export(&actualExporter)
	assert.Equal(t, EndorserExporter{
		Id:                        member.NewMemberIdExporter(f.Id.TenantId, f.Id.MemberId),
		Grade:                     f.Grade,
		AvailableEndorsementCount: values.YearlyEndorsementCount,
		PendingEndorsementCount:   0,
		Version:                   0,
		CreatedAt:                 f.CreatedAt,
	}, actualExporter)
}
