package endorsement

import (
	"time"

	artifact "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	member "github.com/emacsway/grade/grade/internal/domain/member/values"
)

type EndorsementReconstitutor struct {
	SpecialistId      member.MemberIdReconstitutor
	SpecialistGrade   uint8
	SpecialistVersion uint
	ArtifactId        artifact.ArtifactIdReconstitutor
	EndorserId        member.MemberIdReconstitutor
	EndorserGrade     uint8
	EndorserVersion   uint
	CreatedAt         time.Time
}

func (r EndorsementReconstitutor) Reconstitute() (Endorsement, error) {
	specialistId, err := r.SpecialistId.Reconstitute()
	if err != nil {
		return Endorsement{}, err
	}
	specialistGrade, err := grade.DefaultConstructor(r.SpecialistGrade)
	if err != nil {
		return Endorsement{}, err
	}
	artifactId, err := r.ArtifactId.Reconstitute()
	if err != nil {
		return Endorsement{}, err
	}
	endorserId, err := r.EndorserId.Reconstitute()
	if err != nil {
		return Endorsement{}, err
	}
	endorserGrade, err := grade.DefaultConstructor(r.EndorserGrade)
	if err != nil {
		return Endorsement{}, err
	}
	return Endorsement{
		specialistId:      specialistId,
		specialistGrade:   specialistGrade,
		specialistVersion: r.SpecialistVersion,
		artifactId:        artifactId,
		endorserId:        endorserId,
		endorserGrade:     endorserGrade,
		endorserVersion:   r.EndorserVersion,
		createdAt:         r.CreatedAt,
	}, nil
}
