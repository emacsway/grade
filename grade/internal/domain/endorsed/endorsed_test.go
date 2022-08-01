package endorsed

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/assignment"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func TestEndorsedExport(t *testing.T) {
	var actualExporter EndorsedExporter
	ef := NewEndorsedFakeFactory()
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
	assert.Equal(t, EndorsedExporter{
		Id:    member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
		Grade: seedwork.Uint8Exporter(ef.Grade + 1),
		ReceivedEndorsements: []endorsement.EndorsementExporter{
			{
				RecognizerId:      member.NewTenantMemberIdExporter(rf.Id.TenantId, rf.Id.MemberId),
				RecognizerGrade:   seedwork.Uint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				EndorsedId:        member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorsedGrade:     seedwork.Uint8Exporter(ef.Grade),
				EndorsedVersion:   0,
				ArtifactId:        seedwork.Uint64Exporter(ef.ReceivedEndorsements[0].ArtifactId),
				CreatedAt:         ef.ReceivedEndorsements[0].CreatedAt,
			},
			{
				RecognizerId:      member.NewTenantMemberIdExporter(rf.Id.TenantId, rf.Id.MemberId),
				RecognizerGrade:   seedwork.Uint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				EndorsedId:        member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorsedGrade:     seedwork.Uint8Exporter(ef.Grade),
				EndorsedVersion:   1,
				ArtifactId:        seedwork.Uint64Exporter(ef.ReceivedEndorsements[1].ArtifactId),
				CreatedAt:         ef.ReceivedEndorsements[1].CreatedAt,
			},
			{
				RecognizerId:      member.NewTenantMemberIdExporter(rf.Id.TenantId, rf.Id.MemberId),
				RecognizerGrade:   seedwork.Uint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				EndorsedId:        member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorsedGrade:     seedwork.Uint8Exporter(ef.Grade),
				EndorsedVersion:   2,
				ArtifactId:        seedwork.Uint64Exporter(ef.ReceivedEndorsements[2].ArtifactId),
				CreatedAt:         ef.ReceivedEndorsements[2].CreatedAt,
			},
			{
				RecognizerId:      member.NewTenantMemberIdExporter(rf.Id.TenantId, rf.Id.MemberId),
				RecognizerGrade:   seedwork.Uint8Exporter(rf.Grade),
				RecognizerVersion: 0,
				EndorsedId:        member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorsedGrade:     seedwork.Uint8Exporter(ef.Grade + 1),
				EndorsedVersion:   3,
				ArtifactId:        seedwork.Uint64Exporter(ef.ReceivedEndorsements[3].ArtifactId),
				CreatedAt:         ef.ReceivedEndorsements[3].CreatedAt,
			},
		},
		Assignments: []assignment.AssignmentExporter{
			{
				EndorsedId:      member.NewTenantMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorsedVersion: 2,
				AssignedGrade:   seedwork.Uint8Exporter(ef.Grade + 1),
				Reason:          seedwork.StringExporter("Achieved"),
				CreatedAt:       ef.ReceivedEndorsements[2].CreatedAt,
			},
		},
		Version:   4,
		CreatedAt: ef.CreatedAt,
	}, actualExporter)
}
