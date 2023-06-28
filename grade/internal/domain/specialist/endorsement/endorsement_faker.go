package endorsement

import (
	"time"

	artifact "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

func NewEndorsementFaker() EndorsementFaker {
	endorserIdFactory := member.NewTenantMemberIdFaker()
	endorserIdFactory.MemberId = endorser.EndorserMemberIdFakeValue
	return EndorsementFaker{
		SpecialistId:      member.NewTenantMemberIdFaker(),
		SpecialistGrade:   1,
		SpecialistVersion: 5,
		ArtifactId:        artifact.NewTenantArtifactIdFaker(),
		EndorserId:        endorserIdFactory,
		EndorserGrade:     2,
		EndorserVersion:   3,
		CreatedAt:         time.Now().Truncate(time.Microsecond),
	}
}

type EndorsementFaker struct {
	SpecialistId      member.TenantMemberIdFaker
	SpecialistGrade   uint8
	SpecialistVersion uint
	ArtifactId        artifact.TenantArtifactIdFaker
	EndorserId        member.TenantMemberIdFaker
	EndorserGrade     uint8
	EndorserVersion   uint
	CreatedAt         time.Time
}

func (f EndorsementFaker) Create() (Endorsement, error) {
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
