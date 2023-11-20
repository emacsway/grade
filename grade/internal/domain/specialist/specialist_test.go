package specialist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	artifact "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
	"github.com/emacsway/grade/grade/internal/domain/specialist/assignment"
	"github.com/emacsway/grade/grade/internal/domain/specialist/endorsement"
)

func TestSpecialistExport(t *testing.T) {
	var actualExporter SpecialistExporter
	sf := NewSpecialistFaker()
	err := sf.BuildDependencies()
	require.NoError(t, err)
	ef := endorser.NewEndorserFaker()
	for i := 0; i < 4; i++ {
		err := sf.ReceiveEndorsement(ef)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
	s, err := sf.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	s.Export(&actualExporter)
	assert.Equal(t, SpecialistExporter{
		Id:    member.NewMemberIdExporter(sf.Id.TenantId, sf.Id.MemberId),
		Grade: exporters.Uint8Exporter(sf.Grade + 1),
		ReceivedEndorsements: []endorsement.EndorsementExporter{
			{
				SpecialistId:      member.NewMemberIdExporter(sf.Id.TenantId, sf.Id.MemberId),
				SpecialistGrade:   exporters.Uint8Exporter(sf.Grade),
				SpecialistVersion: 0,
				ArtifactId: artifact.NewArtifactIdExporter(
					sf.ReceivedEndorsements[0].Artifact.Id.TenantId,
					sf.ReceivedEndorsements[0].Artifact.Id.ArtifactId,
				),
				EndorserId:      member.NewMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorserGrade:   exporters.Uint8Exporter(ef.Grade),
				EndorserVersion: 0,
				CreatedAt:       sf.ReceivedEndorsements[0].CreatedAt,
			},
			{
				SpecialistId:      member.NewMemberIdExporter(sf.Id.TenantId, sf.Id.MemberId),
				SpecialistGrade:   exporters.Uint8Exporter(sf.Grade),
				SpecialistVersion: 1,
				ArtifactId: artifact.NewArtifactIdExporter(
					sf.ReceivedEndorsements[1].Artifact.Id.TenantId,
					sf.ReceivedEndorsements[1].Artifact.Id.ArtifactId,
				),
				EndorserId:      member.NewMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorserGrade:   exporters.Uint8Exporter(ef.Grade),
				EndorserVersion: 0,
				CreatedAt:       sf.ReceivedEndorsements[1].CreatedAt,
			},
			{
				SpecialistId:      member.NewMemberIdExporter(sf.Id.TenantId, sf.Id.MemberId),
				SpecialistGrade:   exporters.Uint8Exporter(sf.Grade),
				SpecialistVersion: 2,
				ArtifactId: artifact.NewArtifactIdExporter(
					sf.ReceivedEndorsements[2].Artifact.Id.TenantId,
					sf.ReceivedEndorsements[2].Artifact.Id.ArtifactId,
				),
				EndorserId:      member.NewMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorserGrade:   exporters.Uint8Exporter(ef.Grade),
				EndorserVersion: 0,
				CreatedAt:       sf.ReceivedEndorsements[2].CreatedAt,
			},
			{
				SpecialistId:      member.NewMemberIdExporter(sf.Id.TenantId, sf.Id.MemberId),
				SpecialistGrade:   exporters.Uint8Exporter(sf.Grade + 1),
				SpecialistVersion: 3,
				ArtifactId: artifact.NewArtifactIdExporter(
					sf.ReceivedEndorsements[3].Artifact.Id.TenantId,
					sf.ReceivedEndorsements[3].Artifact.Id.ArtifactId,
				),
				EndorserId:      member.NewMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorserGrade:   exporters.Uint8Exporter(ef.Grade),
				EndorserVersion: 0,
				CreatedAt:       sf.ReceivedEndorsements[3].CreatedAt,
			},
		},
		Assignments: []assignment.AssignmentExporter{
			{
				SpecialistId:      member.NewMemberIdExporter(sf.Id.TenantId, sf.Id.MemberId),
				SpecialistVersion: 2,
				AssignedGrade:     exporters.Uint8Exporter(sf.Grade + 1),
				Reason:            exporters.StringExporter("Achieved"),
				CreatedAt:         sf.ReceivedEndorsements[2].CreatedAt,
			},
		},
		Version:   4,
		CreatedAt: sf.CreatedAt,
	}, actualExporter)
}
