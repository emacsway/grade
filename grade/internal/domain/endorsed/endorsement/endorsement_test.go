package endorsement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func TestEndorsementIsEndorsedBy(t *testing.T) {
	cases := []struct {
		RecogniserId     uint64
		ArtifactId       uint64
		TestRecogniserId uint64
		TestArtifactId   uint64
		ExpectedResult   bool
	}{
		{1, 2, 1, 2, true},
		{1, 2, 2, 2, false},
		{1, 2, 1, 1, false},
	}
	f := NewEndorsementFakeFactory()
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			f.RecognizerId.MemberId = c.RecogniserId
			rId, err := member.NewTenantMemberId(f.RecognizerId.TenantId, c.TestRecogniserId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			f.ArtifactId = c.ArtifactId
			aId, err := artifact.NewArtifactId(c.TestArtifactId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			ent, err := f.Create()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			assert.Equal(t, c.ExpectedResult, ent.IsEndorsedBy(rId, aId))
		})
	}
}

func TestEndorsementWeight(t *testing.T) {
	f := NewEndorsementFakeFactory()
	for i := uint8(0); i <= grade.MaxGradeValue; i++ {
		for j := i; j <= grade.MaxGradeValue; j++ {
			t.Run(fmt.Sprintf("Case i=%d j=%d", i, j), func(t *testing.T) {
				var expectedWeight Weight = 1
				f.RecognizerGrade = j
				f.EndorsedGrade = i
				ent, err := f.Create()
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				if j > i {
					expectedWeight = 2
				}
				assert.Equal(t, expectedWeight, ent.GetWeight())
			})
		}
	}
}

func TestEndorsementExport(t *testing.T) {
	var actualExporter EndorsementExporter
	f := NewEndorsementFakeFactory()
	agg, _ := f.Create()
	agg.Export(&actualExporter)
	assert.Equal(t, EndorsementExporter{
		RecognizerId:      member.NewTenantMemberIdExporter(f.RecognizerId.TenantId, f.RecognizerId.MemberId),
		RecognizerGrade:   seedwork.Uint8Exporter(f.RecognizerGrade),
		RecognizerVersion: f.RecognizerVersion,
		EndorsedId:        member.NewTenantMemberIdExporter(f.EndorsedId.TenantId, f.EndorsedId.MemberId),
		EndorsedGrade:     seedwork.Uint8Exporter(f.EndorsedGrade),
		EndorsedVersion:   f.EndorsedVersion,
		ArtifactId:        seedwork.Uint64Exporter(f.ArtifactId),
		CreatedAt:         f.CreatedAt,
	}, actualExporter)
}
