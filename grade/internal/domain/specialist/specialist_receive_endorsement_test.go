package specialist

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
)

func TestSpecialistReceiveEndorsement(t *testing.T) {
	cases := []struct {
		RecogniserTenantId uint64
		RecogniserMemberId uint64
		RecognizerGrade    uint8
		SpecialistTenantId uint64
		SpecialistMemberId uint64
		SpecialistGrade    uint8
		ArtifactAuthorId   uint64
		ArtifactTenantId   uint64
		ExpectedError      error
	}{
		{1, 1, 0, 1, 2, 0, 2, 1, nil},
		{1, 1, 1, 1, 2, 0, 2, 1, nil},
		{1, 1, 0, 1, 2, 1, 2, 1, ErrLowerGradeEndorses},
		{1, 3, 0, 1, 3, 0, 3, 1, ErrEndorsementOneself},
		{1, 1, 0, 2, 2, 0, 2, 1, ErrCrossTenantEndorsement},
		{1, 1, 0, 1, 2, 0, 2, 2, ErrCrossTenantArtifact},
		{1, 1, 0, 1, 2, 0, 3, 1, ErrNotAuthor},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			ef := NewSpecialistFakeFactory()
			rf := recognizer.NewRecognizerFakeFactory()
			af := artifact.NewArtifactFakeFactory()
			ef.Id.TenantId = c.SpecialistTenantId
			ef.Id.MemberId = c.SpecialistMemberId
			ef.Grade = c.SpecialistGrade
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
			aId := ef.Id
			aId.MemberId = c.ArtifactAuthorId
			if err := af.AddAuthorId(aId); err != nil {
				t.Error(err)
				t.FailNow()
			}
			af.Id.TenantId = c.ArtifactTenantId
			af.Id.ArtifactId = ef.CurrentArtifactId
			art, err := af.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = r.ReserveEndorsement()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = e.ReceiveEndorsement(*r, *art, time.Now())
			fmt.Println(err, c.ExpectedError)
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}

func TestSpecialistCanCompleteEndorsement(t *testing.T) {
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
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			ef := NewSpecialistFakeFactory()
			rf := recognizer.NewRecognizerFakeFactory()
			af := artifact.NewArtifactFakeFactory()
			if err := af.AddAuthorId(ef.Id); err != nil {
				t.Error(err)
				t.FailNow()
			}
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
			af.Id.TenantId = ef.Id.TenantId
			af.Id.ArtifactId = ef.CurrentArtifactId
			art, err := af.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = c.Prepare(r)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = e.ReceiveEndorsement(*r, *art, time.Now())
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}
