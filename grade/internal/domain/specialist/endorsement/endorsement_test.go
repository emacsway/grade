package endorsement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	artifact "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func TestEndorsementIsEndorsedBy(t *testing.T) {
	r1 := uint(10)
	a1 := uint(3)
	r2 := uint(11)
	a2 := uint(4)
	cases := []struct {
		RecogniserId     uint
		ArtifactId       uint
		TestRecogniserId uint
		TestArtifactId   uint
		ExpectedResult   bool
	}{
		{r1, a2, r1, a2, true},
		{r1, a2, r2, a2, false},
		{r1, a2, r1, a1, false},
	}
	f := NewEndorsementFakeFactory()
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			f.EndorserId.MemberId = c.RecogniserId
			rId, err := member.NewTenantMemberId(f.EndorserId.TenantId, c.TestRecogniserId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			f.ArtifactId.ArtifactId = c.ArtifactId
			aId, err := artifact.NewTenantArtifactId(f.ArtifactId.TenantId, c.TestArtifactId)
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
				f.EndorserGrade = j
				f.SpecialistGrade = i
				ent, err := f.Create()
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				if j > i {
					expectedWeight = 2
				}
				assert.Equal(t, expectedWeight, ent.Weight())
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
		SpecialistId:      member.NewTenantMemberIdExporter(f.SpecialistId.TenantId, f.SpecialistId.MemberId),
		SpecialistGrade:   exporters.Uint8Exporter(f.SpecialistGrade),
		SpecialistVersion: f.SpecialistVersion,
		ArtifactId:        artifact.NewTenantArtifactIdExporter(f.ArtifactId.TenantId, f.ArtifactId.ArtifactId),
		EndorserId:        member.NewTenantMemberIdExporter(f.EndorserId.TenantId, f.EndorserId.MemberId),
		EndorserGrade:     exporters.Uint8Exporter(f.EndorserGrade),
		EndorserVersion:   f.EndorserVersion,
		CreatedAt:         f.CreatedAt,
	}, actualExporter)
}
