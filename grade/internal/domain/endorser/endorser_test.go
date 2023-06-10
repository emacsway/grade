package endorser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
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
			err = r.CanCompleteEndorsement()
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}

func TestEndorserExport(t *testing.T) {
	var actualExporter EndorserExporter
	f := NewEndorserFakeFactory()
	agg, err := f.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg.Export(&actualExporter)
	assert.Equal(t, EndorserExporter{
		Id:                        member.NewTenantMemberIdExporter(f.Id.TenantId, f.Id.MemberId),
		Grade:                     exporters.Uint8Exporter(f.Grade),
		AvailableEndorsementCount: exporters.UintExporter(YearlyEndorsementCount),
		PendingEndorsementCount:   exporters.UintExporter(0),
		Version:                   0,
		CreatedAt:                 f.CreatedAt,
	}, actualExporter)
}
