package endorsed

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
)

func TestEndorsedReceiveEndorsement(t *testing.T) {
	cases := []struct {
		RecogniserTenantId uint64
		RecogniserMemberId uint64
		RecognizerGrade    uint8
		EndorsedTenantId   uint64
		EndorsedMemberId   uint64
		EndorsedGrade      uint8
		ExpectedError      error
	}{
		{1, 1, 0, 1, 2, 0, nil},
		{1, 1, 1, 1, 2, 0, nil},
		{1, 1, 0, 1, 2, 1, ErrLowerGradeEndorses},
		{1, 3, 0, 1, 3, 0, ErrEndorsementOneself},
		{1, 1, 0, 2, 2, 0, ErrCrossTenantEndorsement},
	}
	ef := NewEndorsedFakeFactory()
	rf := recognizer.NewRecognizerFakeFactory()
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			ef.Id.TenantId = c.EndorsedTenantId
			ef.Id.MemberId = c.EndorsedMemberId
			ef.Grade = c.EndorsedGrade
			rf.Id.TenantId = c.RecogniserTenantId
			rf.Id.MemberId = c.RecogniserMemberId
			rf.Grade = c.RecognizerGrade
			e, err := ef.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			r, err := rf.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			artifactId, err := artifact.NewArtifactId(ef.CurrentArtifactId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = r.ReserveEndorsement()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = e.ReceiveEndorsement(*r, artifactId, time.Now())
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}

func TestEndorsedCanCompleteEndorsement(t *testing.T) {
	cases := []struct {
		Prepare       func(*recognizer.Recognizer) error
		ExpectedError error
	}{
		{func(r *recognizer.Recognizer) error {
			return r.ReserveEndorsement()
		}, nil},
		{func(r *recognizer.Recognizer) error {
			return nil
		}, recognizer.ErrNoEndorsementReservation},
		{func(r *recognizer.Recognizer) error {
			for i := uint(0); i < recognizer.YearlyEndorsementCount; i++ {
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
		}, recognizer.ErrNoEndorsementReservation},
		{func(r *recognizer.Recognizer) error {
			err := r.ReserveEndorsement()
			if err != nil {
				return err
			}
			err = r.ReleaseEndorsementReservation()
			if err != nil {
				return err
			}
			return nil
		}, recognizer.ErrNoEndorsementReservation},
	}
	ef := NewEndorsedFakeFactory()
	rf := recognizer.NewRecognizerFakeFactory()
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			e, err := ef.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			r, err := rf.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			artifactId, err := artifact.NewArtifactId(ef.CurrentArtifactId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = c.Prepare(r)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = e.ReceiveEndorsement(*r, artifactId, time.Now())
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}
