package recognizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func TestRecognizerCanCompleteEndorsement(t *testing.T) {
	cases := []struct {
		Prepare       func(*Recognizer) error
		ExpectedError error
	}{
		{func(r *Recognizer) error {
			return r.ReserveEndorsement()
		}, nil},
		{func(r *Recognizer) error {
			return nil
		}, ErrNoEndorsementReservation},
		{func(r *Recognizer) error {
			for i := uint8(0); i < YearlyEndorsementCount; i++ {
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
		}, ErrNoEndorsementReservation},
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
		}, ErrNoEndorsementReservation},
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
			err = r.CanCompleteEndorsement()
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}

func TestRecognizerExport(t *testing.T) {
	var actualExporter RecognizerExporter
	f := NewRecognizerFakeFactory()
	agg, err := f.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg.Export(&actualExporter)
	assert.Equal(t, RecognizerExporter{
		Id:                        member.NewTenantMemberIdExporter(f.Id.TenantId, f.Id.MemberId),
		Grade:                     seedwork.Uint8Exporter(f.Grade),
		AvailableEndorsementCount: seedwork.NewUint8Exporter(YearlyEndorsementCount),
		PendingEndorsementCount:   seedwork.NewUint8Exporter(0),
		Version:                   0,
		CreatedAt:                 f.CreatedAt,
	}, actualExporter)
}
