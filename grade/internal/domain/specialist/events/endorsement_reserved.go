package events

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
)

func NewEndorsementReceived(
	endorserId member.TenantMemberId,
	endorserGrade grade.Grade,
	endorserVersion uint,
	specialistId member.TenantMemberId,
	specialistGrade grade.Grade,
	specialistVersion uint,
	artifactId artifact.TenantArtifactId,
	createdAt time.Time,
) EndorsementReceived {
	return EndorsementReceived{
		endorserId:        endorserId,
		endorserGrade:     endorserGrade,
		endorserVersion:   endorserVersion,
		specialistId:      specialistId,
		specialistGrade:   specialistGrade,
		specialistVersion: specialistVersion,
		artifactId:        artifactId,
		createdAt:         createdAt,
	}
}

type EndorsementReceived struct {
	endorserId        member.TenantMemberId
	endorserGrade     grade.Grade
	endorserVersion   uint
	specialistId      member.TenantMemberId
	specialistGrade   grade.Grade
	specialistVersion uint
	artifactId        artifact.TenantArtifactId
	createdAt         time.Time
}

func (e EndorsementReceived) EndorserId() member.TenantMemberId {
	return e.endorserId
}

func (e EndorsementReceived) EndorserGrade() grade.Grade {
	return e.endorserGrade
}

func (e EndorsementReceived) EndorserVersion() uint {
	return e.endorserVersion
}

func (e EndorsementReceived) SpecialistId() member.TenantMemberId {
	return e.specialistId
}

func (e EndorsementReceived) SpecialistGrade() grade.Grade {
	return e.specialistGrade
}

func (e EndorsementReceived) SpecialistVersion() uint {
	return e.specialistVersion
}

func (e EndorsementReceived) ArtifactId() artifact.TenantArtifactId {
	return e.artifactId
}

func (e EndorsementReceived) CreatedAt() time.Time {
	return e.createdAt
}
