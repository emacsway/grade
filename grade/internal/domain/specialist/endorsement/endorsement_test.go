package endorsement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func TestEndorsementIsEndorsedBy(t *testing.T) {
	r1 := uuid.ParseSilent("63e8d541-af30-4593-a8ac-761dc268926d")
	a1 := uuid.ParseSilent("7c4435dc-6b5d-4628-a1f8-596dde6704b6")
	r2 := uuid.ParseSilent("e2d9fcaa-565e-4295-9142-bd69e26581cf")
	a2 := uuid.ParseSilent("c8858e26-6bc6-4775-a3bd-084773216b79")
	cases := []struct {
		RecogniserId     uuid.Uuid
		ArtifactId       uuid.Uuid
		TestRecogniserId uuid.Uuid
		TestArtifactId   uuid.Uuid
		ExpectedResult   bool
	}{
		{r1, a2, r1, a2, true},
		{r1, a2, r2, a2, false},
		{r1, a2, r1, a1, false},
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
				f.RecognizerGrade = j
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
		RecognizerId:      member.NewTenantMemberIdExporter(f.RecognizerId.TenantId, f.RecognizerId.MemberId),
		RecognizerGrade:   exporters.Uint8Exporter(f.RecognizerGrade),
		RecognizerVersion: f.RecognizerVersion,
		SpecialistId:      member.NewTenantMemberIdExporter(f.SpecialistId.TenantId, f.SpecialistId.MemberId),
		SpecialistGrade:   exporters.Uint8Exporter(f.SpecialistGrade),
		SpecialistVersion: f.SpecialistVersion,
		ArtifactId:        artifact.NewTenantArtifactIdExporter(f.ArtifactId.TenantId, f.ArtifactId.ArtifactId),
		CreatedAt:         f.CreatedAt,
	}, actualExporter)
}
