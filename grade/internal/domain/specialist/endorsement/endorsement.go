package endorsement

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
)

type Weight uint8

const (
	LowerWeight  = 0
	PeerWeight   = 1
	HigherWeight = 2
)

func NewEndorsement(
	specialistId member.TenantMemberId,
	specialistGrade grade.Grade,
	specialistVersion uint,
	artifactId artifact.TenantArtifactId,
	endorserId member.TenantMemberId,
	endorserGrade grade.Grade,
	endorserVersion uint,
	createdAt time.Time,
) (Endorsement, error) {
	return Endorsement{
		specialistId:      specialistId,
		specialistGrade:   specialistGrade,
		specialistVersion: specialistVersion,
		artifactId:        artifactId,
		endorserId:        endorserId,
		endorserGrade:     endorserGrade,
		endorserVersion:   endorserVersion,
		createdAt:         createdAt,
	}, nil
}

type Endorsement struct {
	specialistId      member.TenantMemberId
	specialistGrade   grade.Grade
	specialistVersion uint
	artifactId        artifact.TenantArtifactId
	endorserId        member.TenantMemberId
	endorserGrade     grade.Grade
	endorserVersion   uint
	createdAt         time.Time
}

func (e Endorsement) IsEndorsedBy(rId member.TenantMemberId, aId artifact.TenantArtifactId) bool {
	return e.endorserId.Equal(rId) && e.artifactId.Equal(aId)
}

func (e Endorsement) SpecialistGrade() grade.Grade {
	return e.specialistGrade
}

func (e Endorsement) Weight() Weight {
	if e.endorserGrade.Equal(e.specialistGrade) {
		return PeerWeight
	} else if e.endorserGrade.GreaterThan(e.specialistGrade) {
		return HigherWeight
	}
	return LowerWeight
}

func (e Endorsement) Export(ex EndorsementExporterSetter) {
	ex.SetSpecialistId(e.specialistId)
	ex.SetSpecialistGrade(e.specialistGrade)
	ex.SetSpecialistVersion(e.specialistVersion)
	ex.SetArtifactId(e.artifactId)
	ex.SetEndorserId(e.endorserId)
	ex.SetEndorserGrade(e.endorserGrade)
	ex.SetEndorserVersion(e.endorserVersion)
	ex.SetCreatedAt(e.createdAt)
}

type EndorsementExporterSetter interface {
	SetSpecialistId(member.TenantMemberId)
	SetSpecialistGrade(grade.Grade)
	SetSpecialistVersion(uint)
	SetArtifactId(artifact.TenantArtifactId)
	SetEndorserId(member.TenantMemberId)
	SetEndorserGrade(grade.Grade)
	SetEndorserVersion(uint)
	SetCreatedAt(time.Time)
}
