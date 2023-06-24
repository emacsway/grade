package endorsement

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

func NewEndorsementFakeFactory() EndorsementFakeFactory {
	endorserIdFactory := member.NewTenantMemberIdFakeFactory()
	endorserIdFactory.MemberId = endorser.EndorserMemberIdFakeValue
	return EndorsementFakeFactory{
		SpecialistId:      member.NewTenantMemberIdFakeFactory(),
		SpecialistGrade:   1,
		SpecialistVersion: 5,
		ArtifactId:        artifact.NewTenantArtifactIdFakeFactory(),
		EndorserId:        endorserIdFactory,
		EndorserGrade:     2,
		EndorserVersion:   3,
		CreatedAt:         time.Now().Truncate(time.Microsecond),
	}
}

type EndorsementFakeFactory struct {
	SpecialistId      member.TenantMemberIdFakeFactory
	SpecialistGrade   uint8
	SpecialistVersion uint
	ArtifactId        artifact.TenantArtifactIdFakeFactory
	EndorserId        member.TenantMemberIdFakeFactory
	EndorserGrade     uint8
	EndorserVersion   uint
	CreatedAt         time.Time
}

func (f EndorsementFakeFactory) Create() (Endorsement, error) {
	specialistId, err := f.SpecialistId.Create()
	if err != nil {
		return Endorsement{}, err
	}
	specialistGrade, err := grade.DefaultConstructor(f.SpecialistGrade)
	if err != nil {
		return Endorsement{}, err
	}
	artifactId, err := f.ArtifactId.Create()
	if err != nil {
		return Endorsement{}, err
	}
	endorserId, err := f.EndorserId.Create()
	if err != nil {
		return Endorsement{}, err
	}
	endorserGrade, err := grade.DefaultConstructor(f.EndorserGrade)
	if err != nil {
		return Endorsement{}, err
	}
	return NewEndorsement(
		specialistId, specialistGrade, f.SpecialistVersion, artifactId,
		endorserId, endorserGrade, f.EndorserVersion, f.CreatedAt,
	)
}
