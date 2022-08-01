package endorsement

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

func NewEndorsementFakeFactory() EndorsementFakeFactory {
	recognizerIdFactory := member.NewTenantMemberIdFakeFactory()
	recognizerIdFactory.MemberId = 1
	endorsedIdFactory := member.NewTenantMemberIdFakeFactory()
	endorsedIdFactory.MemberId = 2
	artifactIdFactory := artifact.NewTenantArtifactIdFakeFactory()
	artifactIdFactory.ArtifactId = 6
	return EndorsementFakeFactory{
		RecognizerId:      recognizerIdFactory,
		RecognizerGrade:   2,
		RecognizerVersion: 3,
		EndorsedId:        endorsedIdFactory,
		EndorsedGrade:     1,
		EndorsedVersion:   5,
		ArtifactId:        artifactIdFactory,
		CreatedAt:         time.Now(),
	}
}

type EndorsementFakeFactory struct {
	RecognizerId      member.TenantMemberIdFakeFactory
	RecognizerGrade   uint8
	RecognizerVersion uint
	EndorsedId        member.TenantMemberIdFakeFactory
	EndorsedGrade     uint8
	EndorsedVersion   uint
	ArtifactId        artifact.TenantArtifactIdFakeFactory
	CreatedAt         time.Time
}

func (f EndorsementFakeFactory) Create() (Endorsement, error) {
	recognizerId, err := member.NewTenantMemberId(f.RecognizerId.TenantId, f.RecognizerId.MemberId)
	if err != nil {
		return Endorsement{}, err
	}
	recognizerGrade, err := grade.DefaultConstructor(f.RecognizerGrade)
	if err != nil {
		return Endorsement{}, err
	}
	endorsedId, err := member.NewTenantMemberId(f.EndorsedId.TenantId, f.EndorsedId.MemberId)
	if err != nil {
		return Endorsement{}, err
	}
	endorsedGrade, err := grade.DefaultConstructor(f.EndorsedGrade)
	if err != nil {
		return Endorsement{}, err
	}
	artifactId, err := artifact.NewTenantArtifactId(f.ArtifactId.TenantId, f.ArtifactId.ArtifactId)
	if err != nil {
		return Endorsement{}, err
	}
	return NewEndorsement(
		recognizerId, recognizerGrade, f.RecognizerVersion,
		endorsedId, endorsedGrade, f.EndorsedVersion,
		artifactId, f.CreatedAt,
	)
}
