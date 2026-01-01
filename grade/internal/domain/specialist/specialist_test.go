package specialist

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	artifact "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/specialist/assignment"
	"github.com/emacsway/grade/grade/internal/domain/specialist/endorsement"
	"github.com/emacsway/grade/grade/internal/seedwork/domain/exporters"
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
					sf.Commands[0].(ReceivedEndorsementFakeCommand).Artifact.Id.TenantId,
					sf.Commands[0].(ReceivedEndorsementFakeCommand).Artifact.Id.ArtifactId,
				),
				EndorserId:      member.NewMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorserGrade:   exporters.Uint8Exporter(ef.Grade),
				EndorserVersion: 0,
				CreatedAt:       sf.Commands[0].(ReceivedEndorsementFakeCommand).CreatedAt,
			},
			{
				SpecialistId:      member.NewMemberIdExporter(sf.Id.TenantId, sf.Id.MemberId),
				SpecialistGrade:   exporters.Uint8Exporter(sf.Grade),
				SpecialistVersion: 1,
				ArtifactId: artifact.NewArtifactIdExporter(
					sf.Commands[1].(ReceivedEndorsementFakeCommand).Artifact.Id.TenantId,
					sf.Commands[1].(ReceivedEndorsementFakeCommand).Artifact.Id.ArtifactId,
				),
				EndorserId:      member.NewMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorserGrade:   exporters.Uint8Exporter(ef.Grade),
				EndorserVersion: 0,
				CreatedAt:       sf.Commands[1].(ReceivedEndorsementFakeCommand).CreatedAt,
			},
			{
				SpecialistId:      member.NewMemberIdExporter(sf.Id.TenantId, sf.Id.MemberId),
				SpecialistGrade:   exporters.Uint8Exporter(sf.Grade),
				SpecialistVersion: 2,
				ArtifactId: artifact.NewArtifactIdExporter(
					sf.Commands[2].(ReceivedEndorsementFakeCommand).Artifact.Id.TenantId,
					sf.Commands[2].(ReceivedEndorsementFakeCommand).Artifact.Id.ArtifactId,
				),
				EndorserId:      member.NewMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorserGrade:   exporters.Uint8Exporter(ef.Grade),
				EndorserVersion: 0,
				CreatedAt:       sf.Commands[2].(ReceivedEndorsementFakeCommand).CreatedAt,
			},
			{
				SpecialistId:      member.NewMemberIdExporter(sf.Id.TenantId, sf.Id.MemberId),
				SpecialistGrade:   exporters.Uint8Exporter(sf.Grade + 1),
				SpecialistVersion: 3,
				ArtifactId: artifact.NewArtifactIdExporter(
					sf.Commands[3].(ReceivedEndorsementFakeCommand).Artifact.Id.TenantId,
					sf.Commands[3].(ReceivedEndorsementFakeCommand).Artifact.Id.ArtifactId,
				),
				EndorserId:      member.NewMemberIdExporter(ef.Id.TenantId, ef.Id.MemberId),
				EndorserGrade:   exporters.Uint8Exporter(ef.Grade),
				EndorserVersion: 0,
				CreatedAt:       sf.Commands[3].(ReceivedEndorsementFakeCommand).CreatedAt,
			},
		},
		Assignments: []assignment.AssignmentExporter{
			{
				SpecialistId:      member.NewMemberIdExporter(sf.Id.TenantId, sf.Id.MemberId),
				SpecialistVersion: 2,
				AssignedGrade:     exporters.Uint8Exporter(sf.Grade + 1),
				Reason:            exporters.StringExporter("Achieved"),
				CreatedAt:         sf.Commands[2].(ReceivedEndorsementFakeCommand).CreatedAt,
			},
		},
		Version:   4,
		CreatedAt: sf.CreatedAt,
	}, actualExporter)
}
