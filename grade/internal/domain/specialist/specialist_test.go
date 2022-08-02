package specialist

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/specialist/assignment"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/specialist/endorsement"
)

func TestSpecialistExport(t *testing.T) {
	var actualExporter SpecialistExporter
	ef := NewSpecialistFakeFactory()
	rf := recognizer.NewRecognizerFakeFactory()
	for i := 0; i < 4; i++ {
		err := ef.ReceiveEndorsement(rf)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
	agg, err := ef.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	agg.Export(&actualExporter)
	assert.Equal(t, SpecialistExporter{
		Id:    member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
		Grade: seedwork.Uint8Exporter(ef.Grade + 1),
		ReceivedEndorsements: []endorsement.EndorsementExporter{
			{
				RecognizerId:      member.NewTenantMemberIdExporter(rf.Id.TenantId, rf.Id.MemberId),
				RecognizerGrade:   seedwork.Uint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				SpecialistId:      member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				SpecialistGrade:   seedwork.Uint8Exporter(ef.Grade),
				SpecialistVersion: 0,
				ArtifactId: artifact.NewTenantArtifactIdExporter(
					ef.ReceivedEndorsements[0].Artifact.Id.TenantId,
					ef.ReceivedEndorsements[0].Artifact.Id.ArtifactId,
				),
				CreatedAt: ef.ReceivedEndorsements[0].CreatedAt,
			},
			{
				RecognizerId:      member.NewTenantMemberIdExporter(rf.Id.TenantId, rf.Id.MemberId),
				RecognizerGrade:   seedwork.Uint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				SpecialistId:      member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				SpecialistGrade:   seedwork.Uint8Exporter(ef.Grade),
				SpecialistVersion: 1,
				ArtifactId: artifact.NewTenantArtifactIdExporter(
					ef.ReceivedEndorsements[1].Artifact.Id.TenantId,
					ef.ReceivedEndorsements[1].Artifact.Id.ArtifactId,
				),
				CreatedAt: ef.ReceivedEndorsements[1].CreatedAt,
			},
			{
				RecognizerId:      member.NewTenantMemberIdExporter(rf.Id.TenantId, rf.Id.MemberId),
				RecognizerGrade:   seedwork.Uint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				SpecialistId:      member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				SpecialistGrade:   seedwork.Uint8Exporter(ef.Grade),
				SpecialistVersion: 2,
				ArtifactId: artifact.NewTenantArtifactIdExporter(
					ef.ReceivedEndorsements[2].Artifact.Id.TenantId,
					ef.ReceivedEndorsements[2].Artifact.Id.ArtifactId,
				),
				CreatedAt: ef.ReceivedEndorsements[2].CreatedAt,
			},
			{
				RecognizerId:      member.NewTenantMemberIdExporter(rf.Id.TenantId, rf.Id.MemberId),
				RecognizerGrade:   seedwork.Uint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				SpecialistId:      member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				SpecialistGrade:   seedwork.Uint8Exporter(ef.Grade + 1),
				SpecialistVersion: 3,
				ArtifactId: artifact.NewTenantArtifactIdExporter(
					ef.ReceivedEndorsements[3].Artifact.Id.TenantId,
					ef.ReceivedEndorsements[3].Artifact.Id.ArtifactId,
				),
				CreatedAt: ef.ReceivedEndorsements[3].CreatedAt,
			},
		},
		Assignments: []assignment.AssignmentExporter{
			{
				SpecialistId:      member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				SpecialistVersion: 2,
				AssignedGrade:     seedwork.Uint8Exporter(ef.Grade + 1),
				Reason:            seedwork.StringExporter("Achieved"),
				CreatedAt:         ef.ReceivedEndorsements[2].CreatedAt,
			},
		},
		Version:   4,
		CreatedAt: ef.CreatedAt,
	}, actualExporter)
}
