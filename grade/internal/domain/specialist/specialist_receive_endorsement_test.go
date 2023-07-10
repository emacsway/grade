package specialist

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	endorserVal "github.com/emacsway/grade/grade/internal/domain/endorser/values"
)

func TestSpecialistReceiveEndorsement(t *testing.T) {
	t1 := uint(10)
	m1 := uint(4)
	t2 := uint(11)
	m2 := uint(1005)
	cases := []struct {
		RecogniserTenantId uint
		RecogniserMemberId uint
		EndorserGrade      uint8
		SpecialistTenantId uint
		SpecialistMemberId uint
		SpecialistGrade    uint8
		ArtifactAuthorId   uint
		ArtifactTenantId   uint
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
			sf := NewSpecialistFaker()
			rf := endorser.NewEndorserFaker()
			af := artifact.NewArtifactFaker()
			sf.Id.TenantId = c.SpecialistTenantId
			sf.Id.MemberId = c.SpecialistMemberId
			sf.Grade = c.SpecialistGrade
			rf.Id.TenantId = c.RecogniserTenantId
			rf.Id.MemberId = c.RecogniserMemberId
			rf.Grade = c.EndorserGrade
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
			af.AddAuthorId(aId)
			af.Id.TenantId = c.ArtifactTenantId
			af.Next()
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
			err = s.ReceiveEndorsement(*r, *a, time.Now().Truncate(time.Microsecond))
			fmt.Println(err, c.ExpectedError)
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}

func TestSpecialistCanCompleteEndorsement(t *testing.T) {
	cases := []struct {
		Prepare       func(*endorser.Endorser) error
		ExpectedError error
	}{
		{func(e *endorser.Endorser) error {
			return e.ReserveEndorsement()
		}, nil},
		{func(e *endorser.Endorser) error {
			return nil
		}, endorser.ErrNoEndorsementReservation},
		{func(e *endorser.Endorser) error {
			for i := uint(0); i < endorserVal.YearlyEndorsementCount; i++ {
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
		}, endorser.ErrNoEndorsementReservation},
		{func(e *endorser.Endorser) error {
			err := e.ReserveEndorsement()
			if err != nil {
				return err
			}
			err = e.ReleaseEndorsementReservation()
			if err != nil {
				return err
			}
			return nil
		}, endorser.ErrNoEndorsementReservation},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			sf := NewSpecialistFaker()
			rf := endorser.NewEndorserFaker()
			af := artifact.NewArtifactFaker()
			af.AddAuthorId(sf.Id)
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
			af.Next()
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
			err = s.ReceiveEndorsement(*r, *a, time.Now().Truncate(time.Microsecond))
			assert.ErrorIs(t, err, c.ExpectedError)
		})
	}
}
