package endorsed

import (
	"fmt"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEndorsedReceiveEndorsement(t *testing.T) {
	cases := []struct {
		RecogniserId    uint64
		RecognizerGrade uint8
		EndorsedId      uint64
		EndorsedGrade   uint8
		ExpectedError   error
	}{
		{1, 0, 2, 0, nil},
		{1, 1, 2, 0, nil},
		{1, 0, 2, 1, endorsement.ErrLowerGradeEndorses},
		{1, 0, 1, 0, endorsement.ErrEndorsementOneself},
	}
	ef := NewEndorsedFakeFactory()
	rf := recognizer.NewRecognizerFakeFactory()
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			ef.Id = c.EndorsedId
			ef.Grade = c.EndorsedGrade
			rf.Id = c.RecogniserId
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
