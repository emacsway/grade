package endorsement

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewEndorsementFakeFactory() EndorsementFakeFactory {
	recognizerIdFactory := member.NewTenantMemberIdFakeFactory()
	recognizerIdFactory.MemberId = uuid.ParseSilent("0372d476-9b88-43ff-8b2e-53f560dfc66a")
	specialistIdFactory := member.NewTenantMemberIdFakeFactory()
	specialistIdFactory.MemberId = uuid.ParseSilent("f11a3651-9e7a-41ff-931e-bfe9ef77ba61")
	artifactIdFactory := artifact.NewTenantArtifactIdFakeFactory()
	artifactIdFactory.ArtifactId = uuid.ParseSilent("13ebf39d-30f1-4904-b1f4-993d2987f7d0")
	return EndorsementFakeFactory{
		RecognizerId:      recognizerIdFactory,
		RecognizerGrade:   2,
		RecognizerVersion: 3,
		SpecialistId:      specialistIdFactory,
		SpecialistGrade:   1,
		SpecialistVersion: 5,
		ArtifactId:        artifactIdFactory,
		CreatedAt:         time.Now(),
	}
}

type EndorsementFakeFactory struct {
	RecognizerId      member.TenantMemberIdFakeFactory
	RecognizerGrade   uint8
	RecognizerVersion uint
	SpecialistId      member.TenantMemberIdFakeFactory
	SpecialistGrade   uint8
	SpecialistVersion uint
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
	specialistId, err := member.NewTenantMemberId(f.SpecialistId.TenantId, f.SpecialistId.MemberId)
	if err != nil {
		return Endorsement{}, err
	}
	specialistGrade, err := grade.DefaultConstructor(f.SpecialistGrade)
	if err != nil {
		return Endorsement{}, err
	}
	artifactId, err := artifact.NewTenantArtifactId(f.ArtifactId.TenantId, f.ArtifactId.ArtifactId)
	if err != nil {
		return Endorsement{}, err
	}
	return NewEndorsement(
		recognizerId, recognizerGrade, f.RecognizerVersion,
		specialistId, specialistGrade, f.SpecialistVersion,
		artifactId, f.CreatedAt,
	)
}
