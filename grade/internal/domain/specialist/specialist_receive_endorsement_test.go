package specialist

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/recognizer"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func TestSpecialistReceiveEndorsement(t *testing.T) {
	t1 := uuid.ParseSilent("63e8d541-af30-4593-a8ac-761dc268926d")
	t2 := uuid.ParseSilent("e2d9fcaa-565e-4295-9142-bd69e26581cf")
	m1 := uuid.ParseSilent("7c4435dc-6b5d-4628-a1f8-596dde6704b6")
	m2 := uuid.ParseSilent("c8858e26-6bc6-4775-a3bd-084773216b79")
	cases := []struct {
		RecogniserTenantId uuid.Uuid
		RecogniserMemberId uuid.Uuid
		RecognizerGrade    uint8
		SpecialistTenantId uuid.Uuid
		SpecialistMemberId uuid.Uuid
		SpecialistGrade    uint8
		ArtifactAuthorId   uuid.Uuid
		ArtifactTenantId   uuid.Uuid
		ExpectedError      error
	}{
		{t1, m1, 0, t1, m2, 0, m2, t1, nil},
		{t1, m1, 1, t1, m2, 0, m2, t1, nil},
		{t1, m1, 0, t1, m2, 1, m2, t1, ErrLowerGradeEndorses},
		{t1, m2, 0, t1, m2, 0, m2, t1, ErrEndorsementOneself},
		{t1, m1, 0, t2, m2, 0, m2, t1, ErrCrossTenantEndorsement},
		{t1, m1, 0, t1, m2, 0, m2, t2, ErrCrossTenantArtifact},
		{t1, m1, 0, t1, m2, 0, m1, t1, ErrNotAuthor},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			sf := NewSpecialistFakeFactory()
			rf := recognizer.NewRecognizerFakeFactory()
			af := artifact.NewArtifactFakeFactory()
			sf.Id.TenantId = c.SpecialistTenantId
			sf.Id.MemberId = c.SpecialistMemberId
			sf.Grade = c.SpecialistGrade
			rf.Id.TenantId = c.RecogniserTenantId
			rf.Id.MemberId = c.RecogniserMemberId
			rf.Grade = c.RecognizerGrade
			s, err := sf.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			r, err := rf.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			aId := sf.Id
			aId.MemberId = c.ArtifactAuthorId
			if err := af.AddAuthorId(aId); err != nil {
				t.Error(err)
				t.FailNow()
			}
			af.Id.TenantId = c.ArtifactTenantId
			af.Id.ArtifactId = sf.CurrentArtifactId
			a, err := af.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = r.ReserveEndorsement()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = s.ReceiveEndorsement(*r, *a, time.Now())
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
			sf := NewSpecialistFakeFactory()
			rf := recognizer.NewRecognizerFakeFactory()
			af := artifact.NewArtifactFakeFactory()
			if err := af.AddAuthorId(sf.Id); err != nil {
				t.Error(err)
				t.FailNow()
			}
			s, err := sf.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			r, err := rf.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			af.Id.TenantId = sf.Id.TenantId
			af.Id.ArtifactId = sf.CurrentArtifactId
			a, err := af.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = c.Prepare(r)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			err = s.ReceiveEndorsement(*r, *a, time.Now())
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}
